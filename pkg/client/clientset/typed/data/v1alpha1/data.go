package v1alpha1

import (
	"github.com/michaelhenkel/introspect/pkg/rest"
)

type DataV1alpha1Interface interface {
	VirtualNetworksGetter
	VirtualMachineInterfacesGetter
	RoutingInstancesGetter
}

type DataplaneV1alpha1Client struct {
	restClient rest.Interface
}

func (c *DataplaneV1alpha1Client) SetHost(host string) {
	c.restClient.SetHost(host)
}

func (c *DataplaneV1alpha1Client) VirtualNetworks() VirtualNetworkInterface {
	return newVirtualNetworks(c)
}

func (c *DataplaneV1alpha1Client) VirtualMachineInterfaces() VirtualMachineInterfaceInterface {
	return newVirtualMachineInterfaces(c)
}

func (c *DataplaneV1alpha1Client) RoutingInstances() RoutingInstanceInterface {
	return newRoutingInstances(c)
}

func New(c rest.Interface) *DataplaneV1alpha1Client {
	return &DataplaneV1alpha1Client{c}
}
