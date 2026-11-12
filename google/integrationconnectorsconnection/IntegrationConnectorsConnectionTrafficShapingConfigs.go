// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionTrafficShapingConfigs struct {
	// Specifies the duration over which the API call quota limits are calculated.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integration_connectors_connection#duration IntegrationConnectorsConnection#duration}
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Maximum number of api calls allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/integration_connectors_connection#quota_limit IntegrationConnectorsConnection#quota_limit}
	QuotaLimit *string `field:"required" json:"quotaLimit" yaml:"quotaLimit"`
}

