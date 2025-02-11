package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"forum/app/models"

	"github.com/gofrs/uuid"
)

// AuthConfig holds OAuth configuration details.
type AuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	AuthURL      string
	TokenURL     string
	Scope        string
	FetchURL     string
}

// GoogleUser represents the Google user profile.
type GoogleUser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Email      string `json:"email"`
}

// GithubUser represents the GitHub user profile.
type GithubUser struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

// Predefined OAuth configurations
var (
	Github = AuthConfig{
		ClientID:     "Ov23li3Any8hCic0eddy",
		ClientSecret: "7f052daad89c01aae86dccaf3d90f8fdf82c2910",
		RedirectURI:  "http://localhost:8080/callback",
		AuthURL:      "https://github.com/login/oauth/authorize",
		TokenURL:     "https://github.com/login/oauth/access_token",
		Scope:        "user:email read:user",
		FetchURL:     "https://api.github.com/user",
	}

	Google = AuthConfig{
		ClientID:     "882176403287-gjj6tgun26m6m40je243cai9sp4ja347.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURI:  "http://localhost:8080/callback",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://oauth2.googleapis.com/token",
		Scope:        "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile",
		FetchURL:     "https://www.googleapis.com/oauth2/v2/userinfo",
	}

	currentService AuthConfig // Stores the current OAuth service (Google or GitHub)
)

// HandleGoogleLogin redirects the user to the OAuth authorization endpoint based on the "service" query parameter.
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
    serviceName := r.URL.Query().Get("service")

    var authURL string
    switch serviceName {
    case "github":
        currentService = Github
        authURL = fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s",
            currentService.AuthURL, currentService.ClientID, currentService.RedirectURI, currentService.Scope)
    case "google":
        currentService = Google
        authURL = fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&prompt=consent",
            currentService.AuthURL, currentService.ClientID, currentService.RedirectURI, currentService.Scope)
    default:
        ErrorController(w, r, http.StatusBadRequest, "Invalid service parameter")
        return
    }

  //  For GitHub, use HTML with JavaScript to ensure proper redirect
    if serviceName == "github" {
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, `
            <!DOCTYPE html>
            <html>
            <head>
                <title>Redirecting to GitHub...</title>
            </head>
            <body>
                <script>
                    window.location.replace(%q);
                </script>
                <noscript>
                    <meta http-equiv="refresh" content="0;url=%s">
                </noscript>
                <p>Redirecting to GitHub login...</p>
            </body>
            </html>
        `, authURL, authURL)
        return
    }

    // For Google, use regular redirect
    http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}
// HandleGoogleCallback processes the OAuth callback.
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		ErrorController(w, r, http.StatusBadRequest, "Missing code parameter")
		return
	}

	accessToken, err := exchangeCodeForToken(code)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, fmt.Sprintf("Failed to exchange code: %v", err))
		return
	}

	email, username, err := fetchUserInfo(accessToken)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, fmt.Sprintf("Failed to fetch user info: %v", err))
		return
	}

	err = performOAuthLogin(w, username, email)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, fmt.Sprintf("OAuth login failed: %v", err))
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// exchangeCodeForToken exchanges the authorization code for an access token.
func exchangeCodeForToken(code string) (string, error) {
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", currentService.ClientID)
	data.Set("client_secret", currentService.ClientSecret)
	data.Set("redirect_uri", currentService.RedirectURI)
	data.Set("grant_type", "authorization_code")

	resp, err := http.PostForm(currentService.TokenURL, data)
	if err != nil {
		return "", fmt.Errorf("failed to exchange code for token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("token exchange failed with status: %d, body: %s", resp.StatusCode, string(body))
	}

	var tokenData map[string]interface{}

	if currentService == Github {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read response body: %w", err)
		}
		values, err := url.ParseQuery(string(body))
		if err != nil {
			return "", fmt.Errorf("failed to parse response body: %w", err)
		}
		accessToken := values.Get("access_token")
		if accessToken == "" {
			return "", fmt.Errorf("no access token in token response")
		}
		return accessToken, nil

	} else {
		if err := json.NewDecoder(resp.Body).Decode(&tokenData); err != nil {
			return "", fmt.Errorf("failed to decode token data: %w", err)
		}
	}
	accessToken, ok := tokenData["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token not found in response")
	}

	return accessToken, nil
}

// fetchUserInfo fetches user information from the respective OAuth provider.
func fetchUserInfo(accessToken string) (email, username string, err error) {
	switch currentService {
	case Google:
		email, username, err = fetchGoogleUserInfo(accessToken)
	case Github:
		email, username, err = fetchGithubUserInfo(accessToken)
	default:
		err = fmt.Errorf("unsupported OAuth service")
	}
	return
}

// fetchGoogleUserInfo fetches user information from Google.
func fetchGoogleUserInfo(accessToken string) (email, username string, err error) {
	userInfo, err := fetchUserProfile(accessToken, currentService.FetchURL)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch Google user info: %w", err)
	}

	var googleUser GoogleUser
	if err := json.Unmarshal(userInfo, &googleUser); err != nil {
		return "", "", fmt.Errorf("failed to unmarshal Google user info: %w", err)
	}
	return googleUser.Email, googleUser.GivenName, nil
}

// fetchGithubUserInfo fetches user information from GitHub.
func fetchGithubUserInfo(accessToken string) (email, username string, error error) {
	userInfo, err := fetchUserProfile(accessToken, currentService.FetchURL)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch Github user info: %w", err)
	}

	var githubUser GithubUser
	if err := json.Unmarshal(userInfo, &githubUser); err != nil {
		return "", "", fmt.Errorf("failed to unmarshal Github user info: %w", err)
	}

	email, err = fetchPrimaryEmailFromGithub(accessToken)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch primary email from Github: %w", err)
	}
	return email, githubUser.Name, nil
}

// fetchUserProfile retrieves the user profile from the specified URL.
func fetchUserProfile(accessToken string, fetchURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", fetchURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("fetch user profile failed with status: %d, body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

// fetchPrimaryEmailFromGithub fetches the primary email address from Github.
func fetchPrimaryEmailFromGithub(accessToken string) (string, error) {
	emailReq, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	emailReq.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	emailResp, err := client.Do(emailReq)
	if err != nil {
		return "", fmt.Errorf("failed to fetch user emails: %w", err)
	}
	defer emailResp.Body.Close()

	if emailResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(emailResp.Body)
		return "", fmt.Errorf("fetch github emails failed with status: %d, body: %s", emailResp.StatusCode, string(body))
	}

	var emails []map[string]interface{}
	if err := json.NewDecoder(emailResp.Body).Decode(&emails); err != nil {
		return "", fmt.Errorf("failed to decode email data: %w", err)
	}

	for _, email := range emails {
		if primary, ok := email["primary"].(bool); ok && primary {
			if emailValue, ok := email["email"].(string); ok {
				return emailValue, nil
			}
		}
	}

	return "", fmt.Errorf("no primary email found")
}

// performOAuthLogin performs the OAuth login or registration process.
func performOAuthLogin(w http.ResponseWriter, username, email string) error {
	userID, err := models.OAuthlogin(username, email)
	if err != nil {
		// User doesn't exist, register them
		userID, err = registerOAuthUser(username, email)
		if err != nil {
			return err
		}
	}

	return createSessionAndSetCookie(w, userID)
}

// registerOAuthUser registers a new user via OAuth.
func registerOAuthUser(username, email string) (int, error) {
	// Generate a more suitable password.  Using 0 or the Github ID is not secure.
	// Consider using a randomly generated string. For simplicity, I'm using a timestamp here, but this should be improved in a real application.
	password := strconv.FormatInt(time.Now().UnixNano(), 10) // Generate password

	user := models.User{UserName: username, Email: email, Password: password}
	userID, err := models.OAuthRegistration(user)
	if err != nil {
		return 0, fmt.Errorf("failed to register user: %w", err)
	}
	return userID, nil
}

// createSessionAndSetCookie creates a session and sets the token cookie.
func createSessionAndSetCookie(w http.ResponseWriter, userID int) error {
	token, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("cannot generate token: %w", err)
	}

	err = models.CreateSession(userID, token.String(), time.Now().Add(24*time.Hour))
	if err != nil {
		return fmt.Errorf("cannot create session: %w", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token.String(),
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	return nil
}
