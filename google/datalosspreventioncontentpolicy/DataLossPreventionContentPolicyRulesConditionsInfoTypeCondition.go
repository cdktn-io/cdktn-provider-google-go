// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyRulesConditionsInfoTypeCondition struct {
	// any_info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#any_info_type DataLossPreventionContentPolicy#any_info_type}
	AnyInfoType *DataLossPreventionContentPolicyRulesConditionsInfoTypeConditionAnyInfoType `field:"optional" json:"anyInfoType" yaml:"anyInfoType"`
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#info_types DataLossPreventionContentPolicy#info_types}
	InfoTypes *DataLossPreventionContentPolicyRulesConditionsInfoTypeConditionInfoTypes `field:"optional" json:"infoTypes" yaml:"infoTypes"`
	// The minimum number of findings required for this condition to be met. Defaults to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#min_count DataLossPreventionContentPolicy#min_count}
	MinCount *float64 `field:"optional" json:"minCount" yaml:"minCount"`
}

