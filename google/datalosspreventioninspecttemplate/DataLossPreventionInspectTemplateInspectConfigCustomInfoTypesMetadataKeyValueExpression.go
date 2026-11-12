// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesMetadataKeyValueExpression struct {
	// The regular expression for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_loss_prevention_inspect_template#key_regex DataLossPreventionInspectTemplate#key_regex}
	KeyRegex *string `field:"required" json:"keyRegex" yaml:"keyRegex"`
	// The regular expression for the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_loss_prevention_inspect_template#value_regex DataLossPreventionInspectTemplate#value_regex}
	ValueRegex *string `field:"required" json:"valueRegex" yaml:"valueRegex"`
}

