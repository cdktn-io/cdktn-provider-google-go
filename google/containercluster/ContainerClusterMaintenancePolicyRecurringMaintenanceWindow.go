// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterMaintenancePolicyRecurringMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#recurrence ContainerCluster#recurrence}.
	Recurrence *string `field:"required" json:"recurrence" yaml:"recurrence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#window_duration ContainerCluster#window_duration}.
	WindowDuration *string `field:"required" json:"windowDuration" yaml:"windowDuration"`
	// window_start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#window_start_time ContainerCluster#window_start_time}
	WindowStartTime *ContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime `field:"required" json:"windowStartTime" yaml:"windowStartTime"`
	// delay_until block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#delay_until ContainerCluster#delay_until}
	DelayUntil *ContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil `field:"optional" json:"delayUntil" yaml:"delayUntil"`
}

