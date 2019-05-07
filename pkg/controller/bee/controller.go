


package bee

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/builders"

	"k8s-p/pkg/apis/insect/v1beta1"
	"k8s-p/pkg/controller/sharedinformers"
	listers "k8s-p/pkg/client/listers_generated/insect/v1beta1"
)

// +controller:group=insect,version=v1beta1,kind=Bee,resource=bees
type BeeControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Bee
	lister listers.BeeLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *BeeControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing bees labels
	c.lister = arguments.GetSharedInformers().Factory.Insect().V1beta1().Bees().Lister()
}

// Reconcile handles enqueued messages
func (c *BeeControllerImpl) Reconcile(u *v1beta1.Bee) error {
	// Implement controller logic here
	log.Printf("Running reconcile Bee for %s\n", u.Name)
	return nil
}

func (c *BeeControllerImpl) Get(namespace, name string) (*v1beta1.Bee, error) {
	return c.lister.Bees(namespace).Get(name)
}
