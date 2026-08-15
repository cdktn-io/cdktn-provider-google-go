// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestool


type CesToolToolFakeConfigCodeBlock struct {
	// Python code which will be invoked in tool fake mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/ces_tool#python_code CesTool#python_code}
	PythonCode *string `field:"required" json:"pythonCode" yaml:"pythonCode"`
}

