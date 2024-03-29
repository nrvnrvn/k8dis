// +build !ignore_autogenerated

// Copyright 2019 The redis-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/k8s/v1alpha1.ContainerSpec": schema_pkg_apis_k8s_v1alpha1_ContainerSpec(ref),
		"./pkg/apis/k8s/v1alpha1.Password":      schema_pkg_apis_k8s_v1alpha1_Password(ref),
		"./pkg/apis/k8s/v1alpha1.Redis":         schema_pkg_apis_k8s_v1alpha1_Redis(ref),
		"./pkg/apis/k8s/v1alpha1.RedisList":     schema_pkg_apis_k8s_v1alpha1_RedisList(ref),
		"./pkg/apis/k8s/v1alpha1.RedisSpec":     schema_pkg_apis_k8s_v1alpha1_RedisSpec(ref),
		"./pkg/apis/k8s/v1alpha1.RedisStatus":   schema_pkg_apis_k8s_v1alpha1_RedisStatus(ref),
	}
}

func schema_pkg_apis_k8s_v1alpha1_ContainerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ContainerSpec allows to set some container-specific attributes",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image is a standard path for a Container image",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Resources describes the compute resource requirements",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"securityContext": {
						SchemaProps: spec.SchemaProps{
							Description: "SecurityContext holds security configuration that will be applied to a container",
							Ref:         ref("k8s.io/api/core/v1.SecurityContext"),
						},
					},
					"initialDelaySeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"image"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.SecurityContext"},
	}
}

func schema_pkg_apis_k8s_v1alpha1_Password(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Password allows to refer to a Secret containing password for Redis Password should be strong enough. Passwords shorter than 8 characters composed of ASCII alphanumeric symbols will lead to a mild warning logged by the Operator. Please note that password hashes are added as annotations to Pods to enable password rotation. Hashes are generated using argon2id KDF. Changing the password in the referenced Secret will not trigger the rolling Statefulset upgrade automatically. However an event in regard to any objects owned by the Redis resource fired afterwards will trigger the rolling upgrade. Redis operator does not store the password internally and reads it from the Secret any time the Reconcile is called. Hence it will not be able to connect to Pods with the ``old'' password. In scenarios when persistence is turned off all the data will be lost during password rotation.",
				Properties: map[string]spec.Schema{
					"secretKeyRef": {
						SchemaProps: spec.SchemaProps{
							Description: "SecretKeyRef is a reference to the Secret in the same namespace containing the password.",
							Ref:         ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
				},
				Required: []string{"secretKeyRef"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.SecretKeySelector"},
	}
}

func schema_pkg_apis_k8s_v1alpha1_Redis(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Redis is the Schema for the redis API",
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
							Description: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/k8s/v1alpha1.RedisSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/k8s/v1alpha1.RedisStatus"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/k8s/v1alpha1.RedisSpec", "./pkg/apis/k8s/v1alpha1.RedisStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_k8s_v1alpha1_RedisList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RedisList is a list of Redis resources",
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
					"items": {
						SchemaProps: spec.SchemaProps{
							Description: "List of Redis resources",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/k8s/v1alpha1.Redis"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/k8s/v1alpha1.Redis"},
	}
}

func schema_pkg_apis_k8s_v1alpha1_RedisSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RedisSpec defines the desired state of Redis",
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replicas is a number of replicas in a Redis failover cluster",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Description: "Config allows to pass custom Redis configuration parameters",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"password": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/k8s/v1alpha1.Password"),
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Description: "Pod annotations",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"securityContext": {
						SchemaProps: spec.SchemaProps{
							Description: "Pod securityContext",
							Ref:         ref("k8s.io/api/core/v1.PodSecurityContext"),
						},
					},
					"affinity": {
						SchemaProps: spec.SchemaProps{
							Description: "Pod affinity",
							Ref:         ref("k8s.io/api/core/v1.Affinity"),
						},
					},
					"tolerations": {
						SchemaProps: spec.SchemaProps{
							Description: "Pod tolerations",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"serviceAccountName": {
						SchemaProps: spec.SchemaProps{
							Description: "Pod ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imagePullSecrets": {
						SchemaProps: spec.SchemaProps{
							Description: "Pod ImagePullSecrets More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.LocalObjectReference"),
									},
								},
							},
						},
					},
					"priorityClassName": {
						SchemaProps: spec.SchemaProps{
							Description: "Pod priorityClassName",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dataVolumeClaimTemplate": {
						SchemaProps: spec.SchemaProps{
							Description: "DataVolumeClaimTemplate for StatefulSet",
							Ref:         ref("k8s.io/api/core/v1.PersistentVolumeClaim"),
						},
					},
					"volumes": {
						SchemaProps: spec.SchemaProps{
							Description: "Volumes for StatefulSet",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Volume"),
									},
								},
							},
						},
					},
					"redis": {
						SchemaProps: spec.SchemaProps{
							Description: "Redis container specification",
							Ref:         ref("./pkg/apis/k8s/v1alpha1.ContainerSpec"),
						},
					},
					"exporter": {
						SchemaProps: spec.SchemaProps{
							Description: "Exporter container specification",
							Ref:         ref("./pkg/apis/k8s/v1alpha1.ContainerSpec"),
						},
					},
					"initContainers": {
						SchemaProps: spec.SchemaProps{
							Description: "Pod initContainers",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Container"),
									},
								},
							},
						},
					},
				},
				Required: []string{"replicas", "redis"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/k8s/v1alpha1.ContainerSpec", "./pkg/apis/k8s/v1alpha1.Password", "k8s.io/api/core/v1.Affinity", "k8s.io/api/core/v1.Container", "k8s.io/api/core/v1.LocalObjectReference", "k8s.io/api/core/v1.PersistentVolumeClaim", "k8s.io/api/core/v1.PodSecurityContext", "k8s.io/api/core/v1.Toleration", "k8s.io/api/core/v1.Volume"},
	}
}

func schema_pkg_apis_k8s_v1alpha1_RedisStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RedisStatus contains the observed state of Redis",
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replicas is the number of active Redis instances in the replication",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"master": {
						SchemaProps: spec.SchemaProps{
							Description: "Master is the current master's Pod name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"replicas", "master"},
			},
		},
		Dependencies: []string{},
	}
}
