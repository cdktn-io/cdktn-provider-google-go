// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestoolset


type CesToolsetConnectorToolsetAuthConfig struct {
	// oauth2_auth_code_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/ces_toolset#oauth2_auth_code_config CesToolset#oauth2_auth_code_config}
	Oauth2AuthCodeConfig *CesToolsetConnectorToolsetAuthConfigOauth2AuthCodeConfig `field:"optional" json:"oauth2AuthCodeConfig" yaml:"oauth2AuthCodeConfig"`
	// oauth2_jwt_bearer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/ces_toolset#oauth2_jwt_bearer_config CesToolset#oauth2_jwt_bearer_config}
	Oauth2JwtBearerConfig *CesToolsetConnectorToolsetAuthConfigOauth2JwtBearerConfig `field:"optional" json:"oauth2JwtBearerConfig" yaml:"oauth2JwtBearerConfig"`
}

