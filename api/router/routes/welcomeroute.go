package routes

import (
	"github.com/vanilla/go-mux-crud/api/controller"
	"net/http"
)

var welcomeRoute = []Route{
	Route{
		URI: "/",
		Method: http.MethodGet,
		Handler: controller.Welcome,
		AuthRequired: false,
	},
}