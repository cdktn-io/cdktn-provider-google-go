// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecuritysecurityprofile


type NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting struct {
	// file_exceptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/network_security_security_profile#file_exceptions NetworkSecuritySecurityProfile#file_exceptions}
	FileExceptions interface{} `field:"optional" json:"fileExceptions" yaml:"fileExceptions"`
	// inline_ml_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/network_security_security_profile#inline_ml_configs NetworkSecuritySecurityProfile#inline_ml_configs}
	InlineMlConfigs interface{} `field:"optional" json:"inlineMlConfigs" yaml:"inlineMlConfigs"`
}

