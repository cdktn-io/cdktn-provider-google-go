// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package firestoreindex


type FirestoreIndexSearchIndexOptions struct {
	// Optional. The language code used to parse all text fields included in the search index, for example, "en", "es", "jp".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/firestore_index#text_language FirestoreIndex#text_language}
	TextLanguage *string `field:"optional" json:"textLanguage" yaml:"textLanguage"`
	// Optional. The path to the field that records the text language override of each document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/firestore_index#text_language_override_field_path FirestoreIndex#text_language_override_field_path}
	TextLanguageOverrideFieldPath *string `field:"optional" json:"textLanguageOverrideFieldPath" yaml:"textLanguageOverrideFieldPath"`
}

