package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//WebGetter represents an HTML Loader that abstracts the Authentication step
type WebGetter interface {
	//DecodeRestJSON Retrieves HTML code and decodes into the data interface{}
	DecodeRestJSON(data interface{}, url string) error
}

//BasicAuthWebGetter Represents a Basic Authenticated Web Getter
type BasicAuthWebGetter struct {
	User string
	Pwd  string
}

//DecodeRestJSON implements the WebGetter.DecodeRestJSON method
func (w *BasicAuthWebGetter) DecodeRestJSON(data interface{}, url string) error {

	//Creates a new Http Client
	client := &http.Client{}

	//Creates a new request. Check for any errors
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	//Use Basic Authentication
	req.SetBasicAuth(w.User, w.Pwd)

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
