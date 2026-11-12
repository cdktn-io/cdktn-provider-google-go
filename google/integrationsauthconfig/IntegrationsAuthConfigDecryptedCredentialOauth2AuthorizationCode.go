// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig


type IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode struct {
	// Indicates if the user has opted in Google Reauth Policy.
	//
	// If opted in, the refresh token will be valid for 20 hours, after which time users must re-authenticate in order to obtain a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integrations_auth_config#apply_reauth_policy IntegrationsAuthConfig#apply_reauth_policy}
	ApplyReauthPolicy interface{} `field:"optional" json:"applyReauthPolicy" yaml:"applyReauthPolicy"`
	// The auth url endpoint to send the auth code request to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integrations_auth_config#auth_endpoint IntegrationsAuthConfig#auth_endpoint}
	AuthEndpoint *string `field:"optional" json:"authEndpoint" yaml:"authEndpoint"`
	// auth_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integrations_auth_config#auth_params IntegrationsAuthConfig#auth_params}
	AuthParams *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeAuthParams `field:"optional" json:"authParams" yaml:"authParams"`
	// The client's id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integrations_auth_config#client_id IntegrationsAuthConfig#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client's secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integrations_auth_config#client_secret IntegrationsAuthConfig#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Represent how to pass parameters to fetch access token Possible values: ["REQUEST_BODY", "QUERY_PARAMETERS", "ENCODED_HEADER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integrations_auth_config#request_type IntegrationsAuthConfig#request_type}
	RequestType *string `field:"optional" json:"requestType" yaml:"requestType"`
	// A space-delimited list of requested scope permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integrations_auth_config#scope IntegrationsAuthConfig#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The token url endpoint to send the token request to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integrations_auth_config#token_endpoint IntegrationsAuthConfig#token_endpoint}
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// token_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integrations_auth_config#token_params IntegrationsAuthConfig#token_params}
	TokenParams *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParams `field:"optional" json:"tokenParams" yaml:"tokenParams"`
}

