// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline


type VertexAiTrainingPipelineModelToUploadContainerSpec struct {
	// URI of the Docker image to be used as the custom container for serving predictions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#image_uri VertexAiTrainingPipeline#image_uri}
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// Specifies the arguments to the command when the container starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#args VertexAiTrainingPipeline#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Specifies the command that runs when the container starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#command VertexAiTrainingPipeline#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Deployment timeout. Limit for deployment timeout is 2 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#deployment_timeout VertexAiTrainingPipeline#deployment_timeout}
	DeploymentTimeout *string `field:"optional" json:"deploymentTimeout" yaml:"deploymentTimeout"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#env VertexAiTrainingPipeline#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// grpc_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#grpc_ports VertexAiTrainingPipeline#grpc_ports}
	GrpcPorts interface{} `field:"optional" json:"grpcPorts" yaml:"grpcPorts"`
	// health_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#health_probe VertexAiTrainingPipeline#health_probe}
	HealthProbe *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbe `field:"optional" json:"healthProbe" yaml:"healthProbe"`
	// HTTP path on the container to send health checks to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#health_route VertexAiTrainingPipeline#health_route}
	HealthRoute *string `field:"optional" json:"healthRoute" yaml:"healthRoute"`
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#ports VertexAiTrainingPipeline#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
	// HTTP path on the container to send prediction requests to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#predict_route VertexAiTrainingPipeline#predict_route}
	PredictRoute *string `field:"optional" json:"predictRoute" yaml:"predictRoute"`
	// The amount of the VM memory to reserve as the shared memory for the model in megabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#shared_memory_size_mb VertexAiTrainingPipeline#shared_memory_size_mb}
	SharedMemorySizeMb *string `field:"optional" json:"sharedMemorySizeMb" yaml:"sharedMemorySizeMb"`
	// startup_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#startup_probe VertexAiTrainingPipeline#startup_probe}
	StartupProbe *VertexAiTrainingPipelineModelToUploadContainerSpecStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
}

