// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkconnectivityspoke


type NetworkConnectivitySpokeGatewayIpRangeReservations struct {
	// A block of IP address ranges used to allocate supporting infrastructure for this gateway—for example, 10.1.2.0/23. The IP address block must be a /23 range. This IP address block must not overlap with subnets in any spoke or peer network that the gateway can communicate with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/network_connectivity_spoke#ip_range NetworkConnectivitySpoke#ip_range}
	IpRange *string `field:"required" json:"ipRange" yaml:"ipRange"`
}

