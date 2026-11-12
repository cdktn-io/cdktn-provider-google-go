// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute


type NetworkServicesHttpRouteRulesActionDirectResponse struct {
	// Status to return as part of HTTP Response. Must be a positive integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/network_services_http_route#status NetworkServicesHttpRoute#status}
	Status *float64 `field:"required" json:"status" yaml:"status"`
	// Response body as base64 encoded bytes. Maximum body size is 4096B.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/network_services_http_route#bytes_body NetworkServicesHttpRoute#bytes_body}
	BytesBody *string `field:"optional" json:"bytesBody" yaml:"bytesBody"`
	// Response body as a string. Maximum body length is 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/network_services_http_route#string_body NetworkServicesHttpRoute#string_body}
	StringBody *string `field:"optional" json:"stringBody" yaml:"stringBody"`
}

