// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestool


type CesToolPythonFunction struct {
	// The name of the Python function to execute.
	//
	// Must match a Python function
	// name defined in the python code. Case sensitive. If the name is not
	// provided, the first function defined in the python code will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/ces_tool#name CesTool#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Python code to execute for the tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/ces_tool#python_code CesTool#python_code}
	PythonCode *string `field:"optional" json:"pythonCode" yaml:"pythonCode"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/ces_tool#service_directory_config CesTool#service_directory_config}
	ServiceDirectoryConfig *CesToolPythonFunctionServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
}

