// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kmsprojectautokeyconfig


type KmsProjectAutokeyConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/kms_project_autokey_config#create KmsProjectAutokeyConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/kms_project_autokey_config#delete KmsProjectAutokeyConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/kms_project_autokey_config#update KmsProjectAutokeyConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

