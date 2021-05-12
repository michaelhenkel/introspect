package v1alpha1

import (
	"context"
	"encoding/xml"
	"fmt"

	v1alpha1 "github.com/michaelhenkel/introspect/pkg/apis/data/v1alpha1"
	"github.com/michaelhenkel/introspect/pkg/rest"
)

type VirtualMachineInterfacesGetter interface {
	VirtualMachineInterfaces() VirtualMachineInterfaceInterface
}

type VirtualMachineInterfaceInterface interface {
	Get(ctx context.Context, name string) (*v1alpha1.VirtualMachineInterface, error)
	List(ctx context.Context) (*v1alpha1.VirtualMachineInterface, error)
}

type VirtualMachineInterfaces struct {
	client rest.Interface
}

func newVirtualMachineInterfaces(c *DataplaneV1alpha1Client) *VirtualMachineInterfaces {
	return &VirtualMachineInterfaces{
		client: c.restClient,
	}
}

func (c *VirtualMachineInterfaces) Get(ctx context.Context, name string) (result *v1alpha1.VirtualMachineInterface, err error) {
	result = &v1alpha1.VirtualMachineInterface{}
	path := fmt.Sprintf("Snh_ItfReq?search_string=%s", name)
	resultBytes, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}
	if err := xml.Unmarshal(resultBytes, result); err != nil {
		return nil, err
	}
	return result, err
}

func (c *VirtualMachineInterfaces) List(ctx context.Context) (result *v1alpha1.VirtualMachineInterface, err error) {
	result = &v1alpha1.VirtualMachineInterface{}
	path := "Snh_ItfReq"
	resultBytes, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}
	if err := xml.Unmarshal(resultBytes, result); err != nil {
		return nil, err
	}
	return result, err
}
