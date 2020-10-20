package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func apiV1ArtistHandler(c buffalo.Context) error {
	things := []string{"artist1", "artist2", "artist3"}
	return c.Render(http.StatusOK, r.JSON(things))
}
