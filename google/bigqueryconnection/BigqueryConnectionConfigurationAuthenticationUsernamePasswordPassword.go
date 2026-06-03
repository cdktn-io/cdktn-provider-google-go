// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigqueryconnection


type BigqueryConnectionConfigurationAuthenticationUsernamePasswordPassword struct {
	// The plaintext password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/bigquery_connection#plaintext BigqueryConnection#plaintext}
	Plaintext *string `field:"required" json:"plaintext" yaml:"plaintext"`
}

