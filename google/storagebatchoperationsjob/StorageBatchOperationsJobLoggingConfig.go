// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagebatchoperationsjob


type StorageBatchOperationsJobLoggingConfig struct {
	// Specifies the actions to be logged. Possible values include 'TRANSFORM'. See [LoggingConfig](https://cloud.google.com/storage/docs/batch-operations/reference/rest/v1/projects.locations.jobs#LoggingConfig) for allowed values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/storage_batch_operations_job#log_actions StorageBatchOperationsJob#log_actions}
	LogActions *[]*string `field:"required" json:"logActions" yaml:"logActions"`
	// States in which Action are logged.If empty, no logs are generated. Possible values include 'SUCCEEDED', 'FAILED'. See [LoggingConfig](https://cloud.google.com/storage/docs/batch-operations/reference/rest/v1/projects.locations.jobs#LoggingConfig) for allowed values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/storage_batch_operations_job#log_action_states StorageBatchOperationsJob#log_action_states}
	LogActionStates *[]*string `field:"required" json:"logActionStates" yaml:"logActionStates"`
}

