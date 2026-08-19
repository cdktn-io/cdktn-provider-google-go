// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#day ContainerCluster#day}.
	Day *float64 `field:"required" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#month ContainerCluster#month}.
	Month *float64 `field:"required" json:"month" yaml:"month"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#year ContainerCluster#year}.
	Year *float64 `field:"required" json:"year" yaml:"year"`
}

