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

		ri, err := client.Control("192.168.64.2:8083").RoutingInstances().Get(context.Background(), "default-domain", "contrail-k8s-kubemanager-mk-default-project", "contrail-k8s-kubemanager-mk-default-podnetwork", "contrail-k8s-kubemanager-mk-default-podnetwork")
		if err != nil {
			panic(err)
		}

		result, err := yaml.Marshal(ri)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(result))

	*/

	bgpNeighbor, err := client.Control("192.168.64.2:8083").BgpNeighbors().Get(context.Background(), "minikube")
	if err != nil {
		panic(err)
	}
	/*
		result, err := yaml.Marshal(bgpNeighbor)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(result))
		fmt.Println(bgpNeighbor.GetAddress())
		fmt.Println(bgpNeighbor.GetState())
	*/
	bgpNeighborRoutingInstance := bgpNeighbor.Neighbors.List.BgpNeighborResp.RoutingInstances.List.BgpNeighborRoutingInstance
	for _, r := range bgpNeighborRoutingInstance {
		fmt.Println("blabla", r.Name.Text, r.State.Text)
	}
	fmt.Println(bgpNeighbor)
}

// Client("192.168.1.1").RoutingInstances().Get("ri")
//clientset.ControlV1alpha1("1.1.1.1").RoutingInstances("").Get()
