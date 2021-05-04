package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Interface interface {
	Post()
	Get(path string) ([]byte, error)
	SetHost(string)
}

type RESTClient struct {
	base   *url.URL
	Client *http.Client
}

func (r *RESTClient) Post() {

}

func (r *RESTClient) SetHost(host string) {
	r.base.Host = host
}

func (r *RESTClient) SetBase(url *url.URL) {
	r.base = url
}

func (r *RESTClient) Get(path string) ([]byte, error) {
	url := fmt.Sprintf("%s://%s/%s", r.base.Scheme, r.base.Host, path)
	result, err := getXML(url)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewRESTClient(scheme string) *RESTClient {
	return &RESTClient{
		base: &url.URL{
			Scheme: scheme,
		},
	}
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("STATUS ERROR: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("READ BODY: %v", err)
	}

	return data, nil
}

//http://192.168.64.2:8083/Snh_ShowRoutingInstanceReq?search_string=default-domain%3Acontrail-k8s-kubemanager-mk-default-project%3Acontrail-k8s-kubemanager-mk-default-podnetwork%3Acontrail-k8s-kubemanager-mk-default-podnetwork
