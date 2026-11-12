// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataformrepositoryworkflowconfig


type DataformRepositoryWorkflowConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_workflow_config#create DataformRepositoryWorkflowConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_workflow_config#delete DataformRepositoryWorkflowConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_workflow_config#update DataformRepositoryWorkflowConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

