// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package modelarmortemplate


type ModelArmorTemplateFilterConfigFilterRuleSettingsRuleSets struct {
	// List of filters this rule set is applied to. Possible values: PROMPT_INJECTION_AND_JAILBREAK RESPONSIBLE_AI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/model_armor_template#filter_types ModelArmorTemplate#filter_types}
	FilterTypes *[]*string `field:"required" json:"filterTypes" yaml:"filterTypes"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/model_armor_template#rules ModelArmorTemplate#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

