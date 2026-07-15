// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanExecutionIdentityServiceAccount struct {
	// Service account email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/dataplex_datascan#email DataplexDatascan#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

