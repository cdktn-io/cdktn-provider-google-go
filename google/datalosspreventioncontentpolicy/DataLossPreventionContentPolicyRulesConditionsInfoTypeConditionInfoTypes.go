// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyRulesConditionsInfoTypeConditionInfoTypes struct {
	// List of info type names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#info_type_names DataLossPreventionContentPolicy#info_type_names}
	InfoTypeNames *[]*string `field:"required" json:"infoTypeNames" yaml:"infoTypeNames"`
}

