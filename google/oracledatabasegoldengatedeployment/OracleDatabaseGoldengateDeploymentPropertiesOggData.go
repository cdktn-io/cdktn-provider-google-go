// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengatedeployment


type OracleDatabaseGoldengateDeploymentPropertiesOggData struct {
	// The Goldengate deployment console username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/oracle_database_goldengate_deployment#admin_username OracleDatabaseGoldengateDeployment#admin_username}
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// The name given to the Goldengate service deployment.
	//
	// The name must be 1 to
	// 32 characters long, must contain only alphanumeric characters and must
	// start with a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/oracle_database_goldengate_deployment#deployment OracleDatabaseGoldengateDeployment#deployment}
	Deployment *string `field:"required" json:"deployment" yaml:"deployment"`
	// The Goldengate deployment console password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/oracle_database_goldengate_deployment#admin_password OracleDatabaseGoldengateDeployment#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Input only. The Goldengate deployment console password secret version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/oracle_database_goldengate_deployment#admin_password_secret_version OracleDatabaseGoldengateDeployment#admin_password_secret_version}
	AdminPasswordSecretVersion *string `field:"optional" json:"adminPasswordSecretVersion" yaml:"adminPasswordSecretVersion"`
	// group_roles_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/oracle_database_goldengate_deployment#group_roles_mapping OracleDatabaseGoldengateDeployment#group_roles_mapping}
	GroupRolesMapping *OracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping `field:"optional" json:"groupRolesMapping" yaml:"groupRolesMapping"`
	// Version of OGG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/oracle_database_goldengate_deployment#ogg_version OracleDatabaseGoldengateDeployment#ogg_version}
	OggVersion *string `field:"optional" json:"oggVersion" yaml:"oggVersion"`
}

