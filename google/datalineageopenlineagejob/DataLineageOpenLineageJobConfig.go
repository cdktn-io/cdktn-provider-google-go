// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalineageopenlineagejob

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataLineageOpenLineageJobConfig struct {
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
	// Name of the OpenLineage job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#name DataLineageOpenLineageJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace of the OpenLineage job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#namespace DataLineageOpenLineageJob#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#deletion_policy DataLineageOpenLineageJob#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Description of the OpenLineage job. Changes to description do not trigger updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#description DataLineageOpenLineageJob#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#id DataLineageOpenLineageJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#input DataLineageOpenLineageJob#input}
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#output DataLineageOpenLineageJob#output}
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/data_lineage_open_lineage_job#timeouts DataLineageOpenLineageJob#timeouts}
	Timeouts *DataLineageOpenLineageJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

