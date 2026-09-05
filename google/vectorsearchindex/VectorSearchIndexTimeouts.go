// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vectorsearchindex


type VectorSearchIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/vector_search_index#create VectorSearchIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/vector_search_index#delete VectorSearchIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/vector_search_index#update VectorSearchIndex#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

