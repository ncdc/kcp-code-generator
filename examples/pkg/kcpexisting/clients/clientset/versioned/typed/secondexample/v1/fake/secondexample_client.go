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

package v1

import (
	"github.com/kcp-dev/logicalcluster/v2"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"k8s.io/client-go/rest"

	secondexamplev1 "acme.corp/pkg/generated/clientset/versioned/typed/secondexample/v1"
	kcpsecondexamplev1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/secondexample/v1"
)

var _ kcpsecondexamplev1.SecondexampleV1ClusterInterface = (*SecondexampleV1ClusterClient)(nil)

type SecondexampleV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *SecondexampleV1ClusterClient) Cluster(cluster logicalcluster.Name) secondexamplev1.SecondexampleV1Interface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &SecondexampleV1Client{Fake: c.Fake, Cluster: cluster}
}

func (c *SecondexampleV1ClusterClient) TestTypes() kcpsecondexamplev1.TestTypeClusterInterface {
	return &testTypesClusterClient{Fake: c.Fake}
}

func (c *SecondexampleV1ClusterClient) ClusterTestTypes() kcpsecondexamplev1.ClusterTestTypeClusterInterface {
	return &clusterTestTypesClusterClient{Fake: c.Fake}
}

var _ secondexamplev1.SecondexampleV1Interface = (*SecondexampleV1Client)(nil)

type SecondexampleV1Client struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *SecondexampleV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *SecondexampleV1Client) TestTypes(namespace string) secondexamplev1.TestTypeInterface {
	return &testTypesClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *SecondexampleV1Client) ClusterTestTypes() secondexamplev1.ClusterTestTypeInterface {
	return &clusterTestTypesClient{Fake: c.Fake, Cluster: c.Cluster}
}