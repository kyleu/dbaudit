// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"

	"github.com/kyleu/dbaudit/app/controller"
)

func generatedRoutes(r *router.Router) {
	r.GET("/db", controller.ConnectionList)
	r.GET("/db/_new", controller.ConnectionCreateForm)
	r.POST("/db/_new", controller.ConnectionCreate)
	r.GET("/db/_random", controller.ConnectionRandom)
	r.GET("/db/{id}", controller.ConnectionDetail)
	r.GET("/db/{id}/edit", controller.ConnectionEditForm)
	r.POST("/db/{id}/edit", controller.ConnectionEdit)
	r.GET("/db/{id}/delete", controller.ConnectionDelete)
	r.GET("/statement", controller.StatementList)
	r.GET("/statement/_new", controller.StatementCreateForm)
	r.POST("/statement/_new", controller.StatementCreate)
	r.GET("/statement/_random", controller.StatementRandom)
	r.GET("/statement/{id}", controller.StatementDetail)
	r.GET("/statement/{id}/edit", controller.StatementEditForm)
	r.POST("/statement/{id}/edit", controller.StatementEdit)
	r.GET("/statement/{id}/delete", controller.StatementDelete)
}
