// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkconnectivityspoke


type NetworkConnectivitySpokeGateway struct {
	// the capacity of the gateway spoke, in Gbps. Possible values: ["CAPACITY_1_GBPS", "CAPACITY_10_GBPS", "CAPACITY_100_GBPS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/network_connectivity_spoke#capacity NetworkConnectivitySpoke#capacity}
	Capacity *string `field:"required" json:"capacity" yaml:"capacity"`
	// ip_range_reservations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/network_connectivity_spoke#ip_range_reservations NetworkConnectivitySpoke#ip_range_reservations}
	IpRangeReservations interface{} `field:"required" json:"ipRangeReservations" yaml:"ipRangeReservations"`
}

