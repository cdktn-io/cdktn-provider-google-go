// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeConfigNodeImageConfig struct {
	// The name of the image to use for this node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/container_cluster#image ContainerCluster#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// The project containing the image to use for this node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/container_cluster#image_project ContainerCluster#image_project}
	ImageProject *string `field:"optional" json:"imageProject" yaml:"imageProject"`
}

