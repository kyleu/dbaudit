// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"

	"github.com/kyleu/dbaudit/app/controller"
)

func generatedRoutes(r *router.Router) {
	r.GET("/statement", controller.StatementList)
	r.GET("/statement/_new", controller.StatementCreateForm)
	r.POST("/statement/_new", controller.StatementCreate)
	r.GET("/statement/_random", controller.StatementRandom)
	r.GET("/statement/{id}", controller.StatementDetail)
	r.GET("/statement/{id}/edit", controller.StatementEditForm)
	r.POST("/statement/{id}/edit", controller.StatementEdit)
	r.GET("/statement/{id}/delete", controller.StatementDelete)
}
