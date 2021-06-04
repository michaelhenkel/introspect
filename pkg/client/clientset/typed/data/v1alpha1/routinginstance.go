package v1alpha1

import (
	"context"
	"encoding/xml"
	"fmt"

	v1alpha1 "github.com/michaelhenkel/introspect/pkg/apis/data/v1alpha1"
	"github.com/michaelhenkel/introspect/pkg/rest"
)

type RoutingInstancesGetter interface {
	RoutingInstances() RoutingInstanceInterface
}

type RoutingInstanceInterface interface {
	Get(ctx context.Context, name string) (*v1alpha1.RoutingInstance, error)
	List(ctx context.Context) (*v1alpha1.RoutingInstance, error)
}

type RoutingInstances struct {
	client rest.Interface
}

func newRoutingInstances(c *DataplaneV1alpha1Client) *RoutingInstances {
	return &RoutingInstances{
		client: c.restClient,
	}
}

func (c *RoutingInstances) Get(ctx context.Context, name string) (result *v1alpha1.RoutingInstance, err error) {
	result = &v1alpha1.RoutingInstance{}
	path := fmt.Sprintf("Snh_RoutingInstanceReq?search_string=%s", name)
	resultBytes, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}
	if err := xml.Unmarshal(resultBytes, result); err != nil {
		return nil, err
	}
	return result, err
}

func (c *RoutingInstances) List(ctx context.Context) (result *v1alpha1.RoutingInstance, err error) {
	result = &v1alpha1.RoutingInstance{}
	path := "Snh_VrfListReq"
	resultBytes, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}
	if err := xml.Unmarshal(resultBytes, result); err != nil {
		return nil, err
	}
	return result, err
}
