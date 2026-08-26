// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestoolset


type CesToolsetOpenApiToolsetTlsConfig struct {
	// ca_certs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/ces_toolset#ca_certs CesToolset#ca_certs}
	CaCerts interface{} `field:"required" json:"caCerts" yaml:"caCerts"`
}

