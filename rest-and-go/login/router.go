package login

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = Routes{
	Route{
		"Login",
		"POST",
		"/login",
		controller.Login,
	}, Route{
		"SignUp",
		"POST",
		"/signup",
		controller.SignUp,
	}, Route{
		"UpdatePassword",
		"POST",
		"/updatepassword",
		controller.UpdatePassword,
	}}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}
