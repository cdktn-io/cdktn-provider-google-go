// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine


type VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentParts struct {
	// audio_transcription block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#audio_transcription VertexAiReasoningEngine#audio_transcription}
	AudioTranscription *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsAudioTranscription `field:"optional" json:"audioTranscription" yaml:"audioTranscription"`
	// code_execution_result block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#code_execution_result VertexAiReasoningEngine#code_execution_result}
	CodeExecutionResult *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsCodeExecutionResult `field:"optional" json:"codeExecutionResult" yaml:"codeExecutionResult"`
	// executable_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#executable_code VertexAiReasoningEngine#executable_code}
	ExecutableCode *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsExecutableCode `field:"optional" json:"executableCode" yaml:"executableCode"`
	// file_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#file_data VertexAiReasoningEngine#file_data}
	FileData *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsFileData `field:"optional" json:"fileData" yaml:"fileData"`
	// function_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#function_call VertexAiReasoningEngine#function_call}
	FunctionCall *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsFunctionCall `field:"optional" json:"functionCall" yaml:"functionCall"`
	// function_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#function_response VertexAiReasoningEngine#function_response}
	FunctionResponse *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsFunctionResponse `field:"optional" json:"functionResponse" yaml:"functionResponse"`
	// inline_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#inline_data VertexAiReasoningEngine#inline_data}
	InlineData *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsInlineData `field:"optional" json:"inlineData" yaml:"inlineData"`
	// The text content of the part.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#text VertexAiReasoningEngine#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// Indicates whether the part represents the model's thought process or reasoning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#thought VertexAiReasoningEngine#thought}
	Thought interface{} `field:"optional" json:"thought" yaml:"thought"`
	// video_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_reasoning_engine#video_metadata VertexAiReasoningEngine#video_metadata}
	VideoMetadata *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsVideoMetadata `field:"optional" json:"videoMetadata" yaml:"videoMetadata"`
}

