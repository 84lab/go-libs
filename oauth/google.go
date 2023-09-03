package oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type GoogleOauthToken struct {
	AccessToken string
	IDToken     string
}

type GoogleUserResult struct {
	ID            string
	Email         string
	VerifiedEmail bool
	Name          string
	GivenName     string
	FamilyName    string
	Picture       string
	Locale        string
}

//nolint:goerr113,forcetypeassert
func GetGoogleOauthToken(code, clientID, clientSecret, oauthRedirectURL string) (*GoogleOauthToken, error) {
	const rootURL = "https://oauth2.googleapis.com/token"

	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", clientID)
	values.Add("client_secret", clientSecret)
	values.Add("redirect_uri", oauthRedirectURL)
	values.Add("access_type", "offline")

	query := values.Encode()

	req, err := http.NewRequest("POST", rootURL, bytes.NewBufferString(query))
	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do a request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	var GoogleOauthTokenRes map[string]interface{}

	if err := json.Unmarshal(resBody, &GoogleOauthTokenRes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	tokenBody := &GoogleOauthToken{
		AccessToken: GoogleOauthTokenRes["access_token"].(string),
		IDToken:     GoogleOauthTokenRes["id_token"].(string),
	}

	return tokenBody, nil
}

//nolint:goerr113,forcetypeassert
func GetGoogleUser(accessToken, idToken string) (*GoogleUserResult, error) {
	rootURL := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", accessToken)

	req, err := http.NewRequest("GET", rootURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", idToken))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do a request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	var GoogleUserRes map[string]interface{}

	if err := json.Unmarshal(resBody, &GoogleUserRes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	userBody := &GoogleUserResult{
		ID:            GoogleUserRes["id"].(string),
		Email:         GoogleUserRes["email"].(string),
		VerifiedEmail: GoogleUserRes["verified_email"].(bool),
		Name:          GoogleUserRes["name"].(string),
		GivenName:     GoogleUserRes["given_name"].(string),
		Picture:       GoogleUserRes["picture"].(string),
		Locale:        GoogleUserRes["locale"].(string),
	}

	return userBody, nil
}
