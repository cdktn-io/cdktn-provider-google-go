// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package privatecacapool


type PrivatecaCaPoolIssuancePolicyPassthroughExtensions struct {
	// additional_extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/privateca_ca_pool#additional_extensions PrivatecaCaPool#additional_extensions}
	AdditionalExtensions interface{} `field:"optional" json:"additionalExtensions" yaml:"additionalExtensions"`
	// A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions. Acceptable values are: 'BASE_KEY_USAGE', 'EXTENDED_KEY_USAGE', 'CA_OPTIONS', 'POLICY_IDS', 'AIA_OCSP_SERVERS', 'NAME_CONSTRAINTS'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/privateca_ca_pool#known_extensions PrivatecaCaPool#known_extensions}
	KnownExtensions *[]*string `field:"optional" json:"knownExtensions" yaml:"knownExtensions"`
}

