package k8s

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	crdClientSet "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type clientWrapper struct {
	client *crdClientSet.Clientset
}

func NewClient() (*clientWrapper, error) {
	var config *rest.Config

	if kubeconfig, err := findKubeConfig(); err == nil {
		config = createClusterExternalClient(kubeconfig)
	} else {
		config = createClusterInternalClient()
	}

	if config == nil {
		return nil, errors.New("Can not create a K8s client")
	}

	crdClient, err := crdClientSet.NewForConfig(config)

	if err != nil {
		return nil, err
	}

	return &clientWrapper{
		client: crdClient,
	}, nil
}

func createClusterInternalClient() (*rest.Config) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return config
}

func createClusterExternalClient(kubeConfig string) (*rest.Config) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return config
}

func findKubeConfig() (string, error) {
	env := os.Getenv("KUBECONFIG")
	if env != "" {
		return env, nil
	}

	if home := homedir.HomeDir(); home != "" {
		return filepath.Join(home, ".kube", "config"), nil
	}

	return "", errors.New("can not detect a kubeconfig file")
}