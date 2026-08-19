// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computerolloutplan


type ComputeRolloutPlanWavesSelectorsResourceHierarchySelector struct {
	// Format: "folders/{folder_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/compute_rollout_plan#included_folders ComputeRolloutPlan#included_folders}
	IncludedFolders *[]*string `field:"optional" json:"includedFolders" yaml:"includedFolders"`
	// Format: "organizations/{organization_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/compute_rollout_plan#included_organizations ComputeRolloutPlan#included_organizations}
	IncludedOrganizations *[]*string `field:"optional" json:"includedOrganizations" yaml:"includedOrganizations"`
	// Format: "projects/{project_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/compute_rollout_plan#included_projects ComputeRolloutPlan#included_projects}
	IncludedProjects *[]*string `field:"optional" json:"includedProjects" yaml:"includedProjects"`
}

