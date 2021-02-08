package k8s

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	crdClientSet "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/tools/clientcmd"
)

type clientWrapper struct {
	client *crdClientSet.Clientset
}

func NewClient() *clientWrapper {
	kubeConfigPath, err := findKubeConfig()

	if err != nil {
		log.Fatal(err)
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	crdClient, err := crdClientSet.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	return &clientWrapper{
		client: crdClient,
	}
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