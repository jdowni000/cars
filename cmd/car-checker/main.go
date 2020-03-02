package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	checkclient "github.com/Comcast/kuberhealthy/v2/pkg/checks/external/checkclient"
)

func main() {
	err := carChecker()
	if err != nil {
		checkclient.ReportFailure([]string{err.Error()})
		return
	}
	checkclient.ReportSuccess()
}

func carChecker() error {
	log.Println("starting check")
	resp, err := http.Get("http://car-server.batman.svc.cluster.local?car=Roadster")
	if err != nil {
		// checkclient.ReportFailure([]string{err.Error()})
		return err
	}
	defer resp.Body.Close()
	log.Println("reading body content")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// checkclient.ReportFailure([]string{err.Error()})
		return err
	}
	stringbody := string(body)
	// scanner := bufio.NewScanner(stringbody)
	if strings.Contains(stringbody, "driving mile 0") {
		log.Println("found expected string")
		return nil
	}
	log.Println("did not find expected string")

	return errors.New("Did not find driving mile 0")
}
