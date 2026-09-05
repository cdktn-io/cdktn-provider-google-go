// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeConfigContainerdConfigRegistryHostsHostsHeader struct {
	// Configures the header key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/container_cluster#key ContainerCluster#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Configures the header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/container_cluster#value ContainerCluster#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

