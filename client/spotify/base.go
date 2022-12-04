package spotify

type SpotifyClient interface {
	GetAccessToken(authorizationCode string) (response *AccessTokenResponse, err error)
}

type spotifyClient struct {
	oauthBaseUrl      string
	apiBaseUrl        string
	oauthClientID     string
	oauthClientSecret string
	oauthRedirectURI  string
}

func NewSpotifyClient(
	oauthClientID string,
	oauthClientSecret string,
	oauthRedirectURI string,
) SpotifyClient {
	return &spotifyClient{
		oauthBaseUrl:      "https://accounts.spotify.com",
		apiBaseUrl:        "https://api.spotify.com",
		oauthClientID:     oauthClientID,
		oauthClientSecret: oauthClientSecret,
		oauthRedirectURI:  oauthRedirectURI,
	}
}
