// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataplexentrylink


type DataplexEntryLinkAspects struct {
	// aspect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/dataplex_entry_link#aspect DataplexEntryLink#aspect}
	Aspect *DataplexEntryLinkAspectsAspect `field:"required" json:"aspect" yaml:"aspect"`
	// The map keys of the Aspects which the service should modify.
	//
	// It should be the aspect type reference in the format '{project_number}.{location_id}.{aspect_type_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/dataplex_entry_link#aspect_key DataplexEntryLink#aspect_key}
	AspectKey *string `field:"required" json:"aspectKey" yaml:"aspectKey"`
}

