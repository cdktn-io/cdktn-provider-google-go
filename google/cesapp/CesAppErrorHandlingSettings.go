// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesapp


type CesAppErrorHandlingSettings struct {
	// end_session_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#end_session_config CesApp#end_session_config}
	EndSessionConfig *CesAppErrorHandlingSettingsEndSessionConfig `field:"optional" json:"endSessionConfig" yaml:"endSessionConfig"`
	// The strategy to use for error handling. Possible values: NONE FALLBACK_RESPONSE END_SESSION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#error_handling_strategy CesApp#error_handling_strategy}
	ErrorHandlingStrategy *string `field:"optional" json:"errorHandlingStrategy" yaml:"errorHandlingStrategy"`
	// fallback_response_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#fallback_response_config CesApp#fallback_response_config}
	FallbackResponseConfig *CesAppErrorHandlingSettingsFallbackResponseConfig `field:"optional" json:"fallbackResponseConfig" yaml:"fallbackResponseConfig"`
}

