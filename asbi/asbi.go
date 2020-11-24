package asbi

import (
	"database/sql"
	"fmt"

	"github.com/columbio/api-gate/config"
	"github.com/columbio/api-gate/loggers"

	// import driver
	_ "github.com/lib/pq"
)

// Location - location description
type Location struct {
	ID             string `json:"ID"`
	OperatorID     string `json:"OperatorID"`
	OperatorName   string `json:"OperatorName"`
	OperatorStatus string `json:"OperatorStatus"`
	Address        string `json:"Address"`
	City           string `json:"City "`
	Region         string `json:"Region"`
	District       string `json:"District"`
	Status         string `json:"Status"`
	Latitude       string `json:"Latitude"`
	Longitude      string `json:"Longitude"`
}

// GetLocationByID - method to get location name by id
func GetLocationByID(locationID string) (*Location, error) {
	loggers.Info.Printf("GetLocationByID() starts")
	//
	DBconnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DaemonConfiguration.USERNAME,
		config.DaemonConfiguration.PASSWORD,
		config.DaemonConfiguration.DBHOSTIP,
		config.DaemonConfiguration.PORT,
		config.DaemonConfiguration.DBNAME)

	loggers.Info.Printf("GetLocationByID() with db conn = %s", DBconnection)
	// new Location object
	var location Location
	location.ID = "null"
	location.OperatorID = "null"
	location.OperatorName = "null"
	location.OperatorStatus = "null"
	location.Address = "null"
	location.City = "null"
	location.Region = "null"
	location.District = "null"
	location.Status = "null"
	location.Latitude = "null"
	location.Longitude = "null"
	// Open connnection
	loggers.Info.Printf("asbiPgsqlQuery() Setup DB connection")
	// Create DB pool
	db, err := sql.Open("postgres", DBconnection)
	defer db.Close()
	if err != nil {
		loggers.Error.Printf("asbiPgsqlQuery(): Failed to open a DB connection: %v", err)
		return &location, err
	}
	// Location Query SQL request
	// SQL query for selecting all field for Location
	locationquery := "SELECT * from " + config.DaemonConfiguration.SCHEME + ".dict_site WHERE id = '" + locationID + "';"
	err = db.QueryRow(locationquery).Scan(&location.ID, &location.OperatorID, &location.Address, &location.City, &location.Region, &location.District, &location.Status)
	if err != nil {
		loggers.Error.Printf("asbiPgsqlQuery(): Failed to execute location query: %v", err)
	}
	// Operator Query SQL request
	// SQL query for selecting all field for Location
	operatorquery := "SELECT name, disabled FROM " + config.DaemonConfiguration.SCHEME + ".dict_operator WHERE id = '" + location.OperatorID + "';"
	err = db.QueryRow(operatorquery).Scan(&location.OperatorName, &location.OperatorStatus)
	if err != nil {
		loggers.Error.Printf("asbiPgsqlQuery(): Failed to execute operator query: %v", err)
	}
	// return values
	return &location, err
}
