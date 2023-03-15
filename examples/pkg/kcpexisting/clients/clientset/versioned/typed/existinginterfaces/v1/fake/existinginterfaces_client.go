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

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"k8s.io/client-go/rest"

	existinginterfacesv1 "acme.corp/pkg/generated/clientset/versioned/typed/existinginterfaces/v1"
	kcpexistinginterfacesv1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/existinginterfaces/v1"
)

var _ kcpexistinginterfacesv1.ExistinginterfacesV1ClusterInterface = (*ExistinginterfacesV1ClusterClient)(nil)

type ExistinginterfacesV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *ExistinginterfacesV1ClusterClient) Cluster(clusterPath logicalcluster.Path) existinginterfacesv1.ExistinginterfacesV1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &ExistinginterfacesV1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *ExistinginterfacesV1ClusterClient) TestTypes() kcpexistinginterfacesv1.TestTypeClusterInterface {
	return &testTypesClusterClient{Fake: c.Fake}
}

func (c *ExistinginterfacesV1ClusterClient) ClusterTestTypes() kcpexistinginterfacesv1.ClusterTestTypeClusterInterface {
	return &clusterTestTypesClusterClient{Fake: c.Fake}
}

var _ existinginterfacesv1.ExistinginterfacesV1Interface = (*ExistinginterfacesV1Client)(nil)

type ExistinginterfacesV1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *ExistinginterfacesV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *ExistinginterfacesV1Client) TestTypes(namespace string) existinginterfacesv1.TestTypeInterface {
	return &testTypesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *ExistinginterfacesV1Client) ClusterTestTypes() existinginterfacesv1.ClusterTestTypeInterface {
	return &clusterTestTypesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
