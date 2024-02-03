package routes

import (
	"github.com/fasthttp/router"

	"github.com/kyleu/dbaudit/app/controller"
)

func parseRoutes(r *router.Router) {
	r.GET("/parse", controller.ParseForm)
	r.POST("/parse", controller.Parse)
}
