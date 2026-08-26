// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengatedeployment


type OracleDatabaseGoldengateDeploymentProperties struct {
	// A valid Goldengate Deployment type. For a list of supported types, use the 'ListGoldengateDeploymentTypes' operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_deployment#deployment_type OracleDatabaseGoldengateDeployment#deployment_type}
	DeploymentType *string `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// ogg_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_deployment#ogg_data OracleDatabaseGoldengateDeployment#ogg_data}
	OggData *OracleDatabaseGoldengateDeploymentPropertiesOggData `field:"required" json:"oggData" yaml:"oggData"`
	// backup_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_deployment#backup_schedule OracleDatabaseGoldengateDeployment#backup_schedule}
	BackupSchedule *OracleDatabaseGoldengateDeploymentPropertiesBackupSchedule `field:"optional" json:"backupSchedule" yaml:"backupSchedule"`
	// The Minimum number of OCPUs to be made available for this Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_deployment#cpu_core_count OracleDatabaseGoldengateDeployment#cpu_core_count}
	CpuCoreCount *float64 `field:"optional" json:"cpuCoreCount" yaml:"cpuCoreCount"`
	// deployment_diagnostic_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_deployment#deployment_diagnostic_data OracleDatabaseGoldengateDeployment#deployment_diagnostic_data}
	DeploymentDiagnosticData *OracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData `field:"optional" json:"deploymentDiagnosticData" yaml:"deploymentDiagnosticData"`
	// The description of the GoldengateDeployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_deployment#description OracleDatabaseGoldengateDeployment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The environment type of the GoldengateDeployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_deployment#environment_type OracleDatabaseGoldengateDeployment#environment_type}
	EnvironmentType *string `field:"optional" json:"environmentType" yaml:"environmentType"`
	// Indicates if auto scaling is enabled for the Deployment's CPU core count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_deployment#is_auto_scaling_enabled OracleDatabaseGoldengateDeployment#is_auto_scaling_enabled}
	IsAutoScalingEnabled interface{} `field:"optional" json:"isAutoScalingEnabled" yaml:"isAutoScalingEnabled"`
	// The Oracle license model that applies to a Deployment. Possible values: LICENSE_INCLUDED BRING_YOUR_OWN_LICENSE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_deployment#license_model OracleDatabaseGoldengateDeployment#license_model}
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// maintenance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_deployment#maintenance_config OracleDatabaseGoldengateDeployment#maintenance_config}
	MaintenanceConfig *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig `field:"optional" json:"maintenanceConfig" yaml:"maintenanceConfig"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/oracle_database_goldengate_deployment#maintenance_window OracleDatabaseGoldengateDeployment#maintenance_window}
	MaintenanceWindow *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
}

