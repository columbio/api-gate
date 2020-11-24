package config

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/columbio/api-gate/loggers"
)

var (
	// DaemonConfiguration - global configuration register
	DaemonConfiguration Configuration

	// RandomGen - random generator
	RandomGen = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Configuration - struct to load config
// from a json file
type Configuration struct {
	DBHOSTIP string `json:"DBHOSTIP"`
	PORT     string `json:"PORT"`
	USERNAME string `json:"USERNAME"`
	PASSWORD string `json:"PASSWORD"`
	DBNAME   string `json:"DBNAME"`
	SCHEME   string `json:"SCHEME"`
}

// ReadConfig - function to read config from yaml file
func ReadConfig() error {
	var configuration Configuration
	//
	loggers.Info.Printf("ReadConfig() starts")
	// --------
	// Open main config file
	configFile, err1 := os.Open("./api-gate.conf.json")
	defer configFile.Close()

	// if we os.Open returns an error then handle it
	if err1 != nil {
		loggers.Error.Println(err1)
		os.Exit(1)
	}
	loggers.Info.Printf("ReadConfig() Successfully opened сonfig files")
	// read our opened xmlFile as a byte array.
	configValue, _ := ioutil.ReadAll(configFile)
	// we initialize our Users array
	parseerr := json.Unmarshal(configValue, &configuration)
	if parseerr != nil {
		loggers.Error.Println(parseerr)
		os.Exit(1)
	}
	// --------
	DaemonConfiguration = configuration
	loggers.Info.Println(DaemonConfiguration)
	return nil
}
