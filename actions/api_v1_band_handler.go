package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func apiV1BandHandler(c buffalo.Context) error {
	things := []string{"band1", "band2", "band3"}
	return c.Render(http.StatusOK, r.JSON(things))
}
