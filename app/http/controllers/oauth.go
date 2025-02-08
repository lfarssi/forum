package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strings"
// )

import (
	"encoding/json"
	"fmt"
	"forum/app/models"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
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
	FetchURL:     "https://api.github.com/user",
}

var Google = Auth{
	ClientId:     "882176403287-gjj6tgun26m6m40je243cai9sp4ja347.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURI:  "http://localhost:8080/callback",
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://oauth2.googleapis.com/token",
	Scope:        "https://www.googleapis.com/auth/userinfo.profile",
	FetchURL:     "https://www.googleapis.com/oauth2/v2/userinfo",
}

type GOOuser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
}

type GithubUser struct {
	Login      string `json:"login"`
	Id         int    `json:"id"`
	Avatar_url string `json:"avatar_url"`
	UserName   string `json:"name"`
	Email      string `json:"email"`
}

var Service Auth

// handleGoogleLogin redirects the user to Google's OAuth 2.0 authorization endpoint.
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("service") == "github" {
		Service = Github
	} else {
		Service = Google
	}
	url := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&prompt=consent",
		Service.AuthURL, Service.ClientId, Service.RedirectURI, Service.Scope)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// handleGoogleCallback processes the OAuth callback from Google.
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Extract the authorization code from the query parameters.
	code := r.URL.Query().Get("code")
	if code == "" {
		ErrorController(w, r, http.StatusBadRequest, "Missing code parameter")
		return
	}

	// Exchange the authorization code for an access token.
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", Service.ClientId)
	data.Set("client_secret", Service.ClientSecret)
	data.Set("redirect_uri", Service.RedirectURI)
	data.Set("grant_type", "authorization_code")

	resp, err := http.PostForm(Service.TokenURL, data)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Failed to exchange code for token")
		return
	}
	var accessToken string
	if Service == Github {
		body, _ := io.ReadAll(resp.Body)
		body1 := string(body)
		b, _ := url.ParseQuery(body1)
		accessToken = b.Get("access_token")
	} else {
		var tokenData map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&tokenData); err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Failed to decode token data")
			return
		}
		accessToken1, ok := tokenData["access_token"].(string)
		if !ok {
			ErrorController(w, r, http.StatusInternalServerError, "No access token in token response")
			return
		}
		accessToken = accessToken1
	}
	defer resp.Body.Close()

	// Fetch the user's profile information.
	req, _ := http.NewRequest("GET", Service.FetchURL, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Failed to fetch user data")
		return
	}
	defer resp.Body.Close()

	// Parse and display the user profile.
	var UserInfo GOOuser
	var UserInfo1 GithubUser
	userInfo, _ := io.ReadAll(resp.Body)
	if Service == Google {
		if err = json.Unmarshal(userInfo, &UserInfo); err != nil {

			ErrorController(w, r, http.StatusNotFound, "Page not found")
			return
		}
		id, err := models.OAuthlogin(UserInfo.Username, UserInfo.Name)
		if err != nil {
			id, err1 := models.OAuthRegistration(models.User{UserName: UserInfo.Username, Email: UserInfo.Name, Password: UserInfo.ID})
			if err1 != nil {
				fmt.Println("jjjjjj")
				fmt.Println(err1)
				ErrorController(w, r, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			token, err := uuid.NewV4()
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "Cannot Generate token")
				return
			}
			err = models.CreateSession(id, token.String(), time.Now().Add((24 * time.Hour)))
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "Cannot Create Sessions")
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    token.String(),
				Expires:  time.Now().Add((24 * time.Hour)),
				HttpOnly: true,
			})
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		token, err := uuid.NewV4()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Generate token")
			return
		}
		err = models.CreateSession(id, token.String(), time.Now().Add((24 * time.Hour)))
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Create Sessions")
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token.String(),
			Expires:  time.Now().Add((24 * time.Hour)),
			HttpOnly: true,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		if err = json.Unmarshal(userInfo, &UserInfo1); err != nil {

			ErrorController(w, r, http.StatusNotFound, "Page not found")
			return
		}
		id, err := models.OAuthlogin(UserInfo1.UserName, UserInfo1.Email)
		if err != nil {
			Id := strconv.Itoa(UserInfo1.Id)
			fmt.Println(Id)
			id, err := models.OAuthRegistration(models.User{UserName: UserInfo1.UserName, Email: UserInfo1.Email, Password: Id})
			if err != nil {
				fmt.Println("suii")
				fmt.Println(err)
				ErrorController(w, r, http.StatusInternalServerError, "Internal Server Error")
				return

			}
			token, err := uuid.NewV4()
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "Cannot Generate token")
				return
			}
			err = models.CreateSession(id, token.String(), time.Now().Add((24 * time.Hour)))
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "Cannot Create Sessions")
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    token.String(),
				Expires:  time.Now().Add((24 * time.Hour)),
				HttpOnly: true,
			})
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		token, err := uuid.NewV4()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Generate token")
			return
		}
		err = models.CreateSession(id, token.String(), time.Now().Add((24 * time.Hour)))
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Create Sessions")
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token.String(),
			Expires:  time.Now().Add((24 * time.Hour)),
			HttpOnly: true,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
