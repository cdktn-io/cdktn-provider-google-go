// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsWorkspaceAlertsSettingsAuthentication struct {
	// claims block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/chronicle_feed#claims ChronicleFeed#claims}
	Claims *ChronicleFeedDetailsWorkspaceAlertsSettingsAuthenticationClaims `field:"optional" json:"claims" yaml:"claims"`
	// rs_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/chronicle_feed#rs_credentials ChronicleFeed#rs_credentials}
	RsCredentials *ChronicleFeedDetailsWorkspaceAlertsSettingsAuthenticationRsCredentials `field:"optional" json:"rsCredentials" yaml:"rsCredentials"`
	// Token endpoint to get the OAuth token from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/chronicle_feed#token_endpoint ChronicleFeed#token_endpoint}
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
}

