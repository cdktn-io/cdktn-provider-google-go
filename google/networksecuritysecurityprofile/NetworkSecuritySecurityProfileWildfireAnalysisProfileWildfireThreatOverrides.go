// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecuritysecurityprofile


type NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireThreatOverrides struct {
	// Threat action override. Possible values: ["WILDFIRE_DEFAULT_ACTION", "WILDFIRE_ALLOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/network_security_security_profile#action NetworkSecuritySecurityProfile#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Vendor-specific ID of a threat to override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/network_security_security_profile#threat_id NetworkSecuritySecurityProfile#threat_id}
	ThreatId *string `field:"required" json:"threatId" yaml:"threatId"`
}

