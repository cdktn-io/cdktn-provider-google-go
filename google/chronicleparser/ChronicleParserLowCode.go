// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicleparser


type ChronicleParserLowCode struct {
	// field_extractors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_parser#field_extractors ChronicleParser#field_extractors}
	FieldExtractors *ChronicleParserLowCodeFieldExtractors `field:"optional" json:"fieldExtractors" yaml:"fieldExtractors"`
	// The log used to create this low code parser in the UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/chronicle_parser#log ChronicleParser#log}
	Log *string `field:"optional" json:"log" yaml:"log"`
}

