


package main

import (
	// Make sure dep tools picks up these dependencies
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "github.com/go-openapi/loads"

	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/cmd/server"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // Enable cloud provider auth

	"k8s-practices/pkg/apis"
	"k8s-practices/pkg/openapi"
)

func main() {
	version := "v0"
	server.StartApiServer("/registry/jw.io", apis.GetAllApiBuilders(), openapi.GetOpenAPIDefinitions, "Api", version)
}
