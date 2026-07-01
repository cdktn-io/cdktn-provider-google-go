// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigBgpPeerConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gkeonprem_bare_metal_admin_cluster#asn GkeonpremBareMetalAdminCluster#asn}.
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// The IP address of the control plane node that connects to the external peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gkeonprem_bare_metal_admin_cluster#control_plane_nodes GkeonpremBareMetalAdminCluster#control_plane_nodes}
	ControlPlaneNodes *[]*string `field:"optional" json:"controlPlaneNodes" yaml:"controlPlaneNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/gkeonprem_bare_metal_admin_cluster#ip_address GkeonpremBareMetalAdminCluster#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

