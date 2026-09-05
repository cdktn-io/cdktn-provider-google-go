// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vectorsearchdataobject


type VectorSearchDataObjectVectorsSparse struct {
	// The indices corresponding to the entries in 'values'. Must have the same length as 'values'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/vector_search_data_object#indices VectorSearchDataObject#indices}
	Indices *[]*float64 `field:"required" json:"indices" yaml:"indices"`
	// The non-zero float values of the sparse vector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/vector_search_data_object#values VectorSearchDataObject#values}
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

