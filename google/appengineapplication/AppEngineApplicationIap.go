// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appengineapplication


type AppEngineApplicationIap struct {
	// OAuth2 client ID to use for the authentication flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/app_engine_application#oauth2_client_id AppEngineApplication#oauth2_client_id}
	Oauth2ClientId *string `field:"required" json:"oauth2ClientId" yaml:"oauth2ClientId"`
	// Adapted for use with the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/app_engine_application#enabled AppEngineApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// OAuth2 client secret to use for the authentication flow.
	//
	// The SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/app_engine_application#oauth2_client_secret AppEngineApplication#oauth2_client_secret}
	Oauth2ClientSecret *string `field:"optional" json:"oauth2ClientSecret" yaml:"oauth2ClientSecret"`
	// OAuth2 client secret to use for the authentication flow.
	//
	// The SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/app_engine_application#oauth2_client_secret_wo AppEngineApplication#oauth2_client_secret_wo}
	Oauth2ClientSecretWo *string `field:"optional" json:"oauth2ClientSecretWo" yaml:"oauth2ClientSecretWo"`
	// Triggers update of `oauth2_client_secret_wo` write-only.
	//
	// Increment this value when an update to `oauth2_client_secret_wo` is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/app_engine_application#oauth2_client_secret_wo_version AppEngineApplication#oauth2_client_secret_wo_version}
	Oauth2ClientSecretWoVersion *string `field:"optional" json:"oauth2ClientSecretWoVersion" yaml:"oauth2ClientSecretWoVersion"`
}

