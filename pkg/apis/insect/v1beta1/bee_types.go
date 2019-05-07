


package v1beta1

import (
	"log"
	"context"

	"k8s.io/apimachinery/pkg/runtime"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"k8s-p/pkg/apis/insect"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Bee
// +k8s:openapi-gen=true
// +resource:path=bees,strategy=BeeStrategy
type Bee struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BeeSpec   `json:"spec,omitempty"`
	Status BeeStatus `json:"status,omitempty"`
}

// BeeSpec defines the desired state of Bee
type BeeSpec struct {
}

// BeeStatus defines the observed state of Bee
type BeeStatus struct {
}

// Validate checks that an instance of Bee is well formed
func (BeeStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*insect.Bee)
	log.Printf("Validating fields for Bee %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Bee field values
func (BeeSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Bee)
	// set default field values here
	log.Printf("Defaulting fields for Bee %s\n", obj.Name)
}
