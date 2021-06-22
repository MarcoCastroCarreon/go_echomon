package main

import (
	"os"

	echoconfig "github.com/MarcoCastroCarreon/go_echomon/src/echo_config"
	"github.com/MarcoCastroCarreon/go_echomon/src/mongo"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var connected bool
var connectionError error

func init() {
	loadError := godotenv.Load(".env")

	if loadError != nil {
		color.Red("Load Env File Error ->", loadError.Error())
	} else {
		mongoUri := os.Getenv("MONGO_URI")
		dbName := os.Getenv("DB_NAME")
		connectionStablished, err := mongo.InitConnection(mongoUri, dbName)
		connected = connectionStablished
		connectionError = err
	}
}

func main() {
	if connected {
		echoconfig.InitServer()
	} else {
		color.Red("DB Connection Error ->", connectionError.Error())
	}

}
