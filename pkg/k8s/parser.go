package k8s

import (
	"context"
	"fmt"
	"log"

	apiextensionsv1"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
    apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
)

type CrdList map[string]*apiextensionsv1.CustomResourceDefinition

type CrdListKey struct {
	name string
	group string
}

var client *clientWrapper

func init() {
	cl, err  := NewClient()

	if err != nil {
		log.Fatalf("Failed to initialize a Kubernetes Client to connect: %s", err.Error())
	}

	client = cl
}


func ParseToCRD(data []byte) (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	object := &apiextensionsv1beta1.CustomResourceDefinition{}

	decoder := serializer.NewCodecFactory(scheme.Scheme).UniversalDecoder()

	if err := runtime.DecodeInto(decoder, data, object); err != nil {
		return nil, err
	}

	return object, nil
}

func ListCRDS(ctx context.Context) (CrdList, error) {
	opts := metav1.ListOptions{}

	crds, err := client.client.ApiextensionsV1().CustomResourceDefinitions().List(ctx, opts)

	var crdList = newCRDList()

	for _, crd := range crds.Items {
		key := newCRDListKey(crd.Name, crd.Spec.Group)
		crdList[key.string()] = &crd
	}

	return crdList, err
}

func Create(ctx context.Context, def *apiextensionsv1beta1.CustomResourceDefinition) {
	opts := metav1.CreateOptions{}

	if _, err := client.client.ApiextensionsV1beta1().CustomResourceDefinitions().Create(ctx, def, opts); err != nil {
		log.Fatal(err)
	}

	log.Printf("Created CRD with name %s and group %s", def.Name, def.Spec.Group)

}

func newCRDList() CrdList {
	crdListMap := make(map[string]*apiextensionsv1.CustomResourceDefinition)
	return crdListMap
}

func newCRDListKey(name string, group string) *CrdListKey {
	return &CrdListKey{
		name,
		group,
	}
}

func (l CrdList) IncludesItem(name string, group string) bool {
	key := newCRDListKey(name, group)

	return l[key.string()] != nil
}

func (k CrdListKey)  string() string {
	return fmt.Sprintf("%s_%s", k.group, k.name)
}