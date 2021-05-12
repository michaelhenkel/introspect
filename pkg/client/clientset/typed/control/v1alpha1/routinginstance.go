package v1alpha1

import (
	"context"
	"encoding/xml"
	"fmt"

	v1alpha1 "github.com/michaelhenkel/introspect/pkg/apis/control/v1alpha1"
	"github.com/michaelhenkel/introspect/pkg/rest"
)

type RoutingInstancesGetter interface {
	RoutingInstances() RoutingInstanceInterface
}

type RoutingInstanceInterface interface {
	Get(ctx context.Context, domain, namespace, virtualNetwork, name string) (*v1alpha1.RoutingInstance, error)
	List(ctx context.Context) (*v1alpha1.RoutingInstance, error)
}

type routingInstances struct {
	client rest.Interface
}

func newRoutingInstances(c *ControlplaneV1alpha1Client) *routingInstances {
	return &routingInstances{
		client: c.restClient,
	}
}

func (c *routingInstances) Get(ctx context.Context, domain, namespace, virtualNetwork, name string) (result *v1alpha1.RoutingInstance, err error) {
	result = &v1alpha1.RoutingInstance{}
	path := fmt.Sprintf("Snh_ShowRoutingInstanceReq?search_string=%s:%s:%s:%s", domain, namespace, virtualNetwork, name)
	resultBytes, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}
	if err := xml.Unmarshal(resultBytes, result); err != nil {
		return nil, err
	}
	return result, err
}

func (c *routingInstances) List(ctx context.Context) (result *v1alpha1.RoutingInstance, err error) {
	result = &v1alpha1.RoutingInstance{}
	path := "Snh_ShowRoutingInstanceReq"
	resultBytes, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}
	if err := xml.Unmarshal(resultBytes, result); err != nil {
		return nil, err
	}
	return result, err
}
