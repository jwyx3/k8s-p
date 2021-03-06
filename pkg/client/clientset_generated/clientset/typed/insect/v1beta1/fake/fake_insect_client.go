// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "k8s-p/pkg/client/clientset_generated/clientset/typed/insect/v1beta1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeInsectV1beta1 struct {
	*testing.Fake
}

func (c *FakeInsectV1beta1) Bees(namespace string) v1beta1.BeeInterface {
	return &FakeBees{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeInsectV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
