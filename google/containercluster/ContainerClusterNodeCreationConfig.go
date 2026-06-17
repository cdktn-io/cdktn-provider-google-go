// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeCreationConfig struct {
	// NodeCreationMode defines the settings of node creation mode.
	//
	// Accepted values are:
	// * VIA_KUBELET: Kubelet registers itself.
	// * VIA_CONTROL_PLANE: gcp-controller-manager automatically creates the node object after CSR approval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/container_cluster#node_creation_mode ContainerCluster#node_creation_mode}
	NodeCreationMode *string `field:"required" json:"nodeCreationMode" yaml:"nodeCreationMode"`
}

