package officegraph

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

type TokenResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func GetAuthToken() (string, error) {

	clientID := viper.GetString("APPCLIENT_ID")
	clientSecret := viper.GetString("APPCLIENT_SECRET")
	tenantID := viper.GetString("APPCLIENT_DOMAIN")

	tokenEndpoint := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenantID)

	// Construct the request body
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "https://graph.microsoft.com/.default")
	//data.Set("scope", "TeamMember.Read.All Place.Read.All")

	// Send the HTTP request to the token endpoint
	resp, err := http.Post(tokenEndpoint, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the JSON response
	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", err
	}

	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("access token not found in response")
	}

	return tokenResp.AccessToken, nil
}
