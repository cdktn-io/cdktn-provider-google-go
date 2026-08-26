// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workstationsworkstationconfig


type WorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfig struct {
	// Whether the instance has confidential compute enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/workstations_workstation_config#enable_confidential_compute WorkstationsWorkstationConfigA#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"optional" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
}

