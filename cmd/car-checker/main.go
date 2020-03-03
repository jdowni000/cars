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

// main reports that the external checker has found problems and env is not empty
func main() {
	if TargetURL == "" {
		log.Println("No URL provided in YAML")
		os.Exit(1)
	}

	if TargetString == "" {
		log.Println("No string provided in YAML")
		os.Exit(1)
	}

	err := stringChecker()
	if err != nil {
		checkclient.ReportFailure([]string{err.Error()})
		return
	}
	checkclient.ReportSuccess()

}

// stringChecker	 reads through supplied URL for supplied string
func stringChecker() error {
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
