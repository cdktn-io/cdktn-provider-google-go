// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefindingsrefinementdeployment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFindingsRefinementDeploymentConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_findings_refinement_deployment#findings_refinement ChronicleFindingsRefinementDeployment#findings_refinement}
	FindingsRefinement *string `field:"required" json:"findingsRefinement" yaml:"findingsRefinement"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_findings_refinement_deployment#instance ChronicleFindingsRefinementDeployment#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_findings_refinement_deployment#location ChronicleFindingsRefinementDeployment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The archive state of the findings refinement deployment.
	//
	// Cannot be set to true unless enabled is set to false.
	// If currently set to true, enabled cannot be updated to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_findings_refinement_deployment#archived ChronicleFindingsRefinementDeployment#archived}
	Archived interface{} `field:"optional" json:"archived" yaml:"archived"`
	// detection_exclusion_application block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_findings_refinement_deployment#detection_exclusion_application ChronicleFindingsRefinementDeployment#detection_exclusion_application}
	DetectionExclusionApplication *ChronicleFindingsRefinementDeploymentDetectionExclusionApplication `field:"optional" json:"detectionExclusionApplication" yaml:"detectionExclusionApplication"`
	// Whether the findings refinement is currently deployed continuously against incoming findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_findings_refinement_deployment#enabled ChronicleFindingsRefinementDeployment#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_findings_refinement_deployment#id ChronicleFindingsRefinementDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_findings_refinement_deployment#project ChronicleFindingsRefinementDeployment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/chronicle_findings_refinement_deployment#timeouts ChronicleFindingsRefinementDeployment#timeouts}
	Timeouts *ChronicleFindingsRefinementDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

