// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataQualitySpecRulesTemplateReferenceValues struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/dataplex_datascan#name DataplexDatascan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The string representation of the parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/dataplex_datascan#value DataplexDatascan#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

