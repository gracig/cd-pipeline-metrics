package influxdb

import "github.com/gracig/cd-runtime-collector/lib"

//Config represents parameters to access or manages an InfluxDB API
type Config struct {
	//APIPWD is the token password of the InfluxDB API
	APIPWD string
	//APIUser is the user of the InfluxDB API
	APIUser string
	//APIURL is the url where the InfluxDB API is hosted
	APIURL string
	//DBName is the database Name to read or write variables to
	DBName string
}

//NewConfigFromEnv creates a new InfluxDBConfig from environment variables.
//Expected environment variables: INFLUXDB_PWD INFLUXDB_USER INFLUXDB_URL INFLUXDB_DB
func NewConfigFromEnv() (cfg Config, err error) {

	if cfg.APIPWD, err = lib.GetEnvVar("INFLUXDB_PWD", ".+"); err != nil {
		return
	}
	if cfg.APIUser, err = lib.GetEnvVar("INFLUXDB_USER", ".+"); err != nil {
		return
	}
	if cfg.APIURL, err = lib.GetEnvVar("INFLUXDB_URL", "http://.*:\\d*"); err != nil {
		return
	}
	if cfg.DBName, err = lib.GetEnvVar("INFLUXDB_DB", ".+"); err != nil {
		return
	}
	return
}
