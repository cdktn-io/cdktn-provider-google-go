// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicleparser


type ChronicleParserVersionInfo struct {
	// Signifies if the parser is disabled for auto upgrade.
	//
	// If true, the parser
	// will not be upgraded by the auto upgrade process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/chronicle_parser#auto_upgrade_disabled ChronicleParser#auto_upgrade_disabled}
	AutoUpgradeDisabled interface{} `field:"required" json:"autoUpgradeDisabled" yaml:"autoUpgradeDisabled"`
}

