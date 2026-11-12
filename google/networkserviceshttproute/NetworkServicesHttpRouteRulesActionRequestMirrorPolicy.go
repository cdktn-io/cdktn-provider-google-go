// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesActionRequestMirrorPolicy struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/network_services_http_route#destination NetworkServicesHttpRoute#destination}
	Destination *NetworkServicesHttpRouteRulesActionRequestMirrorPolicyDestination `field:"optional" json:"destination" yaml:"destination"`
	// The percentage of requests to get mirrored to the desired destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/network_services_http_route#mirror_percent NetworkServicesHttpRoute#mirror_percent}
	MirrorPercent *float64 `field:"optional" json:"mirrorPercent" yaml:"mirrorPercent"`
}

