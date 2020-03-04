package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	checkclient "github.com/Comcast/kuberhealthy/v2/pkg/checks/external/checkclient"
)

// TargetURL retrieves URL that will be used to search for string in response body
var TargetURL = os.Getenv("TARGET_URL")

// TargetString is the string that will be searched for in the server response body
var TargetString = os.Getenv("TARGET_STRING")

// reportErrorAndStop reports to kuberhealthy of error and exits program when called
func reportErrorAndStop(s string) {
	log.Println(s)
	checkclient.ReportFailure([]string{s})
	os.Exit(1)
}

func main() {
	// check to make sure URL is provided
	if TargetURL == "" {
		reportErrorAndStop("No URL provided in YAML")
	}

	//check to make sure string is provided
	if TargetString == "" {
		reportErrorAndStop("No string provided in YAML")
	}

	userURLstring, err := getURLContent(TargetURL)
	if err != nil {
		log.Println(err)
		checkclient.ReportFailure([]string{err.Error()})
	}
	checkclient.ReportSuccess()

	if findStringInContent(userURLstring, TargetString) {
		checkclient.ReportSuccess()
	}
	log.Println(err)
	checkclient.ReportFailure([]string{err.Error()})

}

// getURLContent retrieves bytes and error from URL
func getURLContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err

	}
	defer resp.Body.Close()
	return body, err

}

// findStringInContent parses through URL bytes for specified string and returns bool
func findStringInContent(b []byte, s string) bool {

	stringbody := string(b)
	if strings.Contains(stringbody, s) {
		return true
	}
	return false
}
