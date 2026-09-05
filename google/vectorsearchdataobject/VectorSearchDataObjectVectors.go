// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vectorsearchdataobject


type VectorSearchDataObjectVectors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/vector_search_data_object#field_name VectorSearchDataObject#field_name}.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// dense block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/vector_search_data_object#dense VectorSearchDataObject#dense}
	Dense *VectorSearchDataObjectVectorsDense `field:"optional" json:"dense" yaml:"dense"`
	// sparse block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/vector_search_data_object#sparse VectorSearchDataObject#sparse}
	Sparse *VectorSearchDataObjectVectorsSparse `field:"optional" json:"sparse" yaml:"sparse"`
}

