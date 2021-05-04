package v1alpha1

import (
	"github.com/michaelhenkel/introspect/pkg/rest"
)

type ControlV1alpha1Interface interface {
	RoutingInstancesGetter
	BgpNeighborsGetter
}

type ControlplaneV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ControlplaneV1alpha1Client) SetHost(host string) {
	c.restClient.SetHost(host)
}

func (c *ControlplaneV1alpha1Client) RoutingInstances() RoutingInstanceInterface {
	return newRoutingInstances(c)
}

func (c *ControlplaneV1alpha1Client) BgpNeighbors() BgpNeighborInterface {
	return newBgpNeighbors(c)
}

func New(c rest.Interface) *ControlplaneV1alpha1Client {
	return &ControlplaneV1alpha1Client{c}
}
