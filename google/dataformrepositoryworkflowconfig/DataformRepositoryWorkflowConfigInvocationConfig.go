// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataformrepositoryworkflowconfig


type DataformRepositoryWorkflowConfigInvocationConfig struct {
	// Optional. When set to true, any incremental tables will be fully refreshed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_workflow_config#fully_refresh_incremental_tables_enabled DataformRepositoryWorkflowConfig#fully_refresh_incremental_tables_enabled}
	FullyRefreshIncrementalTablesEnabled interface{} `field:"optional" json:"fullyRefreshIncrementalTablesEnabled" yaml:"fullyRefreshIncrementalTablesEnabled"`
	// Optional. The set of tags to include.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_workflow_config#included_tags DataformRepositoryWorkflowConfig#included_tags}
	IncludedTags *[]*string `field:"optional" json:"includedTags" yaml:"includedTags"`
	// included_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_workflow_config#included_targets DataformRepositoryWorkflowConfig#included_targets}
	IncludedTargets interface{} `field:"optional" json:"includedTargets" yaml:"includedTargets"`
	// Optional. The service account to run workflow invocations under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_workflow_config#service_account DataformRepositoryWorkflowConfig#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Optional. When set to true, transitive dependencies of included actions will be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_workflow_config#transitive_dependencies_included DataformRepositoryWorkflowConfig#transitive_dependencies_included}
	TransitiveDependenciesIncluded interface{} `field:"optional" json:"transitiveDependenciesIncluded" yaml:"transitiveDependenciesIncluded"`
	// Optional. When set to true, transitive dependents of included actions will be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataform_repository_workflow_config#transitive_dependents_included DataformRepositoryWorkflowConfig#transitive_dependents_included}
	TransitiveDependentsIncluded interface{} `field:"optional" json:"transitiveDependentsIncluded" yaml:"transitiveDependentsIncluded"`
}

