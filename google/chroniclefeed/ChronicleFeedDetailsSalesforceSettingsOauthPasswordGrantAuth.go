// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsSalesforceSettingsOauthPasswordGrantAuth struct {
	// Client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#client_id ChronicleFeed#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#client_secret ChronicleFeed#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#password ChronicleFeed#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Token endpoint to get the OAuth token from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#token_endpoint ChronicleFeed#token_endpoint}
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// Username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#user ChronicleFeed#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

