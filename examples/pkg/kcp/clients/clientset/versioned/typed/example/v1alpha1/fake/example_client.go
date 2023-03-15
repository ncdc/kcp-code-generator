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

	examplev1alpha1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1alpha1"
	kcpexamplev1alpha1 "acme.corp/pkg/kcp/clients/clientset/versioned/typed/example/v1alpha1"
)

var _ kcpexamplev1alpha1.ExampleV1alpha1ClusterInterface = (*ExampleV1alpha1ClusterClient)(nil)

type ExampleV1alpha1ClusterClient struct {
	*kcptesting.Fake
}

func (c *ExampleV1alpha1ClusterClient) Cluster(clusterPath logicalcluster.Path) examplev1alpha1.ExampleV1alpha1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &ExampleV1alpha1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *ExampleV1alpha1ClusterClient) TestTypes() kcpexamplev1alpha1.TestTypeClusterInterface {
	return &testTypesClusterClient{Fake: c.Fake}
}

func (c *ExampleV1alpha1ClusterClient) ClusterTestTypes() kcpexamplev1alpha1.ClusterTestTypeClusterInterface {
	return &clusterTestTypesClusterClient{Fake: c.Fake}
}

var _ examplev1alpha1.ExampleV1alpha1Interface = (*ExampleV1alpha1Client)(nil)

type ExampleV1alpha1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *ExampleV1alpha1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *ExampleV1alpha1Client) TestTypes(namespace string) examplev1alpha1.TestTypeInterface {
	return &testTypesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *ExampleV1alpha1Client) ClusterTestTypes() examplev1alpha1.ClusterTestTypeInterface {
	return &clusterTestTypesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
