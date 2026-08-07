// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyNetworkRules struct {
	// from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/network_security_authz_policy#from NetworkSecurityAuthzPolicy#from}
	From *NetworkSecurityAuthzPolicyNetworkRulesFrom `field:"optional" json:"from" yaml:"from"`
	// to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/network_security_authz_policy#to NetworkSecurityAuthzPolicy#to}
	To *NetworkSecurityAuthzPolicyNetworkRulesTo `field:"optional" json:"to" yaml:"to"`
}

