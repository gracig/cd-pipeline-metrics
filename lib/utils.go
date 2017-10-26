package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

//GetEnvVar retrieves an environment variable and validate it against a regular expression
func GetEnvVar(name string, validationRegex string) (value string, err error) {
	value = os.Getenv(name)
	if match, _ := regexp.MatchString(validationRegex, value); !match {
		err = fmt.Errorf("Environment variable [%v] has an invalid value [%v]", name, value)
	}
	return
}

//GetEnvVarAsInt64 retrieves an environment variable as int64
func GetEnvVarAsInt64(name string) (value int64, err error) {

	//Holds the environment variable value
	var evalue string

	//Assigns environment variable into evalue. returns if an error occurred
	if evalue, err = GetEnvVar(name, "\\d+"); err != nil {
		return
	}

	//Convert evalue into value using a decimal number with 64 bits long
	value, err = strconv.ParseInt(evalue, 10, 64)

	return
}

//GetEnvVarAsFloat64 retrieves an environment variable as float64
func GetEnvVarAsFloat64(name string) (value float64, err error) {

	//Holds the environment variable value
	var evalue string

	//Assigns environment variable into evalue. returns if an error occurred
	if evalue, err = GetEnvVar(name, "\\d+"); err != nil {
		return
	}

	//Convert evalue into a float 64 bits number
	value, err = strconv.ParseFloat(evalue, 64)

	return
}

//Text2Float returns a text as a float value inside an Interface, nil if text is invalid
func Text2Float(from string) (value interface{}) {
	var err error
	if value, err = strconv.ParseFloat(from, 64); err != nil {
		log.Printf("Could not parse float value of %v", from)
	}
	return
}

//DecodeRestJSON retrieves the JSON content of a URL and decodes it into the data interface parameter
func DecodeRestJSON(data interface{}, url string, user string, pwd string) error {

	//Creates a new Http Client
	client := &http.Client{}

	//Creates a new request. Check for any errors
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	//Use Basic Authentication
	req.SetBasicAuth(user, pwd)

	//Do the Request. Check for any errors
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	//Tells to close body message after this function ends.
	defer resp.Body.Close()

	//Verifies if url returns a status other then 200. Return the error
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body) // Tries to read the error body

		return fmt.Errorf("Error: [%v]\nResponse: [%v]\nURL: [%v]", resp.Status, string(body), url)
	}

	//Tries do decode the body into the data interface
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return err
	}

	//Returns no errors
	return nil
}
