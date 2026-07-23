// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkserviceswasmplugin


type NetworkServicesWasmPluginTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/network_services_wasm_plugin#create NetworkServicesWasmPlugin#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/network_services_wasm_plugin#delete NetworkServicesWasmPlugin#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/network_services_wasm_plugin#update NetworkServicesWasmPlugin#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

