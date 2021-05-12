package clientset

import (
	v1alpha1Control "github.com/michaelhenkel/introspect/pkg/client/clientset/typed/control/v1alpha1"
	v1alpha1Data "github.com/michaelhenkel/introspect/pkg/client/clientset/typed/data/v1alpha1"
	"github.com/michaelhenkel/introspect/pkg/rest"
)

type Interface interface {
	ControlV1alpha1() v1alpha1Control.ControlV1alpha1Interface
	DataV1alpha1() v1alpha1Data.DataV1alpha1Interface
}

type Clientset struct {
	controlV1alpha1 *v1alpha1Control.ControlplaneV1alpha1Client
	dataV1alpha1    *v1alpha1Data.DataplaneV1alpha1Client
}

func NewClientSet(restClient rest.Interface) Clientset {
	return Clientset{
		controlV1alpha1: v1alpha1Control.New(restClient),
		dataV1alpha1:    v1alpha1Data.New(restClient),
	}
}

func (c *Clientset) Control(host string) v1alpha1Control.ControlV1alpha1Interface {
	c.controlV1alpha1.SetHost(host)
	return c.controlV1alpha1
}

func (c *Clientset) Data(host string) v1alpha1Data.DataV1alpha1Interface {
	c.dataV1alpha1.SetHost(host)
	return c.dataV1alpha1
}
