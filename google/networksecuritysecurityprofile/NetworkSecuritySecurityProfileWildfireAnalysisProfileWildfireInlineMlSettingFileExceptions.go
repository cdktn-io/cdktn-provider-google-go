// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecuritysecurityprofile


type NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSettingFileExceptions struct {
	// Machine learning partial hash of the file to exclude from WildFire Inline ML analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/network_security_security_profile#partial_hash NetworkSecuritySecurityProfile#partial_hash}
	PartialHash *string `field:"required" json:"partialHash" yaml:"partialHash"`
	// The file name associated with the partial hash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/network_security_security_profile#filename NetworkSecuritySecurityProfile#filename}
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
}

