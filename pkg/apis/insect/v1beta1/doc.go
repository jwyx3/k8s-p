


// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=k8s-practices/pkg/apis/insect
// +k8s:defaulter-gen=TypeMeta
// +groupName=insect.jw.io
package v1beta1 // import "k8s-practices/pkg/apis/insect/v1beta1"

