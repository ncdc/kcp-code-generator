//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package informers

import (
	"fmt"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"

	existinginterfacesv1 "github.com/kcp-dev/code-generator/examples/pkg/apis/existinginterfaces/v1"
	upstreaminformers "github.com/kcp-dev/code-generator/examples/pkg/generated/informers/externalversions"
)

type GenericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *GenericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericClusterLister.
func (f *GenericInformer) Lister() cache.GenericLister {
	return kcpcache.NewGenericClusterLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *SharedInformerFactory) ForResource(resource schema.GroupVersionResource) (upstreaminformers.GenericInformer, error) {
	switch resource {
	// Group=ExistingInterfaces, Version=v1
	case existinginterfacesv1.SchemeGroupVersion.WithResource("clustertesttypes"):
		return &GenericInformer{resource: resource.GroupResource(), informer: f.ExistingInterfaces().V1().ClusterTestTypes().Informer()}, nil
	case existinginterfacesv1.SchemeGroupVersion.WithResource("testtypes"):
		return &GenericInformer{resource: resource.GroupResource(), informer: f.ExistingInterfaces().V1().TestTypes().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}