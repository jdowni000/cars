package main

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"strings"

	checkclient "github.com/Comcast/kuberhealthy/pkg/checks/external/checkClient"
)

func main() {
	ok := carChecker()
	if !ok {
		checkclient.ReportFailure([]string{"Test has failed!"})
		return
	}
	checkclient.ReportSuccess()
}

func carChecker() bool {
	resp, err := http.Get("car-server.batman.svc.cluster.local?car=Roadster")
	if err != nil {
		log.fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.fatal(err)
	}
	stringbody := strconv(body)
	scanner := bufio.NewScanner(stringbody)
	if strings.Contains(scanner.Text(), "driving mile 0") {
		return true
	}
	return false
}
