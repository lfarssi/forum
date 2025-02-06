package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var (
	clientID     = "882176403287-gjj6tgun26m6m40je243cai9sp4ja347.apps.googleusercontent.com"
	clientSecret = "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"
	redirectURI  = "http://localhost:8080/callback"
	authURL      = "https://accounts.google.com/o/oauth2/auth"
	tokenURL     = "https://oauth2.googleapis.com/token"
	scope        = "https://www.googleapis.com/auth/userinfo.profile"
)

type Auth struct {
	ClientId     string
	ClientSecret string
	RedirectURI  string
	AuthURL      string
	TokenURL     string
	Scope        string
	FetchURL     string
}

var Github = Auth{
	ClientId:     "Ov23liFzgwDEBpPQYEMO",
	ClientSecret: "dc944e0e15e15837e07382baef81fc8cc932d1dd",
	RedirectURI:  "http://localhost:8080/callback",
	AuthURL:      "https://github.com/login/oauth/authorize",
	TokenURL:     "https://github.com/login/oauth/access_token",
	Scope:        "user",
	FetchURL: "https://api.github.com/user",
}

var Google = Auth{
	ClientId:     "882176403287-gjj6tgun26m6m40je243cai9sp4ja347.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURI:  "http://localhost:8080/callback",
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://oauth2.googleapis.com/token",
	Scope:        "https://www.googleapis.com/auth/userinfo.profile",
	FetchURL: "https://www.googleapis.com/oauth2/v2/userinfo",
}

type GOOuser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
}

type GithubUser struct {
	Login      string `json:"login"`
	Id         string `json:"id"`
	Avatar_url string `json:"avatar_url"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

var User GOOuser

var AuthProvider Auth
func OAuthController(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	if code == "" {
		ErrorController(w, r, http.StatusBadRequest, "Missing code")
		return
	}
	fmt.Println(code)

	// Exchange the authorization code for an access token
	token, err := exchangeCodeForToken(code)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Failed to exchange code for token")
		return
	}
	fmt.Println(token)
	// Use the access token to fetch user data
	err = fetchUserData(token)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Failed to fetch user data")
		return
	}

	fmt.Println(User)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

}

func exchangeCodeForToken(code string) (string, error) {
	// Prepare the request to exchange the code for a token
	reqBody := fmt.Sprintf("client_id=%s&client_secret=%s&code=%s&redirect_uri=%s&grant_type=authorization_code",
		AuthProvider.ClientId, AuthProvider.ClientSecret, code, AuthProvider.RedirectURI)

	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.AccessToken, nil

}

func fetchUserData(accessToken string) error {
	req, err := http.NewRequest("GET", AuthProvider.FetchURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&User); err != nil {
		return err
	}

	return nil
}

func RedirectController(w http.ResponseWriter, r *http.Request) {
	var url string
	if r.Method != http.MethodPost {
		ErrorController(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	if r.URL.Query().Get("service") == "github" {
		AuthProvider = Github
		url = fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&scope=%s",AuthProvider.AuthURL, AuthProvider.ClientId, AuthProvider.RedirectURI, AuthProvider.Scope)
	} else if r.URL.Query().Get("service") == "google" {
		AuthProvider = Google
		url = fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s",AuthProvider.AuthURL, AuthProvider.ClientId, AuthProvider.RedirectURI, AuthProvider.Scope)
	}
	fmt.Println("https://github.com/login/oauth/authorize?client_id=Ov23liFzgwDEBpPQYEMO&redirect_uri=http://localhost:8080/callback&scope=user")
	fmt.Println(url)
	http.Redirect(w, r,url ,http.StatusPermanentRedirect)
}
