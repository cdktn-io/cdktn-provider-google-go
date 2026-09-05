// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkconnectivitygatewayadvertisedroute


type NetworkConnectivityGatewayAdvertisedRouteTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/network_connectivity_gateway_advertised_route#create NetworkConnectivityGatewayAdvertisedRoute#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/network_connectivity_gateway_advertised_route#delete NetworkConnectivityGatewayAdvertisedRoute#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/network_connectivity_gateway_advertised_route#update NetworkConnectivityGatewayAdvertisedRoute#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

