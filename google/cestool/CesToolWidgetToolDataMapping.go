// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestool


type CesToolWidgetToolDataMapping struct {
	// Optional.
	//
	// A map of widget input parameter fields to the corresponding output fields of the source tool.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/ces_tool#field_mappings CesTool#field_mappings}
	FieldMappings *map[string]*string `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// Optional. The mode of the data mapping. Possible values: MODE_UNSPECIFIED FIELD_MAPPING PYTHON_SCRIPT Possible values: ["MODE_UNSPECIFIED", "FIELD_MAPPING", "PYTHON_SCRIPT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/ces_tool#mode CesTool#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// python_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/ces_tool#python_function CesTool#python_function}
	PythonFunction *CesToolWidgetToolDataMappingPythonFunction `field:"optional" json:"pythonFunction" yaml:"pythonFunction"`
	// Optional.
	//
	// The resource name of the tool that provides the data for the widget (e.g., a search tool or a custom function).
	// Format: projects/{project}/locations/{location}/agents/{agent}/tools/{tool}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/ces_tool#source_tool_name CesTool#source_tool_name}
	SourceToolName *string `field:"optional" json:"sourceToolName" yaml:"sourceToolName"`
}

