// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclereferencelist


type ChronicleReferenceListScopeInfoReferenceListScope struct {
	// Optional.
	//
	// The list of scope names of the reference list. The scope names should be
	// full resource names and should be of the format:
	// "projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope_name}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/chronicle_reference_list#scope_names ChronicleReferenceList#scope_names}
	ScopeNames *[]*string `field:"optional" json:"scopeNames" yaml:"scopeNames"`
}

