// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclereferencelist


type ChronicleReferenceListScopeInfo struct {
	// reference_list_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_reference_list#reference_list_scope ChronicleReferenceList#reference_list_scope}
	ReferenceListScope *ChronicleReferenceListScopeInfoReferenceListScope `field:"optional" json:"referenceListScope" yaml:"referenceListScope"`
}

