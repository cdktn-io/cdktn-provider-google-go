// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package privatecacapool


type PrivatecaCaPoolEncryptionSpec struct {
	// The resource name for an existing Cloud KMS key in the format 'projects/* /locations/* /keyRings/* /cryptoKeys/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/privateca_ca_pool#cloud_kms_key PrivatecaCaPool#cloud_kms_key}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	CloudKmsKey *string `field:"optional" json:"cloudKmsKey" yaml:"cloudKmsKey"`
}

