// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowtool


type DialogflowToolOpenApiSpecAuthenticationOauthConfig struct {
	// The client ID from the OAuth provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_tool#client_id DialogflowTool#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// OAuth grant types. Possible values: ["CLIENT_CREDENTIAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_tool#oauth_grant_type DialogflowTool#oauth_grant_type}
	OauthGrantType *string `field:"required" json:"oauthGrantType" yaml:"oauthGrantType"`
	// The token endpoint in the OAuth provider to exchange for an access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_tool#token_endpoint DialogflowTool#token_endpoint}
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// Optional. The client secret from the OAuth provider. If the 'secretVersionForClientSecret' field is set, this field will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_tool#client_secret DialogflowTool#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Optional. The OAuth scopes to grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_tool#scopes DialogflowTool#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Optional.
	//
	// The name of the SecretManager secret version resource storing the client secret.
	// If this field is set, the 'clientSecret' field will be ignored.
	// Format: 'projects/{project}/secrets/{secret}/versions/{version}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_tool#secret_version_for_client_secret DialogflowTool#secret_version_for_client_secret}
	SecretVersionForClientSecret *string `field:"optional" json:"secretVersionForClientSecret" yaml:"secretVersionForClientSecret"`
}

