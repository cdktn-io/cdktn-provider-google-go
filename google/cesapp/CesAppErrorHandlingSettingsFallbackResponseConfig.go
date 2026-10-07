// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesapp


type CesAppErrorHandlingSettingsFallbackResponseConfig struct {
	// The fallback messages in case of system errors (e.g. LLM errors), mapped by supported language code (https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/reference/language).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#custom_fallback_messages CesApp#custom_fallback_messages}
	CustomFallbackMessages *map[string]*string `field:"optional" json:"customFallbackMessages" yaml:"customFallbackMessages"`
	// The maximum number of fallback attempts to make before the agent emitting EndSession Signal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#max_fallback_attempts CesApp#max_fallback_attempts}
	MaxFallbackAttempts *float64 `field:"optional" json:"maxFallbackAttempts" yaml:"maxFallbackAttempts"`
}

