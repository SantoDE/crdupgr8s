package upgrader

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
)

func UpgradeFromHelmChartUrl(url string) {
	chart, err := downloadChart(url)

	if err != nil {
		log.Fatalf("Can not download Helm chart: %s", err.Error())
	}

	crdObjs := chart.CRDObjects()

	upgrade(crdObjs)
}


func downloadChart(url string) (*chart.Chart, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	file, err := ioutil.TempFile("/tmp", "chart_")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		return nil, err
	}

	chartObj, err := loader.Load(file.Name())

	if err != nil {
		return nil, err
	}

	return chartObj, nil
}