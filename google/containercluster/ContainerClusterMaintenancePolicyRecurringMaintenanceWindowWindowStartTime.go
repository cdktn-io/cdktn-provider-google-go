// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#hours ContainerCluster#hours}.
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#minutes ContainerCluster#minutes}.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#seconds ContainerCluster#seconds}.
	Seconds *float64 `field:"required" json:"seconds" yaml:"seconds"`
}

