/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by lister-gen

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "operator/pkg/api/quobyte.com/v1"
)

// QuobyteClientLister helps list QuobyteClients.
type QuobyteClientLister interface {
	// List lists all QuobyteClients in the indexer.
	List(selector labels.Selector) (ret []*v1.QuobyteClient, err error)
	// QuobyteClients returns an object that can list and get QuobyteClients.
	QuobyteClients(namespace string) QuobyteClientNamespaceLister
	QuobyteClientListerExpansion
}

// quobyteClientLister implements the QuobyteClientLister interface.
type quobyteClientLister struct {
	indexer cache.Indexer
}

// NewQuobyteClientLister returns a new QuobyteClientLister.
func NewQuobyteClientLister(indexer cache.Indexer) QuobyteClientLister {
	return &quobyteClientLister{indexer: indexer}
}

// List lists all QuobyteClients in the indexer.
func (s *quobyteClientLister) List(selector labels.Selector) (ret []*v1.QuobyteClient, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.QuobyteClient))
	})
	return ret, err
}

// QuobyteClients returns an object that can list and get QuobyteClients.
func (s *quobyteClientLister) QuobyteClients(namespace string) QuobyteClientNamespaceLister {
	return quobyteClientNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// QuobyteClientNamespaceLister helps list and get QuobyteClients.
type QuobyteClientNamespaceLister interface {
	// List lists all QuobyteClients in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.QuobyteClient, err error)
	// Get retrieves the QuobyteClient from the indexer for a given namespace and name.
	Get(name string) (*v1.QuobyteClient, error)
	QuobyteClientNamespaceListerExpansion
}

// quobyteClientNamespaceLister implements the QuobyteClientNamespaceLister
// interface.
type quobyteClientNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all QuobyteClients in the indexer for a given namespace.
func (s quobyteClientNamespaceLister) List(selector labels.Selector) (ret []*v1.QuobyteClient, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.QuobyteClient))
	})
	return ret, err
}

// Get retrieves the QuobyteClient from the indexer for a given namespace and name.
func (s quobyteClientNamespaceLister) Get(name string) (*v1.QuobyteClient, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("quobyteclient"), name)
	}
	return obj.(*v1.QuobyteClient), nil
}
