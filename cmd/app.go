package cmd

import (
	"flag"

	"github.com/SantoDE/crdupgr8s/pkg/upgrader"
)

var urlFlag string

func init() {
	flag.StringVar(&urlFlag, "url", "", "Url to Helm Chart Release")
	flag.Parse()
}

func RunApp() {
	upgrader.UpgradeFromHelmChartUrl(urlFlag)
}
