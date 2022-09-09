//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedResourceInfo) DeepCopyInto(out *AppliedResourceInfo) {
	*out = *in
	out.Ref = in.Ref
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedResourceInfo.
func (in *AppliedResourceInfo) DeepCopy() *AppliedResourceInfo {
	if in == nil {
		return nil
	}
	out := new(AppliedResourceInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Generator) DeepCopyInto(out *Generator) {
	*out = *in
	if in.PullRequest != nil {
		in, out := &in.PullRequest, &out.PullRequest
		*out = new(PullRequestGenerator)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Generator.
func (in *Generator) DeepCopy() *Generator {
	if in == nil {
		return nil
	}
	out := new(Generator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestGenerator) DeepCopyInto(out *PullRequestGenerator) {
	*out = *in
	if in.Gitlab != nil {
		in, out := &in.Gitlab, &out.Gitlab
		*out = new(PullRequestGeneratorGitlab)
		(*in).DeepCopyInto(*out)
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]PullRequestGeneratorFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestGenerator.
func (in *PullRequestGenerator) DeepCopy() *PullRequestGenerator {
	if in == nil {
		return nil
	}
	out := new(PullRequestGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestGeneratorFilter) DeepCopyInto(out *PullRequestGeneratorFilter) {
	*out = *in
	if in.BranchMatch != nil {
		in, out := &in.BranchMatch, &out.BranchMatch
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestGeneratorFilter.
func (in *PullRequestGeneratorFilter) DeepCopy() *PullRequestGeneratorFilter {
	if in == nil {
		return nil
	}
	out := new(PullRequestGeneratorFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestGeneratorGitlab) DeepCopyInto(out *PullRequestGeneratorGitlab) {
	*out = *in
	if in.TokenRef != nil {
		in, out := &in.TokenRef, &out.TokenRef
		*out = new(SecretRef)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestGeneratorGitlab.
func (in *PullRequestGeneratorGitlab) DeepCopy() *PullRequestGeneratorGitlab {
	if in == nil {
		return nil
	}
	out := new(PullRequestGeneratorGitlab)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRef) DeepCopyInto(out *ResourceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRef.
func (in *ResourceRef) DeepCopy() *ResourceRef {
	if in == nil {
		return nil
	}
	out := new(ResourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTemplate) DeepCopyInto(out *ResourceTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTemplate.
func (in *ResourceTemplate) DeepCopy() *ResourceTemplate {
	if in == nil {
		return nil
	}
	out := new(ResourceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTemplateList) DeepCopyInto(out *ResourceTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTemplateList.
func (in *ResourceTemplateList) DeepCopy() *ResourceTemplateList {
	if in == nil {
		return nil
	}
	out := new(ResourceTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTemplateSpec) DeepCopyInto(out *ResourceTemplateSpec) {
	*out = *in
	out.Interval = in.Interval
	if in.Generators != nil {
		in, out := &in.Generators, &out.Generators
		*out = make([]Generator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]Template, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTemplateSpec.
func (in *ResourceTemplateSpec) DeepCopy() *ResourceTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTemplateStatus) DeepCopyInto(out *ResourceTemplateStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AppliedResources != nil {
		in, out := &in.AppliedResources, &out.AppliedResources
		*out = make([]AppliedResourceInfo, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTemplateStatus.
func (in *ResourceTemplateStatus) DeepCopy() *ResourceTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusReporter) DeepCopyInto(out *StatusReporter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusReporter.
func (in *StatusReporter) DeepCopy() *StatusReporter {
	if in == nil {
		return nil
	}
	out := new(StatusReporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatusReporter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusReporterList) DeepCopyInto(out *StatusReporterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StatusReporter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusReporterList.
func (in *StatusReporterList) DeepCopy() *StatusReporterList {
	if in == nil {
		return nil
	}
	out := new(StatusReporterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatusReporterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusReporterSpec) DeepCopyInto(out *StatusReporterSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusReporterSpec.
func (in *StatusReporterSpec) DeepCopy() *StatusReporterSpec {
	if in == nil {
		return nil
	}
	out := new(StatusReporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusReporterStatus) DeepCopyInto(out *StatusReporterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusReporterStatus.
func (in *StatusReporterStatus) DeepCopy() *StatusReporterStatus {
	if in == nil {
		return nil
	}
	out := new(StatusReporterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Template) DeepCopyInto(out *Template) {
	*out = *in
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = (*in).DeepCopy()
	}
	if in.Raw != nil {
		in, out := &in.Raw, &out.Raw
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Template.
func (in *Template) DeepCopy() *Template {
	if in == nil {
		return nil
	}
	out := new(Template)
	in.DeepCopyInto(out)
	return out
}
