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
	v1 "k8s.io/api/rbac/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeUser) DeepCopyInto(out *KubeUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeUser.
func (in *KubeUser) DeepCopy() *KubeUser {
	if in == nil {
		return nil
	}
	out := new(KubeUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeUserAPIKey) DeepCopyInto(out *KubeUserAPIKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeUserAPIKey.
func (in *KubeUserAPIKey) DeepCopy() *KubeUserAPIKey {
	if in == nil {
		return nil
	}
	out := new(KubeUserAPIKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeUserAPIKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeUserAPIKeyList) DeepCopyInto(out *KubeUserAPIKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeUserAPIKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeUserAPIKeyList.
func (in *KubeUserAPIKeyList) DeepCopy() *KubeUserAPIKeyList {
	if in == nil {
		return nil
	}
	out := new(KubeUserAPIKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeUserAPIKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeUserAPIKeySpec) DeepCopyInto(out *KubeUserAPIKeySpec) {
	*out = *in
	if in.Expired != nil {
		in, out := &in.Expired, &out.Expired
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeUserAPIKeySpec.
func (in *KubeUserAPIKeySpec) DeepCopy() *KubeUserAPIKeySpec {
	if in == nil {
		return nil
	}
	out := new(KubeUserAPIKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeUserAPIKeyStatus) DeepCopyInto(out *KubeUserAPIKeyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeUserAPIKeyStatus.
func (in *KubeUserAPIKeyStatus) DeepCopy() *KubeUserAPIKeyStatus {
	if in == nil {
		return nil
	}
	out := new(KubeUserAPIKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeUserList) DeepCopyInto(out *KubeUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeUserList.
func (in *KubeUserList) DeepCopy() *KubeUserList {
	if in == nil {
		return nil
	}
	out := new(KubeUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeUserSpec) DeepCopyInto(out *KubeUserSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeUserSpec.
func (in *KubeUserSpec) DeepCopy() *KubeUserSpec {
	if in == nil {
		return nil
	}
	out := new(KubeUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeUserStatus) DeepCopyInto(out *KubeUserStatus) {
	*out = *in
	if in.LastLoginTime != nil {
		in, out := &in.LastLoginTime, &out.LastLoginTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeUserStatus.
func (in *KubeUserStatus) DeepCopy() *KubeUserStatus {
	if in == nil {
		return nil
	}
	out := new(KubeUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserKubeConfig) DeepCopyInto(out *UserKubeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserKubeConfig.
func (in *UserKubeConfig) DeepCopy() *UserKubeConfig {
	if in == nil {
		return nil
	}
	out := new(UserKubeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserKubeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserKubeConfigList) DeepCopyInto(out *UserKubeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserKubeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserKubeConfigList.
func (in *UserKubeConfigList) DeepCopy() *UserKubeConfigList {
	if in == nil {
		return nil
	}
	out := new(UserKubeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserKubeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserKubeConfigSpec) DeepCopyInto(out *UserKubeConfigSpec) {
	*out = *in
	if in.ExpiredTime != nil {
		in, out := &in.ExpiredTime, &out.ExpiredTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserKubeConfigSpec.
func (in *UserKubeConfigSpec) DeepCopy() *UserKubeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(UserKubeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserKubeConfigStatus) DeepCopyInto(out *UserKubeConfigStatus) {
	*out = *in
	if in.EncryptedClientCertificateData != nil {
		in, out := &in.EncryptedClientCertificateData, &out.EncryptedClientCertificateData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.EncryptedClientKeyData != nil {
		in, out := &in.EncryptedClientKeyData, &out.EncryptedClientKeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.LastCheck != nil {
		in, out := &in.LastCheck, &out.LastCheck
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserKubeConfigStatus.
func (in *UserKubeConfigStatus) DeepCopy() *UserKubeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(UserKubeConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceRole) DeepCopyInto(out *WorkspaceRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceRole.
func (in *WorkspaceRole) DeepCopy() *WorkspaceRole {
	if in == nil {
		return nil
	}
	out := new(WorkspaceRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceRoleList) DeepCopyInto(out *WorkspaceRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkspaceRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceRoleList.
func (in *WorkspaceRoleList) DeepCopy() *WorkspaceRoleList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceRoleSpec) DeepCopyInto(out *WorkspaceRoleSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]v1.PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterRoleRefs != nil {
		in, out := &in.ClusterRoleRefs, &out.ClusterRoleRefs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RoleRefs != nil {
		in, out := &in.RoleRefs, &out.RoleRefs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceRoleSpec.
func (in *WorkspaceRoleSpec) DeepCopy() *WorkspaceRoleSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceRoleSpec)
	in.DeepCopyInto(out)
	return out
}
