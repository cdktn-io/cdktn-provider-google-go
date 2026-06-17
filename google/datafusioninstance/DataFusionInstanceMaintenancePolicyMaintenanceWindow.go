// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datafusioninstance


type DataFusionInstanceMaintenancePolicyMaintenanceWindow struct {
	// recurring_time_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/data_fusion_instance#recurring_time_window DataFusionInstance#recurring_time_window}
	RecurringTimeWindow *DataFusionInstanceMaintenancePolicyMaintenanceWindowRecurringTimeWindow `field:"required" json:"recurringTimeWindow" yaml:"recurringTimeWindow"`
}

