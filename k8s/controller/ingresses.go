package controller

import (
	"context"

	v1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *ObjectsController) ListIngresses() (*v1beta1.IngressList, error) {
	ingresses, err := c.ExtensionsV1beta1().Ingresses(c.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return ingresses, nil
}
