package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/DropOrg/api/api"
)

const version = "0.1.0"

type config struct {
	DBAddr string
	Env    string
}

var opts config

func initConfig() {
	// Defaults
	viper.SetDefault("db.addr", "mongodb://localhost:27017")
	viper.SetDefault("port", "8000")
	viper.RegisterAlias("DBAddr", "db.addr")
	viper.SetDefault("env", "dev")

	// Config File
	configPaths := []string{".", "$HOME/.drop/", "/etc/drop/"}
	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	// Env variables
	// FIXME for some reason this doesn't seem to work, I've tried testing it
	viper.SetEnvPrefix("drop")
	viper.AutomaticEnv()

	// Arguments
	// TODO. Once flags are supported.

}

func flags() {
	// TODO: Add support for flags using Cobra: https://github.com/spf13/cobra
}

func main() {
	fmt.Printf("Drop server v%s starting up!\n", version)

	// CLI params and flags
	flags()

	// Configuration
	initConfig()

	// Connect to mongo
	// client, err := mongo.NewClient(viper.GetString("db.addr"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	r := mux.NewRouter()
	r.HandleFunc("/", hello).
		Methods("GET")

	api.UsersRoute(r.PathPrefix("/users"))
	api.UserRoute(r.PathPrefix("/user"))

	fmt.Printf("Listening on localhost:%d\n", viper.GetInt("port"))
	http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("port")), r)
}

func hello(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(r.RequestURI)
}
