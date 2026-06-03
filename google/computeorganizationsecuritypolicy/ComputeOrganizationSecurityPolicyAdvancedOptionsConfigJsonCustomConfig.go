// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeorganizationsecuritypolicy


type ComputeOrganizationSecurityPolicyAdvancedOptionsConfigJsonCustomConfig struct {
	// A list of content types to be parsed as JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/compute_organization_security_policy#content_types ComputeOrganizationSecurityPolicy#content_types}
	ContentTypes *[]*string `field:"required" json:"contentTypes" yaml:"contentTypes"`
}

