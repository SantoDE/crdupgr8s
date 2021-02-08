package crds

import (
	"io"
	"io/ioutil"
	"net/http"

	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
)

func DownloadChart(url string) (*chart.Chart, error) {

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