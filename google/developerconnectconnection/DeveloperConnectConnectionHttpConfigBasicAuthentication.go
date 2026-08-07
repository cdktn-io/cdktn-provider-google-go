// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package developerconnectconnection


type DeveloperConnectConnectionHttpConfigBasicAuthentication struct {
	// The username to authenticate as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/developer_connect_connection#username DeveloperConnectConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The password SecretManager secret version to authenticate as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/developer_connect_connection#password_secret_version DeveloperConnectConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
}

