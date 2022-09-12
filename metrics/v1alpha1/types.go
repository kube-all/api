/*
Copyright The Kubeall Authors.

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

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="WorkloadKind",type=string,JSONPath=`.status.workloadKind`
// +kubebuilder:printcolumn:name="Replicas",type=integer,JSONPath=`.status.replicas`
// +kubebuilder:printcolumn:name="Timestamp",type="string",JSONPath=`.status.timestamp`
// +kubebuilder:resource:scope=Namespaced

//Workload workload
type Workload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// +optional
	Status WorkloadStatus `json:"status" yaml:"status" protobuf:"bytes,2,opt,name=status"`
}
type WorkloadStatus struct {
	// pod top owner reference's resource kind, eg: Deployment,StatefulSet
	// +optional
	WorkloadKind string `json:"workloadKind" yaml:"workloadKind" protobuf:"bytes,19,opt,name=workloadKind"`
	// replica, DaemonSet's Replicas will equal number of nodes
	// +optional
	Replicas int32 `json:"replicas,omitempty" yaml:"replicas" protobuf:"varint,14,opt,name=replicas"`
	// Total number of non-terminated pods targeted
	// +optional
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty" protobuf:"varint,15,opt,name=updatedReplicas"`
	// Total number of ready pods targeted
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty" protobuf:"varint,16,opt,name=readyReplicas"`
	// Total number of unavailable pods targeted.
	// +optional
	UnavailableReplicas int32 `json:"unavailableReplicas,omitempty" protobuf:"varint,17,opt,name=unavailableReplicas"`
	// one pod's request and limit
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty" yaml:"resources" protobuf:"bytes,18,opt,name=resources"`
	// The following fields define time interval from which metrics were
	// collected from the interval [Timestamp-Window, Timestamp].
	// +optional
	Timestamp metav1.Time `json:"timestamp,omitempty" yaml:"timestamp" protobuf:"bytes,8,opt,name=timestamp"`
	// +optional
	Window metav1.Duration `json:"window,omitempty" yaml:"window" protobuf:"bytes,9,opt,name=window"`
	// The memory usage is the memory working set.
	// +optional
	Usage corev1.ResourceList `json:"usage,omitempty" yaml:"usage" protobuf:"bytes,10,rep,name=usage,casttype=k8s.io/api/core/v1.ResourceList,castkey=k8s.io/api/core/v1.ResourceName"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkloadList contains a list of Workload
type WorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Workload `json:"items" protobuf:"bytes,2,rep,name=items"`
}
