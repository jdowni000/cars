package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	checkclient "github.com/Comcast/kuberhealthy/pkg/checks/external/checkClient"
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
	resp, err := http.Get("car-server.batman.svc.cluster.local?car=Roadster")
	if err != nil {
		// checkclient.ReportFailure([]string{err.Error()})
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// checkclient.ReportFailure([]string{err.Error()})
		return err
	}
	stringbody := string(body)
	// scanner := bufio.NewScanner(stringbody)
	if strings.Contains(stringbody, "driving mile 0") {
		return nil
	}
	return errors.New("Did not find driving mile 0")
}
