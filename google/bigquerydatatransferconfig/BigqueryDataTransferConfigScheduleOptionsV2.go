// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigquerydatatransferconfig


type BigqueryDataTransferConfigScheduleOptionsV2 struct {
	// event_driven_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/bigquery_data_transfer_config#event_driven_schedule BigqueryDataTransferConfig#event_driven_schedule}
	EventDrivenSchedule *BigqueryDataTransferConfigScheduleOptionsV2EventDrivenSchedule `field:"required" json:"eventDrivenSchedule" yaml:"eventDrivenSchedule"`
}

