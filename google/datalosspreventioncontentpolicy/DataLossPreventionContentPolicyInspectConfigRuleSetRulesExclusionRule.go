// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule struct {
	// How the rule is applied. See the documentation for more information: https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#MatchingType Possible values: ["MATCHING_TYPE_FULL_MATCH", "MATCHING_TYPE_PARTIAL_MATCH", "MATCHING_TYPE_INVERSE_MATCH", "MATCHING_TYPE_RULE_SPECIFIC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#matching_type DataLossPreventionContentPolicy#matching_type}
	MatchingType *string `field:"required" json:"matchingType" yaml:"matchingType"`
	// dictionary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#dictionary DataLossPreventionContentPolicy#dictionary}
	Dictionary *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionary `field:"optional" json:"dictionary" yaml:"dictionary"`
	// exclude_by_hotword block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#exclude_by_hotword DataLossPreventionContentPolicy#exclude_by_hotword}
	ExcludeByHotword *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotword `field:"optional" json:"excludeByHotword" yaml:"excludeByHotword"`
	// exclude_by_image_findings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#exclude_by_image_findings DataLossPreventionContentPolicy#exclude_by_image_findings}
	ExcludeByImageFindings *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings `field:"optional" json:"excludeByImageFindings" yaml:"excludeByImageFindings"`
	// exclude_info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#exclude_info_types DataLossPreventionContentPolicy#exclude_info_types}
	ExcludeInfoTypes *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes `field:"optional" json:"excludeInfoTypes" yaml:"excludeInfoTypes"`
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#regex DataLossPreventionContentPolicy#regex}
	Regex *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegex `field:"optional" json:"regex" yaml:"regex"`
}

