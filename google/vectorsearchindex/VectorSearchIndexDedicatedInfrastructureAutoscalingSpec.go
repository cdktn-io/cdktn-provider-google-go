// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vectorsearchindex


type VectorSearchIndexDedicatedInfrastructureAutoscalingSpec struct {
	// The maximum number of replicas.
	//
	// Must be >= 'min_replica_count'
	// and <= '1000'. If not set or set to '0', defaults to the greater
	// of 'min_replica_count' and '2' (or '5' for the v1beta version).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/vector_search_index#max_replica_count VectorSearchIndex#max_replica_count}
	MaxReplicaCount *float64 `field:"optional" json:"maxReplicaCount" yaml:"maxReplicaCount"`
	// The minimum number of replicas.
	//
	// If not set or set to '0', defaults
	// to '2'. Must be >= '1' and <= '1000'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/vector_search_index#min_replica_count VectorSearchIndex#min_replica_count}
	MinReplicaCount *float64 `field:"optional" json:"minReplicaCount" yaml:"minReplicaCount"`
}

