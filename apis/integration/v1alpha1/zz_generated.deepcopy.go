//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcknowledgeObservation) DeepCopyInto(out *AcknowledgeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcknowledgeObservation.
func (in *AcknowledgeObservation) DeepCopy() *AcknowledgeObservation {
	if in == nil {
		return nil
	}
	out := new(AcknowledgeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcknowledgeParameters) DeepCopyInto(out *AcknowledgeParameters) {
	*out = *in
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]FilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Note != nil {
		in, out := &in.Note, &out.Note
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcknowledgeParameters.
func (in *AcknowledgeParameters) DeepCopy() *AcknowledgeParameters {
	if in == nil {
		return nil
	}
	out := new(AcknowledgeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Action) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionList) DeepCopyInto(out *ActionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Action, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionList.
func (in *ActionList) DeepCopy() *ActionList {
	if in == nil {
		return nil
	}
	out := new(ActionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionObservation) DeepCopyInto(out *ActionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionObservation.
func (in *ActionObservation) DeepCopy() *ActionObservation {
	if in == nil {
		return nil
	}
	out := new(ActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionParameters) DeepCopyInto(out *ActionParameters) {
	*out = *in
	if in.Acknowledge != nil {
		in, out := &in.Acknowledge, &out.Acknowledge
		*out = make([]AcknowledgeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AddNote != nil {
		in, out := &in.AddNote, &out.AddNote
		*out = make([]AddNoteParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Close != nil {
		in, out := &in.Close, &out.Close
		*out = make([]CloseParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = make([]CreateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ignore != nil {
		in, out := &in.Ignore, &out.Ignore
		*out = make([]IgnoreParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IntegrationID != nil {
		in, out := &in.IntegrationID, &out.IntegrationID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionParameters.
func (in *ActionParameters) DeepCopy() *ActionParameters {
	if in == nil {
		return nil
	}
	out := new(ActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionSpec) DeepCopyInto(out *ActionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionSpec.
func (in *ActionSpec) DeepCopy() *ActionSpec {
	if in == nil {
		return nil
	}
	out := new(ActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionStatus) DeepCopyInto(out *ActionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionStatus.
func (in *ActionStatus) DeepCopy() *ActionStatus {
	if in == nil {
		return nil
	}
	out := new(ActionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddNoteFilterObservation) DeepCopyInto(out *AddNoteFilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddNoteFilterObservation.
func (in *AddNoteFilterObservation) DeepCopy() *AddNoteFilterObservation {
	if in == nil {
		return nil
	}
	out := new(AddNoteFilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddNoteFilterParameters) DeepCopyInto(out *AddNoteFilterParameters) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]FilterConditionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddNoteFilterParameters.
func (in *AddNoteFilterParameters) DeepCopy() *AddNoteFilterParameters {
	if in == nil {
		return nil
	}
	out := new(AddNoteFilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddNoteObservation) DeepCopyInto(out *AddNoteObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddNoteObservation.
func (in *AddNoteObservation) DeepCopy() *AddNoteObservation {
	if in == nil {
		return nil
	}
	out := new(AddNoteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddNoteParameters) DeepCopyInto(out *AddNoteParameters) {
	*out = *in
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]AddNoteFilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Note != nil {
		in, out := &in.Note, &out.Note
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddNoteParameters.
func (in *AddNoteParameters) DeepCopy() *AddNoteParameters {
	if in == nil {
		return nil
	}
	out := new(AddNoteParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloseFilterConditionsObservation) DeepCopyInto(out *CloseFilterConditionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloseFilterConditionsObservation.
func (in *CloseFilterConditionsObservation) DeepCopy() *CloseFilterConditionsObservation {
	if in == nil {
		return nil
	}
	out := new(CloseFilterConditionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloseFilterConditionsParameters) DeepCopyInto(out *CloseFilterConditionsParameters) {
	*out = *in
	if in.ExpectedValue != nil {
		in, out := &in.ExpectedValue, &out.ExpectedValue
		*out = new(string)
		**out = **in
	}
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(bool)
		**out = **in
	}
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloseFilterConditionsParameters.
func (in *CloseFilterConditionsParameters) DeepCopy() *CloseFilterConditionsParameters {
	if in == nil {
		return nil
	}
	out := new(CloseFilterConditionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloseFilterObservation) DeepCopyInto(out *CloseFilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloseFilterObservation.
func (in *CloseFilterObservation) DeepCopy() *CloseFilterObservation {
	if in == nil {
		return nil
	}
	out := new(CloseFilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloseFilterParameters) DeepCopyInto(out *CloseFilterParameters) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CloseFilterConditionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloseFilterParameters.
func (in *CloseFilterParameters) DeepCopy() *CloseFilterParameters {
	if in == nil {
		return nil
	}
	out := new(CloseFilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloseObservation) DeepCopyInto(out *CloseObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloseObservation.
func (in *CloseObservation) DeepCopy() *CloseObservation {
	if in == nil {
		return nil
	}
	out := new(CloseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloseParameters) DeepCopyInto(out *CloseParameters) {
	*out = *in
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]CloseFilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Note != nil {
		in, out := &in.Note, &out.Note
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloseParameters.
func (in *CloseParameters) DeepCopy() *CloseParameters {
	if in == nil {
		return nil
	}
	out := new(CloseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionsObservation) DeepCopyInto(out *ConditionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionsObservation.
func (in *ConditionsObservation) DeepCopy() *ConditionsObservation {
	if in == nil {
		return nil
	}
	out := new(ConditionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionsParameters) DeepCopyInto(out *ConditionsParameters) {
	*out = *in
	if in.ExpectedValue != nil {
		in, out := &in.ExpectedValue, &out.ExpectedValue
		*out = new(string)
		**out = **in
	}
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(bool)
		**out = **in
	}
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionsParameters.
func (in *ConditionsParameters) DeepCopy() *ConditionsParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateFilterConditionsObservation) DeepCopyInto(out *CreateFilterConditionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateFilterConditionsObservation.
func (in *CreateFilterConditionsObservation) DeepCopy() *CreateFilterConditionsObservation {
	if in == nil {
		return nil
	}
	out := new(CreateFilterConditionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateFilterConditionsParameters) DeepCopyInto(out *CreateFilterConditionsParameters) {
	*out = *in
	if in.ExpectedValue != nil {
		in, out := &in.ExpectedValue, &out.ExpectedValue
		*out = new(string)
		**out = **in
	}
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(bool)
		**out = **in
	}
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateFilterConditionsParameters.
func (in *CreateFilterConditionsParameters) DeepCopy() *CreateFilterConditionsParameters {
	if in == nil {
		return nil
	}
	out := new(CreateFilterConditionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateFilterObservation) DeepCopyInto(out *CreateFilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateFilterObservation.
func (in *CreateFilterObservation) DeepCopy() *CreateFilterObservation {
	if in == nil {
		return nil
	}
	out := new(CreateFilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateFilterParameters) DeepCopyInto(out *CreateFilterParameters) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CreateFilterConditionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateFilterParameters.
func (in *CreateFilterParameters) DeepCopy() *CreateFilterParameters {
	if in == nil {
		return nil
	}
	out := new(CreateFilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateObservation) DeepCopyInto(out *CreateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateObservation.
func (in *CreateObservation) DeepCopy() *CreateObservation {
	if in == nil {
		return nil
	}
	out := new(CreateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateParameters) DeepCopyInto(out *CreateParameters) {
	*out = *in
	if in.AlertActions != nil {
		in, out := &in.AlertActions, &out.AlertActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.AppendAttachments != nil {
		in, out := &in.AppendAttachments, &out.AppendAttachments
		*out = new(bool)
		**out = **in
	}
	if in.CustomPriority != nil {
		in, out := &in.CustomPriority, &out.CustomPriority
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Entity != nil {
		in, out := &in.Entity, &out.Entity
		*out = new(string)
		**out = **in
	}
	if in.ExtraProperties != nil {
		in, out := &in.ExtraProperties, &out.ExtraProperties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]CreateFilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IgnoreAlertActionsFromPayload != nil {
		in, out := &in.IgnoreAlertActionsFromPayload, &out.IgnoreAlertActionsFromPayload
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreExtraPropertiesFromPayload != nil {
		in, out := &in.IgnoreExtraPropertiesFromPayload, &out.IgnoreExtraPropertiesFromPayload
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreRespondersFromPayload != nil {
		in, out := &in.IgnoreRespondersFromPayload, &out.IgnoreRespondersFromPayload
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreTagsFromPayload != nil {
		in, out := &in.IgnoreTagsFromPayload, &out.IgnoreTagsFromPayload
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreTeamsFromPayload != nil {
		in, out := &in.IgnoreTeamsFromPayload, &out.IgnoreTeamsFromPayload
		*out = new(bool)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Note != nil {
		in, out := &in.Note, &out.Note
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(string)
		**out = **in
	}
	if in.Responders != nil {
		in, out := &in.Responders, &out.Responders
		*out = make([]RespondersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateParameters.
func (in *CreateParameters) DeepCopy() *CreateParameters {
	if in == nil {
		return nil
	}
	out := new(CreateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterConditionsObservation) DeepCopyInto(out *FilterConditionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterConditionsObservation.
func (in *FilterConditionsObservation) DeepCopy() *FilterConditionsObservation {
	if in == nil {
		return nil
	}
	out := new(FilterConditionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterConditionsParameters) DeepCopyInto(out *FilterConditionsParameters) {
	*out = *in
	if in.ExpectedValue != nil {
		in, out := &in.ExpectedValue, &out.ExpectedValue
		*out = new(string)
		**out = **in
	}
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(bool)
		**out = **in
	}
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterConditionsParameters.
func (in *FilterConditionsParameters) DeepCopy() *FilterConditionsParameters {
	if in == nil {
		return nil
	}
	out := new(FilterConditionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterObservation) DeepCopyInto(out *FilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterObservation.
func (in *FilterObservation) DeepCopy() *FilterObservation {
	if in == nil {
		return nil
	}
	out := new(FilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterParameters) DeepCopyInto(out *FilterParameters) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ConditionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterParameters.
func (in *FilterParameters) DeepCopy() *FilterParameters {
	if in == nil {
		return nil
	}
	out := new(FilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IgnoreFilterConditionsObservation) DeepCopyInto(out *IgnoreFilterConditionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IgnoreFilterConditionsObservation.
func (in *IgnoreFilterConditionsObservation) DeepCopy() *IgnoreFilterConditionsObservation {
	if in == nil {
		return nil
	}
	out := new(IgnoreFilterConditionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IgnoreFilterConditionsParameters) DeepCopyInto(out *IgnoreFilterConditionsParameters) {
	*out = *in
	if in.ExpectedValue != nil {
		in, out := &in.ExpectedValue, &out.ExpectedValue
		*out = new(string)
		**out = **in
	}
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(bool)
		**out = **in
	}
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IgnoreFilterConditionsParameters.
func (in *IgnoreFilterConditionsParameters) DeepCopy() *IgnoreFilterConditionsParameters {
	if in == nil {
		return nil
	}
	out := new(IgnoreFilterConditionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IgnoreFilterObservation) DeepCopyInto(out *IgnoreFilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IgnoreFilterObservation.
func (in *IgnoreFilterObservation) DeepCopy() *IgnoreFilterObservation {
	if in == nil {
		return nil
	}
	out := new(IgnoreFilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IgnoreFilterParameters) DeepCopyInto(out *IgnoreFilterParameters) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]IgnoreFilterConditionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IgnoreFilterParameters.
func (in *IgnoreFilterParameters) DeepCopy() *IgnoreFilterParameters {
	if in == nil {
		return nil
	}
	out := new(IgnoreFilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IgnoreObservation) DeepCopyInto(out *IgnoreObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IgnoreObservation.
func (in *IgnoreObservation) DeepCopy() *IgnoreObservation {
	if in == nil {
		return nil
	}
	out := new(IgnoreObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IgnoreParameters) DeepCopyInto(out *IgnoreParameters) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]IgnoreFilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IgnoreParameters.
func (in *IgnoreParameters) DeepCopy() *IgnoreParameters {
	if in == nil {
		return nil
	}
	out := new(IgnoreParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RespondersObservation) DeepCopyInto(out *RespondersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RespondersObservation.
func (in *RespondersObservation) DeepCopy() *RespondersObservation {
	if in == nil {
		return nil
	}
	out := new(RespondersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RespondersParameters) DeepCopyInto(out *RespondersParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RespondersParameters.
func (in *RespondersParameters) DeepCopy() *RespondersParameters {
	if in == nil {
		return nil
	}
	out := new(RespondersParameters)
	in.DeepCopyInto(out)
	return out
}
