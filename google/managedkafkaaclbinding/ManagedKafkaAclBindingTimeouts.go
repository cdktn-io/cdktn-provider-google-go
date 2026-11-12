// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package managedkafkaaclbinding


type ManagedKafkaAclBindingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#create ManagedKafkaAclBinding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#delete ManagedKafkaAclBinding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#update ManagedKafkaAclBinding#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

