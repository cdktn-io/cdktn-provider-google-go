// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dnspolicy


type DnsPolicyAlternativeNameServerConfigOutboundEndpoints struct {
	// Specified as a resource path including project and region. This should be formatted like 'projects/{project}/locations/{location}/outboundEndpoints/{endpoint}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dns_policy#name DnsPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

