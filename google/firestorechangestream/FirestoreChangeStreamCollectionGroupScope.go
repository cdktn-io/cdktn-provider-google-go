// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package firestorechangestream


type FirestoreChangeStreamCollectionGroupScope struct {
	// The ID of the collection group to track.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/firestore_change_stream#collection_group_id FirestoreChangeStream#collection_group_id}
	CollectionGroupId *string `field:"required" json:"collectionGroupId" yaml:"collectionGroupId"`
}

