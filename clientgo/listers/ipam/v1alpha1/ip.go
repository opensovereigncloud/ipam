// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ipam/api/ipam/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IPLister helps list IPs.
// All objects returned here must be treated as read-only.
type IPLister interface {
	// List lists all IPs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IP, err error)
	// IPs returns an object that can list and get IPs.
	IPs(namespace string) IPNamespaceLister
	IPListerExpansion
}

// iPLister implements the IPLister interface.
type iPLister struct {
	indexer cache.Indexer
}

// NewIPLister returns a new IPLister.
func NewIPLister(indexer cache.Indexer) IPLister {
	return &iPLister{indexer: indexer}
}

// List lists all IPs in the indexer.
func (s *iPLister) List(selector labels.Selector) (ret []*v1alpha1.IP, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IP))
	})
	return ret, err
}

// IPs returns an object that can list and get IPs.
func (s *iPLister) IPs(namespace string) IPNamespaceLister {
	return iPNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IPNamespaceLister helps list and get IPs.
// All objects returned here must be treated as read-only.
type IPNamespaceLister interface {
	// List lists all IPs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IP, err error)
	// Get retrieves the IP from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IP, error)
	IPNamespaceListerExpansion
}

// iPNamespaceLister implements the IPNamespaceLister
// interface.
type iPNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IPs in the indexer for a given namespace.
func (s iPNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IP, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IP))
	})
	return ret, err
}

// Get retrieves the IP from the indexer for a given namespace and name.
func (s iPNamespaceLister) Get(name string) (*v1alpha1.IP, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ip"), name)
	}
	return obj.(*v1alpha1.IP), nil
}
