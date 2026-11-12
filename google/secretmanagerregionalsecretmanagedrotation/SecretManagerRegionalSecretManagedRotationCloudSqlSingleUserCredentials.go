// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretmanagerregionalsecretmanagedrotation


type SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentials struct {
	// The name of the Cloud SQL instance. The instance must be in the same project as the regional secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/secret_manager_regional_secret_managed_rotation#instance_id SecretManagerRegionalSecretManagedRotation#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The username of the Cloud SQL database user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/secret_manager_regional_secret_managed_rotation#username SecretManagerRegionalSecretManagedRotation#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The initial password to set for the database user. If not provided, a random password will be generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/secret_manager_regional_secret_managed_rotation#password SecretManagerRegionalSecretManagedRotation#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The initial password to set for the database user. If not provided, a random password will be generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/secret_manager_regional_secret_managed_rotation#password_wo SecretManagerRegionalSecretManagedRotation#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// Triggers update of 'password_wo' write-only.
	//
	// Increment this value when an update to 'password_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/secret_manager_regional_secret_managed_rotation#password_wo_version SecretManagerRegionalSecretManagedRotation#password_wo_version}
	PasswordWoVersion *string `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
}

