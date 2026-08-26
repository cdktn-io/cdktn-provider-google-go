// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestoolset


type CesToolsetConnectorToolsetAuthConfigOauth2JwtBearerConfig struct {
	// Client parameter name to pass through. Must be in the format '$context.variables.<name_of_variable>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/ces_toolset#client_key CesToolset#client_key}
	ClientKey *string `field:"required" json:"clientKey" yaml:"clientKey"`
	// Issuer parameter name to pass through. Must be in the format '$context.variables.<name_of_variable>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/ces_toolset#issuer CesToolset#issuer}
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Subject parameter name to pass through. Must be in the format '$context.variables.<name_of_variable>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/ces_toolset#subject CesToolset#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

