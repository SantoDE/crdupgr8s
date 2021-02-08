package k8s

import (
	"os"

	crdClientSet "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/tools/clientcmd"
)

type clientWrapper struct {
	client *crdClientSet.Clientset
}

func NewClient() *clientWrapper {

	kubeConfigPath := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err.Error())
	}

	crdClient, err := crdClientSet.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	return &clientWrapper{
		client: crdClient,
	}
}