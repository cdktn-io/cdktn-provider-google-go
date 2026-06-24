// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsSalesforceSettings struct {
	// API hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#hostname ChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// oauth_jwt_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#oauth_jwt_credentials ChronicleFeed#oauth_jwt_credentials}
	OauthJwtCredentials *ChronicleFeedDetailsSalesforceSettingsOauthJwtCredentials `field:"optional" json:"oauthJwtCredentials" yaml:"oauthJwtCredentials"`
	// oauth_password_grant_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#oauth_password_grant_auth ChronicleFeed#oauth_password_grant_auth}
	OauthPasswordGrantAuth *ChronicleFeedDetailsSalesforceSettingsOauthPasswordGrantAuth `field:"optional" json:"oauthPasswordGrantAuth" yaml:"oauthPasswordGrantAuth"`
}

