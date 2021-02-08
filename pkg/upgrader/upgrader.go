package upgrader

import (
	"context"
	"log"
	"time"

	"github.com/SantoDE/crdupgr8s/pkg/crds"
	"github.com/SantoDE/crdupgr8s/pkg/k8s"
)

func UpgradeCRDS(url string) {
	chart, err := crds.DownloadChart(url)

	if err != nil {
		log.Fatalf("Can not download Helm chart: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	existingCrds, err := k8s.ListCRDS(ctx)

	if err != nil {
		log.Fatalf("Can not detect current existing CRDs in the cluster: %s", err.Error())
	}

	crdObjs := chart.CRDObjects()

	log.Printf("Detected %d CRDs in cluster and %d in the Helm Chart", len(existingCrds), len(crdObjs))

	for _, crd := range crdObjs {
		obj, err := k8s.ParseToCRD(crd.File.Data)

		if err != nil {
			log.Fatalf("Can not parse CRDs: %s", err.Error())
		}

		if !existingCrds.IncludesItem(obj.Name, obj.Spec.Group) {
			log.Printf("Detected that the CRD %s is missing in the Cluster, creating it now", obj.Name)
			k8s.Create(ctx, obj)
		}
	}

}