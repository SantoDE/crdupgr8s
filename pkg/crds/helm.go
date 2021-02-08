package crds

import (
	"io"
	"io/ioutil"
	"net/http"

	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
)

func DownloadChart(url string) (*chart.Chart, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	file, err := ioutil.TempFile("/tmp", "tar_")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	chartObj, err := loader.Load(file.Name())

	if err != nil {
		return nil, err
	}

	return chartObj, nil
}