// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workstationsworkstationcluster


type WorkstationsWorkstationClusterPrivateClusterConfig struct {
	// Whether Workstations endpoint is private.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/workstations_workstation_cluster#enable_private_endpoint WorkstationsWorkstationCluster#enable_private_endpoint}
	EnablePrivateEndpoint interface{} `field:"required" json:"enablePrivateEndpoint" yaml:"enablePrivateEndpoint"`
	// Additional project IDs that are allowed to attach to the workstation cluster's service attachment.
	//
	// By default, the workstation cluster's project and the VPC host project (if different) are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/workstations_workstation_cluster#allowed_projects WorkstationsWorkstationCluster#allowed_projects}
	AllowedProjects *[]*string `field:"optional" json:"allowedProjects" yaml:"allowedProjects"`
}

