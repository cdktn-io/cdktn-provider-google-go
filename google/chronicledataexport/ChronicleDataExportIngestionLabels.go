// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicledataexport


type ChronicleDataExportIngestionLabels struct {
	// The key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_data_export#key ChronicleDataExport#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_data_export#value ChronicleDataExport#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

