// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyNetworkRulesFromSources struct {
	// ip_blocks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/network_security_authz_policy#ip_blocks NetworkSecurityAuthzPolicy#ip_blocks}
	IpBlocks interface{} `field:"optional" json:"ipBlocks" yaml:"ipBlocks"`
	// principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/network_security_authz_policy#principals NetworkSecurityAuthzPolicy#principals}
	Principals interface{} `field:"optional" json:"principals" yaml:"principals"`
}

