// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datafusioninstance


type DataFusionInstanceMaintenancePolicy struct {
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/data_fusion_instance#maintenance_window DataFusionInstance#maintenance_window}
	MaintenanceWindow *DataFusionInstanceMaintenancePolicyMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
}

