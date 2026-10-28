// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecuritysecurityprofile


type NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireSubmissionRules struct {
	// Direction of traffic to match for a rule. Possible values: ["UPLOAD", "DOWNLOAD", "BOTH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/network_security_security_profile#direction NetworkSecuritySecurityProfile#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Defines the file selection mode for a rule. Possible values: ["ALL_FILE_TYPES", "CUSTOM_FILE_TYPES"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/network_security_security_profile#file_selection_mode NetworkSecuritySecurityProfile#file_selection_mode}
	FileSelectionMode *string `field:"required" json:"fileSelectionMode" yaml:"fileSelectionMode"`
	// custom_file_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/network_security_security_profile#custom_file_types NetworkSecuritySecurityProfile#custom_file_types}
	CustomFileTypes *NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireSubmissionRulesCustomFileTypes `field:"optional" json:"customFileTypes" yaml:"customFileTypes"`
}

