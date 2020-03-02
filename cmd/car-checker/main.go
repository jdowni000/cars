package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	checkclient "github.com/Comcast/kuberhealthy/v2/pkg/checks/external/checkclient"
)

// TargetURL variable set from YAML env
var TargetURL = os.Getenv("TARGET_URL")

//TargetString variable set from YAML env
var TargetString = os.Getenv("TARGET_STRING")

// init verifies env are not empty from YAML
func init() {
	if TargetURL == "" {
		log.Println("No URL provided in YAML")
		return
	}
	if TargetString == "" {

		log.Println("No string provided in YAML")
		return
	}
}

// main reports that the external checker has found problems
func main() {
	err := carChecker()
	if err != nil {
		checkclient.ReportFailure([]string{err.Error()})
		return
	}
	checkclient.ReportSuccess()
}

// carChecker	 reads through supplied URL for supplied string
func carChecker() error {
	log.Println("starting check")
	resp, err := http.Get(TargetURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Println("reading body content")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	stringbody := string(body)
	if strings.Contains(stringbody, TargetString) {
		log.Println("found expected string")
		return nil
	}
	log.Println("did not find expected string")

	return errors.New("Did not find wanted string")
}
