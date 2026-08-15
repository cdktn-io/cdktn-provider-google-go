// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerrepository


type SecureSourceManagerRepositoryScanConfigSecretScanConfig struct {
	// Enables secret scanning for the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/secure_source_manager_repository#enabled SecureSourceManagerRepository#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The DLP inspect template to use for secret scanning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/secure_source_manager_repository#inspect_template SecureSourceManagerRepository#inspect_template}
	InspectTemplate *string `field:"optional" json:"inspectTemplate" yaml:"inspectTemplate"`
}

