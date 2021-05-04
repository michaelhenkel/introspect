package v1alpha1

import (
	"context"
	"encoding/xml"
	"fmt"

	v1alpha1 "github.com/michaelhenkel/introspect/pkg/apis/control/v1alpha1"
	"github.com/michaelhenkel/introspect/pkg/rest"
)

type BgpNeighborsGetter interface {
	BgpNeighbors() BgpNeighborInterface
}

type BgpNeighborInterface interface {
	Get(ctx context.Context, name string) (*v1alpha1.BgpNeighbor, error)
}

type BgpNeighbors struct {
	client rest.Interface
}

func newBgpNeighbors(c *ControlplaneV1alpha1Client) *BgpNeighbors {
	return &BgpNeighbors{
		client: c.restClient,
	}
}

func (c *BgpNeighbors) Get(ctx context.Context, name string) (result *v1alpha1.BgpNeighbor, err error) {
	result = &v1alpha1.BgpNeighbor{}
	path := fmt.Sprintf("Snh_BgpNeighborReq?search_string=%s", name)
	resultBytes, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}
	if err := xml.Unmarshal(resultBytes, result); err != nil {
		return nil, err
	}
	return result, err
}
