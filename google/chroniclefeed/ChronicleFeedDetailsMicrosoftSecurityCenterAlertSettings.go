// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings struct {
	// API Auth Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#auth_endpoint ChronicleFeed#auth_endpoint}
	AuthEndpoint *string `field:"optional" json:"authEndpoint" yaml:"authEndpoint"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsMicrosoftSecurityCenterAlertSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#hostname ChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Subscription ID of the Microsoft security center alert settings alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#subscription_id ChronicleFeed#subscription_id}
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// Tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#tenant_id ChronicleFeed#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

