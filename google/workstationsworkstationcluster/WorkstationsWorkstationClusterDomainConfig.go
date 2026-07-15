// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workstationsworkstationcluster


type WorkstationsWorkstationClusterDomainConfig struct {
	// Domain used by Workstations for HTTP ingress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/workstations_workstation_cluster#domain WorkstationsWorkstationCluster#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

