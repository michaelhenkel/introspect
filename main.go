package main

import (
	"context"
	"fmt"

	"github.com/michaelhenkel/introspect/pkg/client/clientset"
	"github.com/michaelhenkel/introspect/pkg/rest"
)

func main() {
	restClient := rest.NewRESTClient("http")
	client := clientset.NewClientSet(restClient)

	/*
			bgpNeighbor, err := client.Control("192.168.64.4:8083").BgpNeighbors().List(context.Background())
			if err != nil {
				panic(err)
			}
			bgpNeighbors := bgpNeighbor.Neighbors.List.BgpNeighborResp
			for _, r := range bgpNeighbors {
				fmt.Println(r.Peer.Text, r.State.Text)
			}


		ri, err := client.Control("192.168.64.4:8083").RoutingInstances().List(context.Background())
		if err != nil {
			panic(err)
		}
		for _, r := range ri.Instances.List.ShowRoutingInstance {

			fmt.Println(r.Name.Text)
		}

		result, err := yaml.Marshal(ri)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(result))

		for _, r := range ri.Instances.List.ShowRoutingInstance {

			fmt.Println(r.Name.Text)
		}
	*/

	vn, err := client.Data("192.168.64.7:8085").VirtualNetworks().List(context.Background())
	if err != nil {
		panic(err)
	}
	for _, bla := range vn.VnListResp.VnList.List.VnSandeshData {
		fmt.Println(bla)
	}

	vmi, err := client.Data("192.168.64.7:8085").VirtualMachineInterfaces().List(context.Background())
	if err != nil {
		panic(err)
	}
	for _, bla := range vmi.ItfResp.ItfList.List.ItfSandeshData {
		fmt.Println(bla)
	}
}
