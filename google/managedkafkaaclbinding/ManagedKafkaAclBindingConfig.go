// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package managedkafkaaclbinding

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ManagedKafkaAclBindingConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the acl this entry belongs to, which is the final component of the acl's name.
	//
	// The structure of 'aclId' defines the Resource
	// Pattern (resource_type, resource_name, pattern_type) of the acl.
	// 'aclId' is structured like one of the following:
	// For acls on the cluster: 'cluster'
	// For acls on a single resource within the cluster: 'topic/{resource_name}' 'consumerGroup/{resource_name}' 'transactionalId/{resource_name}'
	// For acls on all resources that match a prefix: 'topicPrefixed/{resource_name}' 'consumerGroupPrefixed/{resource_name}' 'transactionalIdPrefixed/{resource_name}'
	// For acls on all resources of a given type (i.e. the wildcard literal '*''): 'allTopics' (represents 'topic/*') 'allConsumerGroups' (represents 'consumerGroup/*') 'allTransactionalIds' (represents 'transactionalId/*').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#acl_id ManagedKafkaAclBinding#acl_id}
	AclId *string `field:"required" json:"aclId" yaml:"aclId"`
	// The cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#cluster ManagedKafkaAclBinding#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// ID of the location of the Kafka resource. See https://cloud.google.com/managed-kafka/docs/locations for a list of supported locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#location ManagedKafkaAclBinding#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The operation type.
	//
	// Allowed values are (case insensitive): ALL, READ,
	// WRITE, CREATE, DELETE, ALTER, DESCRIBE, CLUSTER_ACTION, DESCRIBE_CONFIGS,
	// ALTER_CONFIGS, and IDEMPOTENT_WRITE. See
	// https://kafka.apache.org/documentation/#operations_resources_and_protocols
	// for valid combinations of resource_type and operation for different
	// Kafka API requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#operation ManagedKafkaAclBinding#operation}
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// The principal.
	//
	// Specified as a Google Cloud account, with the Kafka
	// StandardAuthorizer prefix "User:". For example:
	// "User:test-kafka-client@test-project.iam.gserviceaccount.com". Can be
	// the wildcard "User:*" to refer to all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#principal ManagedKafkaAclBinding#principal}
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// Defaults to "DELETE".
	// When a 'terraform destroy' or 'terraform apply' would delete the instance,
	// the command will fail if this field is set to "PREVENT" in Terraform state.
	// When set to "ABANDON", the command will remove the resource from Terraform
	// management without updating or deleting the resource in the API.
	// When set to "DELETE", deleting the resource is allowed.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#deletion_policy ManagedKafkaAclBinding#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The host. Must be set to "*" for Managed Service for Apache Kafka.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#host ManagedKafkaAclBinding#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#id ManagedKafkaAclBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The permission type. Accepted values are (case insensitive): ALLOW, DENY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#permission_type ManagedKafkaAclBinding#permission_type}
	PermissionType *string `field:"optional" json:"permissionType" yaml:"permissionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#project ManagedKafkaAclBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/managed_kafka_acl_binding#timeouts ManagedKafkaAclBinding#timeouts}
	Timeouts *ManagedKafkaAclBindingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

