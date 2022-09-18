//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDashboard) DeepCopyInto(out *ClusterDashboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDashboard.
func (in *ClusterDashboard) DeepCopy() *ClusterDashboard {
	if in == nil {
		return nil
	}
	out := new(ClusterDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterDashboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDashboardList) DeepCopyInto(out *ClusterDashboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterDashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDashboardList.
func (in *ClusterDashboardList) DeepCopy() *ClusterDashboardList {
	if in == nil {
		return nil
	}
	out := new(ClusterDashboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterDashboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDashboardSpec) DeepCopyInto(out *ClusterDashboardSpec) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]DashboardMetrics, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDashboardSpec.
func (in *ClusterDashboardSpec) DeepCopy() *ClusterDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterDashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDashboardStatus) DeepCopyInto(out *ClusterDashboardStatus) {
	*out = *in
	if in.LastReportTime != nil {
		in, out := &in.LastReportTime, &out.LastReportTime
		*out = (*in).DeepCopy()
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make(map[string][]DashboardMetricsValue, len(*in))
		for key, val := range *in {
			var outVal []DashboardMetricsValue
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]DashboardMetricsValue, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDashboardStatus.
func (in *ClusterDashboardStatus) DeepCopy() *ClusterDashboardStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterDashboardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardMetrics) DeepCopyInto(out *DashboardMetrics) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardMetrics.
func (in *DashboardMetrics) DeepCopy() *DashboardMetrics {
	if in == nil {
		return nil
	}
	out := new(DashboardMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardMetricsValue) DeepCopyInto(out *DashboardMetricsValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardMetricsValue.
func (in *DashboardMetricsValue) DeepCopy() *DashboardMetricsValue {
	if in == nil {
		return nil
	}
	out := new(DashboardMetricsValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Meter) DeepCopyInto(out *Meter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Meter.
func (in *Meter) DeepCopy() *Meter {
	if in == nil {
		return nil
	}
	out := new(Meter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricValue) DeepCopyInto(out *MetricValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricValue.
func (in *MetricValue) DeepCopy() *MetricValue {
	if in == nil {
		return nil
	}
	out := new(MetricValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectArchitecture) DeepCopyInto(out *ProjectArchitecture) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectArchitecture.
func (in *ProjectArchitecture) DeepCopy() *ProjectArchitecture {
	if in == nil {
		return nil
	}
	out := new(ProjectArchitecture)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectArchitecture) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectArchitectureList) DeepCopyInto(out *ProjectArchitectureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProjectArchitecture, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectArchitectureList.
func (in *ProjectArchitectureList) DeepCopy() *ProjectArchitectureList {
	if in == nil {
		return nil
	}
	out := new(ProjectArchitectureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectArchitectureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectArchitectureSpec) DeepCopyInto(out *ProjectArchitectureSpec) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make(map[string]ProjectService, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.PrometheusInformation != nil {
		in, out := &in.PrometheusInformation, &out.PrometheusInformation
		*out = make(map[string]PrometheusInfo, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectArchitectureSpec.
func (in *ProjectArchitectureSpec) DeepCopy() *ProjectArchitectureSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectArchitectureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectArchitectureStatus) DeepCopyInto(out *ProjectArchitectureStatus) {
	*out = *in
	if in.PrometheusStatus != nil {
		in, out := &in.PrometheusStatus, &out.PrometheusStatus
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = make(map[string]ServiceMetricsStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectArchitectureStatus.
func (in *ProjectArchitectureStatus) DeepCopy() *ProjectArchitectureStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectArchitectureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectService) DeepCopyInto(out *ProjectService) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProjectServiceItems != nil {
		in, out := &in.ProjectServiceItems, &out.ProjectServiceItems
		*out = make(map[string]ProjectServiceItem, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectService.
func (in *ProjectService) DeepCopy() *ProjectService {
	if in == nil {
		return nil
	}
	out := new(ProjectService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectServiceItem) DeepCopyInto(out *ProjectServiceItem) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceRelations != nil {
		in, out := &in.ServiceRelations, &out.ServiceRelations
		*out = make([]ServiceRelation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]ServiceMetrics, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectServiceItem.
func (in *ProjectServiceItem) DeepCopy() *ProjectServiceItem {
	if in == nil {
		return nil
	}
	out := new(ProjectServiceItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusInfo) DeepCopyInto(out *PrometheusInfo) {
	*out = *in
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusInfo.
func (in *PrometheusInfo) DeepCopy() *PrometheusInfo {
	if in == nil {
		return nil
	}
	out := new(PrometheusInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceItemMetricsStatus) DeepCopyInto(out *ServiceItemMetricsStatus) {
	*out = *in
	if in.VectorData != nil {
		in, out := &in.VectorData, &out.VectorData
		*out = make([]VectorDataLine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VectorMatrix != nil {
		in, out := &in.VectorMatrix, &out.VectorMatrix
		*out = make([]VectorMatrixLine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceItemMetricsStatus.
func (in *ServiceItemMetricsStatus) DeepCopy() *ServiceItemMetricsStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceItemMetricsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceItemStatus) DeepCopyInto(out *ServiceItemStatus) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make(map[string]ServiceItemMetricsStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceItemStatus.
func (in *ServiceItemStatus) DeepCopy() *ServiceItemStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceItemStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMetrics) DeepCopyInto(out *ServiceMetrics) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMetrics.
func (in *ServiceMetrics) DeepCopy() *ServiceMetrics {
	if in == nil {
		return nil
	}
	out := new(ServiceMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMetricsStatus) DeepCopyInto(out *ServiceMetricsStatus) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make(map[string]ServiceItemStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMetricsStatus.
func (in *ServiceMetricsStatus) DeepCopy() *ServiceMetricsStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceMetricsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceRelation) DeepCopyInto(out *ServiceRelation) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceRelation.
func (in *ServiceRelation) DeepCopy() *ServiceRelation {
	if in == nil {
		return nil
	}
	out := new(ServiceRelation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorDataLine) DeepCopyInto(out *VectorDataLine) {
	*out = *in
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Value = in.Value
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorDataLine.
func (in *VectorDataLine) DeepCopy() *VectorDataLine {
	if in == nil {
		return nil
	}
	out := new(VectorDataLine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorMatrixLine) DeepCopyInto(out *VectorMatrixLine) {
	*out = *in
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]MetricValue, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorMatrixLine.
func (in *VectorMatrixLine) DeepCopy() *VectorMatrixLine {
	if in == nil {
		return nil
	}
	out := new(VectorMatrixLine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadList) DeepCopyInto(out *WorkloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadList.
func (in *WorkloadList) DeepCopy() *WorkloadList {
	if in == nil {
		return nil
	}
	out := new(WorkloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadStatus) DeepCopyInto(out *WorkloadStatus) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	in.Timestamp.DeepCopyInto(&out.Timestamp)
	out.Window = in.Window
	if in.Usage != nil {
		in, out := &in.Usage, &out.Usage
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadStatus.
func (in *WorkloadStatus) DeepCopy() *WorkloadStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceDashboard) DeepCopyInto(out *WorkspaceDashboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceDashboard.
func (in *WorkspaceDashboard) DeepCopy() *WorkspaceDashboard {
	if in == nil {
		return nil
	}
	out := new(WorkspaceDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceDashboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceDashboardList) DeepCopyInto(out *WorkspaceDashboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkspaceDashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceDashboardList.
func (in *WorkspaceDashboardList) DeepCopy() *WorkspaceDashboardList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceDashboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceDashboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceDashboardSpec) DeepCopyInto(out *WorkspaceDashboardSpec) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]DashboardMetrics, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceDashboardSpec.
func (in *WorkspaceDashboardSpec) DeepCopy() *WorkspaceDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceDashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceDashboardStatus) DeepCopyInto(out *WorkspaceDashboardStatus) {
	*out = *in
	if in.LastReportTime != nil {
		in, out := &in.LastReportTime, &out.LastReportTime
		*out = (*in).DeepCopy()
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make(map[string][]DashboardMetricsValue, len(*in))
		for key, val := range *in {
			var outVal []DashboardMetricsValue
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]DashboardMetricsValue, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Meters != nil {
		in, out := &in.Meters, &out.Meters
		*out = make([]Meter, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceDashboardStatus.
func (in *WorkspaceDashboardStatus) DeepCopy() *WorkspaceDashboardStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceDashboardStatus)
	in.DeepCopyInto(out)
	return out
}
