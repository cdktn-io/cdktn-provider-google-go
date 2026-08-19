// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package colabnotebookexecution


type ColabNotebookExecutionCustomEnvironmentSpecShieldedInstanceConfig struct {
	// Defines whether the instance has integrity monitoring enabled.
	//
	// Enables monitoring and attestation of the boot integrity of the instance. The attestation is performed against the integrity policy baseline. This baseline is initially derived from the implicitly trusted boot image when the instance is created. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/colab_notebook_execution#enable_integrity_monitoring ColabNotebookExecution#enable_integrity_monitoring}
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Defines whether the instance has Secure Boot enabled.
	//
	// Secure Boot helps ensure that the system only runs authentic software by verifying the digital signature of all boot components, and halting the boot process if signature verification fails. Disabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/colab_notebook_execution#enable_secure_boot ColabNotebookExecution#enable_secure_boot}
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Defines whether the instance has the vTPM enabled. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/colab_notebook_execution#enable_vtpm ColabNotebookExecution#enable_vtpm}
	EnableVtpm interface{} `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}

