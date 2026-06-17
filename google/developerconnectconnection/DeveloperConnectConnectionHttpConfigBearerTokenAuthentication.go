// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package developerconnectconnection


type DeveloperConnectConnectionHttpConfigBearerTokenAuthentication struct {
	// The token SecretManager secret version to authenticate as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/developer_connect_connection#token_secret_version DeveloperConnectConnection#token_secret_version}
	TokenSecretVersion *string `field:"optional" json:"tokenSecretVersion" yaml:"tokenSecretVersion"`
}

