package controller

import (
	"github.com/DaviAraujoCC/dns-manager-gke/k8s/auth"
	"k8s.io/client-go/kubernetes"
)

type ObjectsController struct {
	*kubernetes.Clientset
	Namespace string
}

func NewObjectsController(n string) (*ObjectsController, error) {
	clientset, err := auth.NewClient()
	if err != nil {
		return nil, err
	}
	return &ObjectsController{
		clientset,
		n,
	}, nil
}
