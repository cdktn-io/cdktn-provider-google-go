// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package modelarmortemplate


type ModelArmorTemplateFilterConfigFilterRuleSettingsRuleSetsRulesExclusionRule struct {
	// dictionary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/model_armor_template#dictionary ModelArmorTemplate#dictionary}
	Dictionary *ModelArmorTemplateFilterConfigFilterRuleSettingsRuleSetsRulesExclusionRuleDictionary `field:"optional" json:"dictionary" yaml:"dictionary"`
	// Specifies the scope of the input that should be matched against the exclusion rule. If unspecified, defaults to MATCHING_SCOPE_PARTIAL_MATCH. Possible values: MATCHING_SCOPE_FULL_MATCH MATCHING_SCOPE_PARTIAL_MATCH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/model_armor_template#matching_scope ModelArmorTemplate#matching_scope}
	MatchingScope *string `field:"optional" json:"matchingScope" yaml:"matchingScope"`
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/model_armor_template#regex ModelArmorTemplate#regex}
	Regex *ModelArmorTemplateFilterConfigFilterRuleSettingsRuleSetsRulesExclusionRuleRegex `field:"optional" json:"regex" yaml:"regex"`
}

