// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesguardrail


type CesGuardrailActionRespondImmediatelyResponses struct {
	// Text for the agent to respond with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#text CesGuardrail#text}
	Text *string `field:"required" json:"text" yaml:"text"`
	// Whether the response is disabled. Disabled responses are not used by the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#disabled CesGuardrail#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

