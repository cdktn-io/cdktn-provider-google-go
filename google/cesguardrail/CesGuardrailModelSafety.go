// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesguardrail


type CesGuardrailModelSafety struct {
	// safety_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/ces_guardrail#safety_settings CesGuardrail#safety_settings}
	SafetySettings interface{} `field:"required" json:"safetySettings" yaml:"safetySettings"`
}

