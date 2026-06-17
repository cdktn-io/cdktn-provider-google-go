// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestool


type CesToolDataStoreToolEngineSourceDataStoreSources struct {
	// data_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/ces_tool#data_store CesTool#data_store}
	DataStore *CesToolDataStoreToolEngineSourceDataStoreSourcesDataStore `field:"optional" json:"dataStore" yaml:"dataStore"`
	// Filter specification for the DataStore. See: https://cloud.google.com/generative-ai-app-builder/docs/filter-search-metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/ces_tool#filter CesTool#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

