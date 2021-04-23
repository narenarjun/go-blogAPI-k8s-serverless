package routes

import (
	"blogapi/api/controllers"
	"net/http"
)

var loginRoutes = []Route{
	Route{
		Uri:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
