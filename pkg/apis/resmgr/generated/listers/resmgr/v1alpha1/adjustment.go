// Copyright 2019-2020 Intel Corporation. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/intel/nri-resmgr/pkg/apis/resmgr/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AdjustmentLister helps list Adjustments.
// All objects returned here must be treated as read-only.
type AdjustmentLister interface {
	// List lists all Adjustments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Adjustment, err error)
	// Adjustments returns an object that can list and get Adjustments.
	Adjustments(namespace string) AdjustmentNamespaceLister
	AdjustmentListerExpansion
}

// adjustmentLister implements the AdjustmentLister interface.
type adjustmentLister struct {
	indexer cache.Indexer
}

// NewAdjustmentLister returns a new AdjustmentLister.
func NewAdjustmentLister(indexer cache.Indexer) AdjustmentLister {
	return &adjustmentLister{indexer: indexer}
}

// List lists all Adjustments in the indexer.
func (s *adjustmentLister) List(selector labels.Selector) (ret []*v1alpha1.Adjustment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Adjustment))
	})
	return ret, err
}

// Adjustments returns an object that can list and get Adjustments.
func (s *adjustmentLister) Adjustments(namespace string) AdjustmentNamespaceLister {
	return adjustmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AdjustmentNamespaceLister helps list and get Adjustments.
// All objects returned here must be treated as read-only.
type AdjustmentNamespaceLister interface {
	// List lists all Adjustments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Adjustment, err error)
	// Get retrieves the Adjustment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Adjustment, error)
	AdjustmentNamespaceListerExpansion
}

// adjustmentNamespaceLister implements the AdjustmentNamespaceLister
// interface.
type adjustmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Adjustments in the indexer for a given namespace.
func (s adjustmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Adjustment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Adjustment))
	})
	return ret, err
}

// Get retrieves the Adjustment from the indexer for a given namespace and name.
func (s adjustmentNamespaceLister) Get(name string) (*v1alpha1.Adjustment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("adjustment"), name)
	}
	return obj.(*v1alpha1.Adjustment), nil
}