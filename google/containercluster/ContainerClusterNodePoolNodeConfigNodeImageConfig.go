// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNodeConfigNodeImageConfig struct {
	// The Operating System image for the node pool.
	//
	// This is a private feature, please contact your Google account team for allowlisting this feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/container_cluster#image ContainerCluster#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// The GCP project storing the Operating System image for the node pool.
	//
	// This is a private feature, please contact your Google account team for allowlisting this feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/container_cluster#image_project ContainerCluster#image_project}
	ImageProject *string `field:"optional" json:"imageProject" yaml:"imageProject"`
}

