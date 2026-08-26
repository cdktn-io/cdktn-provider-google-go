// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyNetworkRulesFromSourcesPrincipalsPrincipal struct {
	// The input string must match exactly the string specified here. Examples: * abc only matches the value abc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/network_security_authz_policy#exact NetworkSecurityAuthzPolicy#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
}

