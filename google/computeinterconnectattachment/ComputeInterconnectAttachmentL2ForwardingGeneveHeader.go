// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeinterconnectattachment


type ComputeInterconnectAttachmentL2ForwardingGeneveHeader struct {
	// VNI is a 24-bit unique virtual network identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/compute_interconnect_attachment#vni ComputeInterconnectAttachment#vni}
	Vni *float64 `field:"optional" json:"vni" yaml:"vni"`
}

