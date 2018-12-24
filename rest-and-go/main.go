package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-and-go/login"

	"github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port              int
	Static_Variable   string
	Connection_String string
}

// func loadConfiguration(filename string) (Configuration, error) {
// 	var config Configuration
// 	filepath := "./config/config.json"
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		return config, err
// 	}
// 	decoder := json.NewDecoder(file)
// 	err = decoder.Decode(config)
// 	if err != nil {
// 		return config, err
// 	}
// 	return config, err
// }
type App struct {
	Router *mux.Router
}

func main() {
	// router := mux.NewRouter()
	// router.HandleFunc("/login", login).Methods("POST")
	// log.Fatal(http.ListenAndServe("127.0.0.1:8000", router))
	var a App
	a.Serve()
	a.Run()

}

func (a *App) Run() {
	configuration := Configuration{}
	err := gonfig.GetConf("config/config.json", &configuration)
	if err != nil {
		fmt.Println(err)
	}

	// cp := string(configuration.Port)
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(allowedOrigins, allowedMethods)(a.Router)))
}

func (a *App) Serve() {
	a.Router = login.NewRouter()
}
