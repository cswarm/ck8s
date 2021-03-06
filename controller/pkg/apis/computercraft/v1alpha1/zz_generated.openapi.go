// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/computercraft/v1alpha1.Computer":                 schema_pkg_apis_computercraft_v1alpha1_Computer(ref),
		"./pkg/apis/computercraft/v1alpha1.ComputerDeployment":       schema_pkg_apis_computercraft_v1alpha1_ComputerDeployment(ref),
		"./pkg/apis/computercraft/v1alpha1.ComputerDeploymentSpec":   schema_pkg_apis_computercraft_v1alpha1_ComputerDeploymentSpec(ref),
		"./pkg/apis/computercraft/v1alpha1.ComputerDeploymentStatus": schema_pkg_apis_computercraft_v1alpha1_ComputerDeploymentStatus(ref),
		"./pkg/apis/computercraft/v1alpha1.ComputerPod":              schema_pkg_apis_computercraft_v1alpha1_ComputerPod(ref),
		"./pkg/apis/computercraft/v1alpha1.ComputerPodSpec":          schema_pkg_apis_computercraft_v1alpha1_ComputerPodSpec(ref),
		"./pkg/apis/computercraft/v1alpha1.ComputerSpec":             schema_pkg_apis_computercraft_v1alpha1_ComputerSpec(ref),
	}
}

func schema_pkg_apis_computercraft_v1alpha1_Computer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Computer is the Schema for the computers API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/computercraft/v1alpha1.ComputerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.NodeStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/computercraft/v1alpha1.ComputerSpec", "k8s.io/api/core/v1.NodeStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_computercraft_v1alpha1_ComputerDeployment(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ComputerDeployment is the Schema for the computerdeployments API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/computercraft/v1alpha1.ComputerDeploymentSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/computercraft/v1alpha1.ComputerDeploymentStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/computercraft/v1alpha1.ComputerDeploymentSpec", "./pkg/apis/computercraft/v1alpha1.ComputerDeploymentStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_computercraft_v1alpha1_ComputerDeploymentSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ComputerDeploymentSpec defines the desired state of ComputerDeployment",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_computercraft_v1alpha1_ComputerDeploymentStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ComputerDeploymentStatus defines the observed state of ComputerDeployment",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_computercraft_v1alpha1_ComputerPod(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ComputerPod is the Schema for the computerpods API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/computercraft/v1alpha1.ComputerPodSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/computercraft/v1alpha1.ComputerPodStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/computercraft/v1alpha1.ComputerPodSpec", "./pkg/apis/computercraft/v1alpha1.ComputerPodStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_computercraft_v1alpha1_ComputerPodSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ComputerPodSpec defines the desired state of ComputerPod",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"computerType": {
						SchemaProps: spec.SchemaProps{
							Description: "ComputerType is a type of computer to run on",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "URL is a URL to download source code from",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_computercraft_v1alpha1_ComputerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ComputerSpec defines the desired state of Computer",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"id": {
						SchemaProps: spec.SchemaProps{
							Description: "ID of the computer",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"id"},
			},
		},
	}
}
