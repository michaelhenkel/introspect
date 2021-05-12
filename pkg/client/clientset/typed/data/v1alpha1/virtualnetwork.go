package v1alpha1

import (
	"context"
	"encoding/xml"
	"fmt"

	v1alpha1 "github.com/michaelhenkel/introspect/pkg/apis/data/v1alpha1"
	"github.com/michaelhenkel/introspect/pkg/rest"
)

type VirtualNetworksGetter interface {
	VirtualNetworks() VirtualNetworkInterface
}

type VirtualNetworkInterface interface {
	Get(ctx context.Context, name string) (*v1alpha1.VirtualNetwork, error)
	List(ctx context.Context) (*v1alpha1.VirtualNetwork, error)
}

type VirtualNetworks struct {
	client rest.Interface
}

func newVirtualNetworks(c *DataplaneV1alpha1Client) *VirtualNetworks {
	return &VirtualNetworks{
		client: c.restClient,
	}
}

func (c *VirtualNetworks) Get(ctx context.Context, name string) (result *v1alpha1.VirtualNetwork, err error) {
	result = &v1alpha1.VirtualNetwork{}
	path := fmt.Sprintf("Snh_VirtualNetworkReq?search_string=%s", name)
	resultBytes, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}
	if err := xml.Unmarshal(resultBytes, result); err != nil {
		return nil, err
	}
	return result, err
}

func (c *VirtualNetworks) List(ctx context.Context) (result *v1alpha1.VirtualNetwork, err error) {
	result = &v1alpha1.VirtualNetwork{}
	path := "Snh_VnListReq"
	resultBytes, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}
	if err := xml.Unmarshal(resultBytes, result); err != nil {
		return nil, err
	}
	return result, err
}
