package health

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/yeldiRium/spotify-rules-based-playlists-backend/server/api"
)

func Handle() api.Route {
	return func(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) *api.Error {
		writer.Write([]byte(http.StatusText(http.StatusOK)))

		return nil
	}
}
