// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestool


type CesToolDataStoreToolModalityConfigsSnippetsConfig struct {
	// Whether snippets are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/ces_tool#enable_snippets CesTool#enable_snippets}
	EnableSnippets interface{} `field:"optional" json:"enableSnippets" yaml:"enableSnippets"`
	// Number of snippets to return per query. If unset, returns all snippets from the service by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/ces_tool#max_snippets CesTool#max_snippets}
	MaxSnippets *float64 `field:"optional" json:"maxSnippets" yaml:"maxSnippets"`
}

