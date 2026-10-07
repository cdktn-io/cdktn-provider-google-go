// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesapp


type CesAppErrorHandlingSettingsEndSessionConfig struct {
	// Whether to escalate the session in EndSession. If session is escalated, metadata in EndSession will contain session_escalated = true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_app#escalate_session CesApp#escalate_session}
	EscalateSession interface{} `field:"optional" json:"escalateSession" yaml:"escalateSession"`
}

