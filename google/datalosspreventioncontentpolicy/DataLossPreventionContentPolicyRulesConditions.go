// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyRulesConditions struct {
	// info_type_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#info_type_condition DataLossPreventionContentPolicy#info_type_condition}
	InfoTypeCondition *DataLossPreventionContentPolicyRulesConditionsInfoTypeCondition `field:"optional" json:"infoTypeCondition" yaml:"infoTypeCondition"`
}

