package asbi

import (
	"database/sql"

	"github.com/duoflow/whois-email-resolver/loggers"
	// import driver
	_ "github.com/lib/pq"
)

const (
	// DBconnection - TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DBconnection = "postgres://postgres:VBXHYjXo@172.18.0.108:5432/events"
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
	loggers.Info.Printf("asbiGetLocationByID() starts")
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
	locationquery := "SELECT * from events_asbi_test.dict_site WHERE id = '" + locationID + "';"
	err = db.QueryRow(locationquery).Scan(&location.ID, &location.OperatorID, &location.Address, &location.City, &location.Region, &location.District, &location.Status)
	if err != nil {
		loggers.Error.Printf("asbiPgsqlQuery(): Failed to execute location query: %v", err)
	}
	// Operator Query SQL request
	// SQL query for selecting all field for Location
	operatorquery := "SELECT name, disabled FROM events_asbi_test.dict_operator WHERE id = '" + location.OperatorID + "';"
	err = db.QueryRow(operatorquery).Scan(&location.OperatorName, &location.OperatorStatus)
	if err != nil {
		loggers.Error.Printf("asbiPgsqlQuery(): Failed to execute operator query: %v", err)
	}
	// return values
	return &location, err
}
