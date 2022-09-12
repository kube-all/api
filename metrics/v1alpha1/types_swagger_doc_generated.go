/*
Copyright 2022 The kubeall.com Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_Workload = map[string]string{
	"": "Workload workload",
}

func (Workload) SwaggerDoc() map[string]string {
	return map_Workload
}

var map_WorkloadList = map[string]string{
	"": "WorkloadList contains a list of Workload",
}

func (WorkloadList) SwaggerDoc() map[string]string {
	return map_WorkloadList
}

var map_WorkloadStatus = map[string]string{
	"workloadKind":        "pod top owner reference's resource kind, eg: Deployment,StatefulSet",
	"replicas":            "replica, DaemonSet's Replicas will equal number of nodes",
	"updatedReplicas":     "Total number of non-terminated pods targeted",
	"readyReplicas":       "Total number of ready pods targeted",
	"unavailableReplicas": "Total number of unavailable pods targeted.",
	"resources":           "one pod's request and limit",
	"timestamp":           "The following fields define time interval from which metrics were collected from the interval [Timestamp-Window, Timestamp].",
	"usage":               "The memory usage is the memory working set.",
}

func (WorkloadStatus) SwaggerDoc() map[string]string {
	return map_WorkloadStatus
}

// AUTO-GENERATED FUNCTIONS END HERE