// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed


type ChronicleFeedDetailsAzureBlobStoreV2SettingsAuthenticationAzureV2WorkloadIdentityFederation struct {
	// OAuth client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#client_id ChronicleFeed#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Subject ID of the Azure subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#subject_id ChronicleFeed#subject_id}
	SubjectId *string `field:"required" json:"subjectId" yaml:"subjectId"`
	// Tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_feed#tenant_id ChronicleFeed#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

