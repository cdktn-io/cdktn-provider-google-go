// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression struct {
	// The regular expression for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#key_regex DataLossPreventionContentPolicy#key_regex}
	KeyRegex *string `field:"required" json:"keyRegex" yaml:"keyRegex"`
	// The regular expression for the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#value_regex DataLossPreventionContentPolicy#value_regex}
	ValueRegex *string `field:"required" json:"valueRegex" yaml:"valueRegex"`
}

