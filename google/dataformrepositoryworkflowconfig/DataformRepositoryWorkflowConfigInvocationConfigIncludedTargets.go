// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataformrepositoryworkflowconfig


type DataformRepositoryWorkflowConfigInvocationConfigIncludedTargets struct {
	// The action's database (Google Cloud project ID).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_workflow_config#database DataformRepositoryWorkflowConfig#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The action's name, within database and schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_workflow_config#name DataformRepositoryWorkflowConfig#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The action's schema (BigQuery dataset ID), within database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_workflow_config#schema DataformRepositoryWorkflowConfig#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

