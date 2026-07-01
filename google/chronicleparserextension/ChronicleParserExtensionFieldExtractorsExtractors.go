// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicleparserextension


type ChronicleParserExtensionFieldExtractorsExtractors struct {
	// Path in generated event which is to be populated.
	//
	// This is required if the
	// FieldExtractor is used to specify the parser extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/chronicle_parser_extension#destination_path ChronicleParserExtension#destination_path}
	DestinationPath *string `field:"optional" json:"destinationPath" yaml:"destinationPath"`
	// Field path could be a json path, xml path or csv column name depending on log format.
	//
	// It refers to a section or substring in raw log.
	// This is required if the FieldExtractor is used to specify the parser
	// extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/chronicle_parser_extension#field_path ChronicleParserExtension#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
	// Operator used for precondition. Possible values: EQUALS NOT_EQUALS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/chronicle_parser_extension#precondition_op ChronicleParserExtension#precondition_op}
	PreconditionOp *string `field:"optional" json:"preconditionOp" yaml:"preconditionOp"`
	// Precondition path could be a json path, xml path or csv column name depending on log format.
	//
	// It refers to a section or substring in raw log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/chronicle_parser_extension#precondition_path ChronicleParserExtension#precondition_path}
	PreconditionPath *string `field:"optional" json:"preconditionPath" yaml:"preconditionPath"`
	// Precondition value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/chronicle_parser_extension#precondition_value ChronicleParserExtension#precondition_value}
	PreconditionValue *string `field:"optional" json:"preconditionValue" yaml:"preconditionValue"`
	// Value to be mapped to the destination path directly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/chronicle_parser_extension#value ChronicleParserExtension#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

