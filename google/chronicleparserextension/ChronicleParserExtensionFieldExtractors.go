// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicleparserextension


type ChronicleParserExtensionFieldExtractors struct {
	// Whether to append repeated fields or not. When false, repeated fields will be replaced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_parser_extension#append_repeated_fields ChronicleParserExtension#append_repeated_fields}
	AppendRepeatedFields interface{} `field:"optional" json:"appendRepeatedFields" yaml:"appendRepeatedFields"`
	// extractors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_parser_extension#extractors ChronicleParserExtension#extractors}
	Extractors interface{} `field:"optional" json:"extractors" yaml:"extractors"`
	// Possible values: JSON CSV XML.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_parser_extension#log_format ChronicleParserExtension#log_format}
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// preprocess_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_parser_extension#preprocess_config ChronicleParserExtension#preprocess_config}
	PreprocessConfig *ChronicleParserExtensionFieldExtractorsPreprocessConfig `field:"optional" json:"preprocessConfig" yaml:"preprocessConfig"`
}

