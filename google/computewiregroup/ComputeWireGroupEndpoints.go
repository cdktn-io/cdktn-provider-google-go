// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computewiregroup


type ComputeWireGroupEndpoints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/compute_wire_group#endpoint ComputeWireGroup#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// interconnects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/compute_wire_group#interconnects ComputeWireGroup#interconnects}
	Interconnects interface{} `field:"optional" json:"interconnects" yaml:"interconnects"`
}

