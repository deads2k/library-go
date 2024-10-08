// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// InfrastructureLister helps list Infrastructures.
// All objects returned here must be treated as read-only.
type InfrastructureLister interface {
	// List lists all Infrastructures in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Infrastructure, err error)
	// Get retrieves the Infrastructure from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Infrastructure, error)
	InfrastructureListerExpansion
}

// infrastructureLister implements the InfrastructureLister interface.
type infrastructureLister struct {
	listers.ResourceIndexer[*v1.Infrastructure]
}

// NewInfrastructureLister returns a new InfrastructureLister.
func NewInfrastructureLister(indexer cache.Indexer) InfrastructureLister {
	return &infrastructureLister{listers.New[*v1.Infrastructure](indexer, v1.Resource("infrastructure"))}
}
