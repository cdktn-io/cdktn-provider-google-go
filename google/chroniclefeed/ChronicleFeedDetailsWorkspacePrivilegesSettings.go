// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsWorkspacePrivilegesSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsWorkspacePrivilegesSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_feed#workspace_customer_id ChronicleFeed#workspace_customer_id}
	WorkspaceCustomerId *string `field:"optional" json:"workspaceCustomerId" yaml:"workspaceCustomerId"`
}

