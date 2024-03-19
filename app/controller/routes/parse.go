package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/dbaudit/app/controller"
)

func parseRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/parse", controller.ParseForm)
	makeRoute(r, http.MethodPost, "/parse", controller.Parse)
}
