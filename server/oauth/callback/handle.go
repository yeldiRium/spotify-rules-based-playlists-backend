package callback

import (
	"fmt"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/julienschmidt/httprouter"
	"github.com/yeldiRium/spotify-rules-based-playlists-backend/client/spotify"
	"github.com/yeldiRium/spotify-rules-based-playlists-backend/server/api"
)

func Handle(
	spotifyClient spotify.SpotifyClient,
	redirectURL string,
	sessionManager *scs.SessionManager,
) api.Route {
	return func(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) *api.Error {
		// This should probably somehow validate the state parameter. Oh well. I don't care at the moment.

		authorizationCode := r.URL.Query().Get("code")
		if authorizationCode == "" {
			return &api.Error{
				Code:    http.StatusBadRequest,
				Message: "the request does not contain an authorization code",
			}
		}

		accessTokenResponse, err := spotifyClient.GetAccessToken(authorizationCode)
		if err != nil {
			return &api.Error{
				Code:    http.StatusForbidden,
				Message: fmt.Sprintf("could not acquire an access token from spotify: %s", err),
			}
		}

		sessionManager.Put(r.Context(), "accessToken", accessTokenResponse.AccessToken)
		sessionManager.Put(r.Context(), "refreshToken", accessTokenResponse.RefreshToken)
		sessionManager.Put(r.Context(), "expiresIn", accessTokenResponse.ExpiresIn)

		// TODO: fetch user information and store it too

		http.Redirect(writer, r, redirectURL, http.StatusSeeOther)

		return nil
	}
}
