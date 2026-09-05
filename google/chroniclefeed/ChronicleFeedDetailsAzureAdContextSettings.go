// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAzureAdContextSettings struct {
	// API Auth Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#auth_endpoint ChronicleFeed#auth_endpoint}
	AuthEndpoint *string `field:"optional" json:"authEndpoint" yaml:"authEndpoint"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#authentication ChronicleFeed#authentication}
	Authentication *ChronicleFeedDetailsAzureAdContextSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#hostname ChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Whether to retrieve device information in user context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#retrieve_devices ChronicleFeed#retrieve_devices}
	RetrieveDevices interface{} `field:"optional" json:"retrieveDevices" yaml:"retrieveDevices"`
	// Whether to retrieve group information in user context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#retrieve_groups ChronicleFeed#retrieve_groups}
	RetrieveGroups interface{} `field:"optional" json:"retrieveGroups" yaml:"retrieveGroups"`
	// Tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_feed#tenant_id ChronicleFeed#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

