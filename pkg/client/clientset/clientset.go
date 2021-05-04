package clientset

import (
	v1alpha1Control "github.com/michaelhenkel/introspect/pkg/client/clientset/typed/control/v1alpha1"
	"github.com/michaelhenkel/introspect/pkg/rest"
)

type Interface interface {
	ControlV1alpha1() v1alpha1Control.ControlV1alpha1Interface
}

type Clientset struct {
	controlV1alpha1 *v1alpha1Control.ControlplaneV1alpha1Client
}

func NewClientSet(restClient rest.Interface) Clientset {
	return Clientset{
		controlV1alpha1: v1alpha1Control.New(restClient),
	}
}

func (c *Clientset) Control(host string) v1alpha1Control.ControlV1alpha1Interface {
	c.controlV1alpha1.SetHost(host)
	return c.controlV1alpha1
}
