// +build !ignore_autogenerated

// Tencent is pleased to support the open source community by making TKEStack
// available.

// Copyright (C) 2012-2020 Tencent. All Rights Reserved.

// Licensed under the Apache License, Version 2.0 (the "License"); you may not use
// this file except in compliance with the License. You may obtain a copy of the
// License at

// https://opensource.org/licenses/Apache-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OF ANY KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterGateway) DeepCopyInto(out *JupyterGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterGateway.
func (in *JupyterGateway) DeepCopy() *JupyterGateway {
	if in == nil {
		return nil
	}
	out := new(JupyterGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JupyterGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterGatewayList) DeepCopyInto(out *JupyterGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JupyterGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterGatewayList.
func (in *JupyterGatewayList) DeepCopy() *JupyterGatewayList {
	if in == nil {
		return nil
	}
	out := new(JupyterGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JupyterGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterGatewaySpec) DeepCopyInto(out *JupyterGatewaySpec) {
	*out = *in
	if in.Kernels != nil {
		in, out := &in.Kernels, &out.Kernels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DefaultKernel != nil {
		in, out := &in.DefaultKernel, &out.DefaultKernel
		*out = new(string)
		**out = **in
	}
	if in.CullIdleTimeout != nil {
		in, out := &in.CullIdleTimeout, &out.CullIdleTimeout
		*out = new(int32)
		**out = **in
	}
	if in.CullInterval != nil {
		in, out := &in.CullInterval, &out.CullInterval
		*out = new(int32)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterRole != nil {
		in, out := &in.ClusterRole, &out.ClusterRole
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterGatewaySpec.
func (in *JupyterGatewaySpec) DeepCopy() *JupyterGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(JupyterGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterGatewayStatus) DeepCopyInto(out *JupyterGatewayStatus) {
	*out = *in
	in.DeploymentStatus.DeepCopyInto(&out.DeploymentStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterGatewayStatus.
func (in *JupyterGatewayStatus) DeepCopy() *JupyterGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(JupyterGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernel) DeepCopyInto(out *JupyterKernel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernel.
func (in *JupyterKernel) DeepCopy() *JupyterKernel {
	if in == nil {
		return nil
	}
	out := new(JupyterKernel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JupyterKernel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernelCRDSpec) DeepCopyInto(out *JupyterKernelCRDSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernelCRDSpec.
func (in *JupyterKernelCRDSpec) DeepCopy() *JupyterKernelCRDSpec {
	if in == nil {
		return nil
	}
	out := new(JupyterKernelCRDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernelCondition) DeepCopyInto(out *JupyterKernelCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernelCondition.
func (in *JupyterKernelCondition) DeepCopy() *JupyterKernelCondition {
	if in == nil {
		return nil
	}
	out := new(JupyterKernelCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernelList) DeepCopyInto(out *JupyterKernelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JupyterKernel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernelList.
func (in *JupyterKernelList) DeepCopy() *JupyterKernelList {
	if in == nil {
		return nil
	}
	out := new(JupyterKernelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JupyterKernelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernelSpec) DeepCopyInto(out *JupyterKernelSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernelSpec.
func (in *JupyterKernelSpec) DeepCopy() *JupyterKernelSpec {
	if in == nil {
		return nil
	}
	out := new(JupyterKernelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JupyterKernelSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernelSpecList) DeepCopyInto(out *JupyterKernelSpecList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JupyterKernelSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernelSpecList.
func (in *JupyterKernelSpecList) DeepCopy() *JupyterKernelSpecList {
	if in == nil {
		return nil
	}
	out := new(JupyterKernelSpecList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JupyterKernelSpecList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernelSpecSpec) DeepCopyInto(out *JupyterKernelSpecSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernelSpecSpec.
func (in *JupyterKernelSpecSpec) DeepCopy() *JupyterKernelSpecSpec {
	if in == nil {
		return nil
	}
	out := new(JupyterKernelSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernelSpecStatus) DeepCopyInto(out *JupyterKernelSpecStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernelSpecStatus.
func (in *JupyterKernelSpecStatus) DeepCopy() *JupyterKernelSpecStatus {
	if in == nil {
		return nil
	}
	out := new(JupyterKernelSpecStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernelStatus) DeepCopyInto(out *JupyterKernelStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]JupyterKernelCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.LastReconcileTime != nil {
		in, out := &in.LastReconcileTime, &out.LastReconcileTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernelStatus.
func (in *JupyterKernelStatus) DeepCopy() *JupyterKernelStatus {
	if in == nil {
		return nil
	}
	out := new(JupyterKernelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernelTemplate) DeepCopyInto(out *JupyterKernelTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernelTemplate.
func (in *JupyterKernelTemplate) DeepCopy() *JupyterKernelTemplate {
	if in == nil {
		return nil
	}
	out := new(JupyterKernelTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JupyterKernelTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernelTemplateList) DeepCopyInto(out *JupyterKernelTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JupyterKernelTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernelTemplateList.
func (in *JupyterKernelTemplateList) DeepCopy() *JupyterKernelTemplateList {
	if in == nil {
		return nil
	}
	out := new(JupyterKernelTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JupyterKernelTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernelTemplateSpec) DeepCopyInto(out *JupyterKernelTemplateSpec) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(v1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernelTemplateSpec.
func (in *JupyterKernelTemplateSpec) DeepCopy() *JupyterKernelTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(JupyterKernelTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterKernelTemplateStatus) DeepCopyInto(out *JupyterKernelTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterKernelTemplateStatus.
func (in *JupyterKernelTemplateStatus) DeepCopy() *JupyterKernelTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(JupyterKernelTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterNotebook) DeepCopyInto(out *JupyterNotebook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterNotebook.
func (in *JupyterNotebook) DeepCopy() *JupyterNotebook {
	if in == nil {
		return nil
	}
	out := new(JupyterNotebook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JupyterNotebook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterNotebookList) DeepCopyInto(out *JupyterNotebookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JupyterNotebook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterNotebookList.
func (in *JupyterNotebookList) DeepCopy() *JupyterNotebookList {
	if in == nil {
		return nil
	}
	out := new(JupyterNotebookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JupyterNotebookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterNotebookSpec) DeepCopyInto(out *JupyterNotebookSpec) {
	*out = *in
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(v1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterNotebookSpec.
func (in *JupyterNotebookSpec) DeepCopy() *JupyterNotebookSpec {
	if in == nil {
		return nil
	}
	out := new(JupyterNotebookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JupyterNotebookStatus) DeepCopyInto(out *JupyterNotebookStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JupyterNotebookStatus.
func (in *JupyterNotebookStatus) DeepCopy() *JupyterNotebookStatus {
	if in == nil {
		return nil
	}
	out := new(JupyterNotebookStatus)
	in.DeepCopyInto(out)
	return out
}
