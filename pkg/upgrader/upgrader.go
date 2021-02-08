package upgrader

import (
	"log"

	"github.com/SantoDE/crdupgr8s/pkg/crds"
	"github.com/SantoDE/crdupgr8s/pkg/k8s"
)

func UpgradeCRDS(url string) {

	chart, err := crds.DownloadChart(url)

	if err != nil {
		log.Fatal(err.Error())
	}

	existingCrds, err := k8s.ListCRDS()
	crdObjs := chart.CRDObjects()

	log.Printf("Detected %d CRDs in cluster and %d in the Helm Chart", len(existingCrds), len(crdObjs))

	for _, crd := range crdObjs {
		obj := k8s.ParseToCRD(crd.File.Data)

		if !existingCrds.InludesItem(obj.Name, obj.Spec.Group) {
			log.Printf("Detected that the CRD %s is missing in the Cluster, creating it now", obj.Name)
			k8s.Create(obj)
		}
	}

}