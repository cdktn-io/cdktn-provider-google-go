// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package managedkafkacluster


type ManagedKafkaClusterGcpConfigAccessConfigPublicClusterConfig struct {
	// A list of IPv4 addresses or CIDR ranges that are allowed to connect to the cluster.
	//
	// To protect your cluster, allow access from only trusted external IP ranges. Don't expose your cluster to untrusted ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/managed_kafka_cluster#allowed_source_ip_ranges ManagedKafkaCluster#allowed_source_ip_ranges}
	AllowedSourceIpRanges *[]*string `field:"required" json:"allowedSourceIpRanges" yaml:"allowedSourceIpRanges"`
}

