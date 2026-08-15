// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package firestoreindex


type FirestoreIndexFieldsSearchConfig struct {
	// geo_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/firestore_index#geo_spec FirestoreIndex#geo_spec}
	GeoSpec *FirestoreIndexFieldsSearchConfigGeoSpec `field:"optional" json:"geoSpec" yaml:"geoSpec"`
	// text_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/firestore_index#text_spec FirestoreIndex#text_spec}
	TextSpec *FirestoreIndexFieldsSearchConfigTextSpec `field:"optional" json:"textSpec" yaml:"textSpec"`
}

