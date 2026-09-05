// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesguardrail


type CesGuardrailLlmPromptSecurity struct {
	// custom_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/ces_guardrail#custom_policy CesGuardrail#custom_policy}
	CustomPolicy *CesGuardrailLlmPromptSecurityCustomPolicy `field:"optional" json:"customPolicy" yaml:"customPolicy"`
	// default_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/ces_guardrail#default_settings CesGuardrail#default_settings}
	DefaultSettings *CesGuardrailLlmPromptSecurityDefaultSettings `field:"optional" json:"defaultSettings" yaml:"defaultSettings"`
	// Determines the behavior when the guardrail encounters an LLM error.
	//
	// - If true: the guardrail is bypassed.
	// - If false (default): the guardrail triggers/blocks.
	// Note: If a custom policy is provided, this field is ignored in favor of
	// the policy's 'failOpen' configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/ces_guardrail#fail_open CesGuardrail#fail_open}
	FailOpen interface{} `field:"optional" json:"failOpen" yaml:"failOpen"`
}

