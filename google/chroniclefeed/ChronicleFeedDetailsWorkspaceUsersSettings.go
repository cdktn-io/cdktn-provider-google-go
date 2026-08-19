// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsWorkspaceUsersSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsWorkspaceUsersSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Projection Type. Possible values: BASIC_PROJECTION FULL_PROJECTION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#projection_type ChronicleFeed#projection_type}
	ProjectionType *string `field:"optional" json:"projectionType" yaml:"projectionType"`
	// Customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/chronicle_feed#workspace_customer_id ChronicleFeed#workspace_customer_id}
	WorkspaceCustomerId *string `field:"optional" json:"workspaceCustomerId" yaml:"workspaceCustomerId"`
}

