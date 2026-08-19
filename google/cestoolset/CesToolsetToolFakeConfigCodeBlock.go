// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestoolset


type CesToolsetToolFakeConfigCodeBlock struct {
	// Python code which will be invoked in tool fake mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/ces_toolset#python_code CesToolset#python_code}
	PythonCode *string `field:"required" json:"pythonCode" yaml:"pythonCode"`
}

