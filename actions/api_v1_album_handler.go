package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func apiV1AlbumHandler(c buffalo.Context) error {
	things := []string{"album1", "album2", "album3"}
	return c.Render(http.StatusOK, r.JSON(things))
}
