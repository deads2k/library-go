// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// ConsoleLister helps list Consoles.
// All objects returned here must be treated as read-only.
type ConsoleLister interface {
	// List lists all Consoles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*operatorv1.Console, err error)
	// Get retrieves the Console from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*operatorv1.Console, error)
	ConsoleListerExpansion
}

// consoleLister implements the ConsoleLister interface.
type consoleLister struct {
	listers.ResourceIndexer[*operatorv1.Console]
}

// NewConsoleLister returns a new ConsoleLister.
func NewConsoleLister(indexer cache.Indexer) ConsoleLister {
	return &consoleLister{listers.New[*operatorv1.Console](indexer, operatorv1.Resource("console"))}
}
