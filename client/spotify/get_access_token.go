package spotify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/rs/zerolog/log"
)

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func (spotifyClient *spotifyClient) GetAccessToken(authorizationCode string) (response *AccessTokenResponse, err error) {
	form := url.Values{}
	form.Set("code", authorizationCode)
	form.Set("redirect_uri", spotifyClient.oauthRedirectURI)
	form.Set("grant_type", "authorization_code")
	bodyReader := strings.NewReader(form.Encode())

	accessTokenRequestURL := spotifyClient.oauthBaseUrl + "/api/token"
	accessTokenRequest, err := http.NewRequest(http.MethodPost, accessTokenRequestURL, bodyReader)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("could not create http request")
	}

	basicAuthValue := base64.StdEncoding.EncodeToString([]byte(spotifyClient.oauthClientID + ":" + spotifyClient.oauthClientSecret))
	accessTokenRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	accessTokenRequest.Header.Set("Authorization", fmt.Sprintf("Basic %s", basicAuthValue))

	log.Debug().
		Str("basic auth value", basicAuthValue).
		Str("token endpoint", accessTokenRequestURL).
		Str("form body", form.Encode()).
		Msg("retrieving access token")

	accessTokenResponse, err := http.DefaultClient.Do(accessTokenRequest)
	if err != nil {
		return nil, errors.New("could not retrieve access token from spotify")
	}

	accessTokenResponseBody, err := io.ReadAll(accessTokenResponse.Body)
	if err != nil {
		return nil, errors.New("could not read response from spotify access token endpoint")
	}

	var accessTokenResponseContent AccessTokenResponse
	err = json.Unmarshal(accessTokenResponseBody, &accessTokenResponseContent)
	if err != nil {
		log.Error().
			Err(err).
			Str("access token response", string(accessTokenResponseBody)).
			Msg("could not parse response from spotify access token endpoint")
		return nil, errors.New("could not parse response from spotify access token endpoint")
	}

	return &accessTokenResponseContent, nil
}
