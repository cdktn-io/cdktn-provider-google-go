// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataformrepositoryreleaseconfig


type DataformRepositoryReleaseConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_release_config#create DataformRepositoryReleaseConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_release_config#delete DataformRepositoryReleaseConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_release_config#update DataformRepositoryReleaseConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

