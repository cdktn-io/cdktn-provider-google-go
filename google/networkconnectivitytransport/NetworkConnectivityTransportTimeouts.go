// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkconnectivitytransport


type NetworkConnectivityTransportTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/network_connectivity_transport#create NetworkConnectivityTransport#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/network_connectivity_transport#delete NetworkConnectivityTransport#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/network_connectivity_transport#update NetworkConnectivityTransport#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

