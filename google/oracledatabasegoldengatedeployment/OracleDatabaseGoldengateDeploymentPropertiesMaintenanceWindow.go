// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengatedeployment


type OracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow struct {
	// Possible values: MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/oracle_database_goldengate_deployment#day OracleDatabaseGoldengateDeployment#day}
	Day *string `field:"required" json:"day" yaml:"day"`
	// Start hour for maintenance period. Hour is in UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/oracle_database_goldengate_deployment#start_hour OracleDatabaseGoldengateDeployment#start_hour}
	StartHour *float64 `field:"required" json:"startHour" yaml:"startHour"`
}

