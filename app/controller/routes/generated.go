// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/dbaudit/app/controller"
)

func generatedRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/statement", controller.StatementList)
	makeRoute(r, http.MethodGet, "/statement/_new", controller.StatementCreateForm)
	makeRoute(r, http.MethodPost, "/statement/_new", controller.StatementCreate)
	makeRoute(r, http.MethodGet, "/statement/_random", controller.StatementRandom)
	makeRoute(r, http.MethodGet, "/statement/{id}", controller.StatementDetail)
	makeRoute(r, http.MethodGet, "/statement/{id}/edit", controller.StatementEditForm)
	makeRoute(r, http.MethodPost, "/statement/{id}/edit", controller.StatementEdit)
	makeRoute(r, http.MethodGet, "/statement/{id}/delete", controller.StatementDelete)
	makeRoute(r, http.MethodGet, "/db", controller.ConnectionList)
	makeRoute(r, http.MethodGet, "/db/_new", controller.ConnectionCreateForm)
	makeRoute(r, http.MethodPost, "/db/_new", controller.ConnectionCreate)
	makeRoute(r, http.MethodGet, "/db/_random", controller.ConnectionRandom)
	makeRoute(r, http.MethodGet, "/db/{id}", controller.ConnectionDetail)
	makeRoute(r, http.MethodGet, "/db/{id}/edit", controller.ConnectionEditForm)
	makeRoute(r, http.MethodPost, "/db/{id}/edit", controller.ConnectionEdit)
	makeRoute(r, http.MethodGet, "/db/{id}/delete", controller.ConnectionDelete)
}
