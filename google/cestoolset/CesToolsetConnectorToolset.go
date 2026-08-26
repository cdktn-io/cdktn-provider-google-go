// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestoolset


type CesToolsetConnectorToolset struct {
	// The full resource name of the referenced Integration Connectors Connection. Format: 'projects/{project}/locations/{location}/connections/{connection}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/ces_toolset#connection CesToolset#connection}
	Connection *string `field:"required" json:"connection" yaml:"connection"`
	// connector_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/ces_toolset#connector_actions CesToolset#connector_actions}
	ConnectorActions interface{} `field:"required" json:"connectorActions" yaml:"connectorActions"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/ces_toolset#auth_config CesToolset#auth_config}
	AuthConfig *CesToolsetConnectorToolsetAuthConfig `field:"optional" json:"authConfig" yaml:"authConfig"`
}

