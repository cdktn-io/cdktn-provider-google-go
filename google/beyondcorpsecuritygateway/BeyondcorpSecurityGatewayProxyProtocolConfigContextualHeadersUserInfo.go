// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package beyondcorpsecuritygateway


type BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo struct {
	// The output type of the delegated user info. Possible values: ["PROTOBUF", "JSON", "NONE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/beyondcorp_security_gateway#output_type BeyondcorpSecurityGateway#output_type}
	OutputType *string `field:"optional" json:"outputType" yaml:"outputType"`
}

