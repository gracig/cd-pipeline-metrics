package influxdb

import (
	"time"

	"github.com/gracig/cd-runtime-collector/lib"

	"github.com/influxdata/influxdb/client/v2"
)

//NewMetricAdapter returns a new influxdb implementation of the metric.batchManager interface.
//It instantiates a new client.BatchPoints from the influxdb api
func NewMetricAdapter() (ma lib.MetricAdapter, err error) {

	var (
		bp  client.BatchPoints //Represent the object that will store the metrics
		c   client.Client      //Holds the Influxdb client connection
		cfg Config             //Influxdb Configuration
	)

	//Retrieve Influxdb configuration from environment
	if cfg, err = NewConfigFromEnv(); err != nil {
		return
	}

	//Use the inlfluxdb api to create a new BatchPoint interface. Return if err is not null
	if bp, err = client.NewBatchPoints(client.BatchPointsConfig{Database: cfg.DBName}); err != nil {
		return
	}

	//Initiates a connection to the Infludb Server API
	if c, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     cfg.APIURL,
		Username: cfg.APIUser,
		Password: cfg.APIPWD}); err != nil {
		return
	}

	//Instantiates a new datasource object
	ma = &metricAdapter{c: c, bp: bp}

	return
}

//batchManager implements the BatchManager interface using influxdb as backend
type metricAdapter struct {
	bp client.BatchPoints //The BatchPoints interface from Influxdb API
	c  client.Client      //Holds the Influxdb client connection
}

//AddDataPoint implements the method BatchManager.AddDataPoint using influxdb as backend
func (ma *metricAdapter) AddDataPoint(name string, tags map[string]string, fields map[string]interface{}, t time.Time) (err error) {

	var point *client.Point //a data Point for the infludb API

	//Create a new data point. Tags will be searchable fields and fields will hold values. The name will be the metric Name
	if point, err = client.NewPoint(name, tags, fields, t); err != nil {
		return
	}

	//Adds the new point into the batchpoints object.
	ma.bp.AddPoint(point)

	return
}

//WriteData implements the metric.Datasource.WriteData method using influxdb as backend
func (ma *metricAdapter) WriteData() (err error) {

	//Calls method write of influxdb client api using the Batchpoints
	err = ma.c.Write(ma.bp)

	return
}
