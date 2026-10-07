// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminigdaobservabilitysettingbinding


type GeminiGdaObservabilitySettingBindingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/gemini_gda_observability_setting_binding#create GeminiGdaObservabilitySettingBinding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/gemini_gda_observability_setting_binding#delete GeminiGdaObservabilitySettingBinding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/gemini_gda_observability_setting_binding#update GeminiGdaObservabilitySettingBinding#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

