// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsDummyLogTypeSettings struct {
	// Full API Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#api_endpoint ChronicleFeed#api_endpoint}
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsDummyLogTypeSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
}

