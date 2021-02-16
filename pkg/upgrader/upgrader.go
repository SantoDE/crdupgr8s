package upgrader

import (
	"context"
	"log"
	"time"

	"github.com/SantoDE/crdupgr8s/pkg/k8s"
	"helm.sh/helm/v3/pkg/chart"
)

func upgrade(new []chart.CRD) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	existing, err := k8s.ListCRDS(ctx)

	if err != nil {
		log.Fatalf("Can not detect current existing CRDs in the cluster: %s", err.Error())
	}

	for _, crd := range new {
		obj, err := k8s.ParseToCRD(crd.File.Data)

		if err != nil {
			log.Fatalf("Can not parse CRDs: %s", err.Error())
		}

		if !existing.IncludesItem(obj.Name, obj.Spec.Group) {
			log.Printf("Detected that the CRD %s is missing in the Cluster, creating it now", obj.Name)
			k8s.Create(ctx, obj)
		}
	}
}