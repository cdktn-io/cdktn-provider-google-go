// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate


type TranscoderJobTemplateConfigEncryptionsMpegCenc struct {
	// Specify the encryption scheme.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/transcoder_job_template#scheme TranscoderJobTemplate#scheme}
	Scheme *string `field:"required" json:"scheme" yaml:"scheme"`
}

