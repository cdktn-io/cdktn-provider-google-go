// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerrepository


type SecureSourceManagerRepositoryScanConfig struct {
	// secret_scan_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/secure_source_manager_repository#secret_scan_config SecureSourceManagerRepository#secret_scan_config}
	SecretScanConfig *SecureSourceManagerRepositoryScanConfigSecretScanConfig `field:"optional" json:"secretScanConfig" yaml:"secretScanConfig"`
}

