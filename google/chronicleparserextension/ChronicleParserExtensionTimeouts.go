// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicleparserextension


type ChronicleParserExtensionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_parser_extension#create ChronicleParserExtension#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_parser_extension#delete ChronicleParserExtension#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

