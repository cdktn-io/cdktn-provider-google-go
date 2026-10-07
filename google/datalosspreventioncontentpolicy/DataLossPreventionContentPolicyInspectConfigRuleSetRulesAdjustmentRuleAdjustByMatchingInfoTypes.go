// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes struct {
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#info_types DataLossPreventionContentPolicy#info_types}
	InfoTypes interface{} `field:"required" json:"infoTypes" yaml:"infoTypes"`
	// How the adjustment rule is applied. Possible values: ["MATCHING_TYPE_FULL_MATCH", "MATCHING_TYPE_PARTIAL_MATCH", "MATCHING_TYPE_INVERSE_MATCH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#matching_type DataLossPreventionContentPolicy#matching_type}
	MatchingType *string `field:"required" json:"matchingType" yaml:"matchingType"`
	// Minimum likelihood of the adjustByMatchingInfoTypes infoTypes finding. Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#min_likelihood DataLossPreventionContentPolicy#min_likelihood}
	MinLikelihood *string `field:"required" json:"minLikelihood" yaml:"minLikelihood"`
}

