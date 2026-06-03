// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loggingsavedquery


type LoggingSavedQueryLoggingQuerySummaryFields struct {
	// The field from the LogEntry to include in the summary line.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/logging_saved_query#field LoggingSavedQuery#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
}

