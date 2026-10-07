// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowtool


type DialogflowToolFunctionSpec struct {
	// Optional.
	//
	// The JSON schema is encapsulated in a [google.protobuf.Struct](https://protobuf.dev/reference/protobuf/google.protobuf/#struct) to describe the input of the function.
	// This input is a JSON object that contains the function's parameters as properties of the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_tool#input_schema DialogflowTool#input_schema}
	InputSchema *string `field:"optional" json:"inputSchema" yaml:"inputSchema"`
	// Optional.
	//
	// The method type of the function. If not specified, the default value is GET. Possible values: ["GET", "POST", "PUT", "DELETE", "PATCH"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_tool#method_type DialogflowTool#method_type}
	MethodType *string `field:"optional" json:"methodType" yaml:"methodType"`
	// Optional.
	//
	// The JSON schema is encapsulated in a [google.protobuf.Struct](https://protobuf.dev/reference/protobuf/google.protobuf/#struct) to describe the output of the function.
	// This output is a JSON object that contains the function's parameters as properties of the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_tool#output_schema DialogflowTool#output_schema}
	OutputSchema *string `field:"optional" json:"outputSchema" yaml:"outputSchema"`
}

