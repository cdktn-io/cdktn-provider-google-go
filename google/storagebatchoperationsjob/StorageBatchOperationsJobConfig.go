// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagebatchoperationsjob

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StorageBatchOperationsJobConfig struct {
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
	// bucket_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/storage_batch_operations_job#bucket_list StorageBatchOperationsJob#bucket_list}
	BucketList *StorageBatchOperationsJobBucketListStruct `field:"optional" json:"bucketList" yaml:"bucketList"`
	// delete_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/storage_batch_operations_job#delete_object StorageBatchOperationsJob#delete_object}
	DeleteObject *StorageBatchOperationsJobDeleteObject `field:"optional" json:"deleteObject" yaml:"deleteObject"`
	// If set to 'true', the storage batch operation job will not be deleted and new job will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/storage_batch_operations_job#delete_protection StorageBatchOperationsJob#delete_protection}
	DeleteProtection interface{} `field:"optional" json:"deleteProtection" yaml:"deleteProtection"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/storage_batch_operations_job#deletion_policy StorageBatchOperationsJob#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/storage_batch_operations_job#description StorageBatchOperationsJob#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/storage_batch_operations_job#id StorageBatchOperationsJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/storage_batch_operations_job#job_id StorageBatchOperationsJob#job_id}
	JobId *string `field:"optional" json:"jobId" yaml:"jobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/storage_batch_operations_job#project StorageBatchOperationsJob#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// put_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/storage_batch_operations_job#put_metadata StorageBatchOperationsJob#put_metadata}
	PutMetadata *StorageBatchOperationsJobPutMetadata `field:"optional" json:"putMetadata" yaml:"putMetadata"`
	// put_object_hold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/storage_batch_operations_job#put_object_hold StorageBatchOperationsJob#put_object_hold}
	PutObjectHold *StorageBatchOperationsJobPutObjectHold `field:"optional" json:"putObjectHold" yaml:"putObjectHold"`
	// rewrite_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/storage_batch_operations_job#rewrite_object StorageBatchOperationsJob#rewrite_object}
	RewriteObject *StorageBatchOperationsJobRewriteObject `field:"optional" json:"rewriteObject" yaml:"rewriteObject"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/storage_batch_operations_job#timeouts StorageBatchOperationsJob#timeouts}
	Timeouts *StorageBatchOperationsJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

