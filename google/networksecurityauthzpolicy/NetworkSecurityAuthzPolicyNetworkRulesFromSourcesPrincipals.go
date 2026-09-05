// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyNetworkRulesFromSourcesPrincipals struct {
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/network_security_authz_policy#principal NetworkSecurityAuthzPolicy#principal}
	Principal *NetworkSecurityAuthzPolicyNetworkRulesFromSourcesPrincipalsPrincipal `field:"optional" json:"principal" yaml:"principal"`
	// An enum to decide what principal value the principal rule will match against.
	//
	// If not specified, the PrincipalSelector is CLIENT_CERT_URI_SAN. Default value: "CLIENT_CERT_URI_SAN" Possible values: ["PRINCIPAL_SELECTOR_UNSPECIFIED", "CLIENT_CERT_URI_SAN", "CLIENT_CERT_DNS_NAME_SAN", "CLIENT_CERT_COMMON_NAME"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/network_security_authz_policy#principal_selector NetworkSecurityAuthzPolicy#principal_selector}
	PrincipalSelector *string `field:"optional" json:"principalSelector" yaml:"principalSelector"`
}

