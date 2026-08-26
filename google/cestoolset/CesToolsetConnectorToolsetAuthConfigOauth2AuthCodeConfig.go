// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestoolset


type CesToolsetConnectorToolsetAuthConfigOauth2AuthCodeConfig struct {
	// Oauth token parameter name to pass through. Must be in the format '$context.variables.<name_of_variable>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/ces_toolset#oauth_token CesToolset#oauth_token}
	OauthToken *string `field:"required" json:"oauthToken" yaml:"oauthToken"`
}

