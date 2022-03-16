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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AutoCloseActionObservation struct {
}

type AutoCloseActionParameters struct {

	// +kubebuilder:validation:Required
	Duration []DurationParameters `json:"duration" tf:"duration,omitempty"`
}

type AutoRestartActionDurationObservation struct {
}

type AutoRestartActionDurationParameters struct {

	// +kubebuilder:validation:Required
	TimeAmount *float64 `json:"timeAmount" tf:"time_amount,omitempty"`

	// +kubebuilder:validation:Optional
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type AutoRestartActionObservation struct {
}

type AutoRestartActionParameters struct {

	// +kubebuilder:validation:Required
	Duration []AutoRestartActionDurationParameters `json:"duration" tf:"duration,omitempty"`

	// +kubebuilder:validation:Required
	MaxRepeatCount *float64 `json:"maxRepeatCount" tf:"max_repeat_count,omitempty"`
}

type ConditionsObservation struct {
}

type ConditionsParameters struct {

	// User defined value that will be compared with alert field according to the operation. Default value is empty string
	// +kubebuilder:validation:Optional
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// If 'field' is set as 'extra-properties', key could be used for key-value pair
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Indicates behaviour of the given operation. Default value is false
	// +kubebuilder:validation:Optional
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// +kubebuilder:validation:Required
	Operation *string `json:"operation" tf:"operation,omitempty"`

	// Order of the condition in conditions list
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type DeDuplicationActionDurationObservation struct {
}

type DeDuplicationActionDurationParameters struct {

	// +kubebuilder:validation:Required
	TimeAmount *float64 `json:"timeAmount" tf:"time_amount,omitempty"`

	// +kubebuilder:validation:Optional
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type DeDuplicationActionObservation struct {
}

type DeDuplicationActionParameters struct {

	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`

	// +kubebuilder:validation:Required
	DeDuplicationActionType *string `json:"deDuplicationActionType" tf:"de_duplication_action_type,omitempty"`

	// +kubebuilder:validation:Optional
	Duration []DeDuplicationActionDurationParameters `json:"duration,omitempty" tf:"duration,omitempty"`
}

type DelayActionDurationObservation struct {
}

type DelayActionDurationParameters struct {

	// +kubebuilder:validation:Required
	TimeAmount *float64 `json:"timeAmount" tf:"time_amount,omitempty"`

	// +kubebuilder:validation:Optional
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type DelayActionObservation struct {
}

type DelayActionParameters struct {

	// +kubebuilder:validation:Required
	DelayOption *string `json:"delayOption" tf:"delay_option,omitempty"`

	// +kubebuilder:validation:Optional
	Duration []DelayActionDurationParameters `json:"duration,omitempty" tf:"duration,omitempty"`

	// +kubebuilder:validation:Optional
	UntilHour *float64 `json:"untilHour,omitempty" tf:"until_hour,omitempty"`

	// +kubebuilder:validation:Optional
	UntilMinute *float64 `json:"untilMinute,omitempty" tf:"until_minute,omitempty"`
}

type DurationObservation struct {
}

type DurationParameters struct {

	// +kubebuilder:validation:Required
	TimeAmount *float64 `json:"timeAmount" tf:"time_amount,omitempty"`

	// +kubebuilder:validation:Optional
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// +kubebuilder:validation:Optional
	Conditions []ConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PolicyParameters struct {

	// +kubebuilder:validation:Optional
	AutoCloseAction []AutoCloseActionParameters `json:"autoCloseAction,omitempty" tf:"auto_close_action,omitempty"`

	// +kubebuilder:validation:Optional
	AutoRestartAction []AutoRestartActionParameters `json:"autoRestartAction,omitempty" tf:"auto_restart_action,omitempty"`

	// +kubebuilder:validation:Optional
	DeDuplicationAction []DeDuplicationActionParameters `json:"deDuplicationAction,omitempty" tf:"de_duplication_action,omitempty"`

	// +kubebuilder:validation:Optional
	DelayAction []DelayActionParameters `json:"delayAction,omitempty" tf:"delay_action,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Filter []FilterParameters `json:"filter" tf:"filter,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyDescription *string `json:"policyDescription,omitempty" tf:"policy_description,omitempty"`

	// +kubebuilder:validation:Optional
	Suppress *bool `json:"suppress,omitempty" tf:"suppress,omitempty"`

	// +kubebuilder:validation:Required
	TeamID *string `json:"teamId" tf:"team_id,omitempty"`

	// +kubebuilder:validation:Optional
	TimeRestriction []TimeRestrictionParameters `json:"timeRestriction,omitempty" tf:"time_restriction,omitempty"`
}

type RestrictionObservation struct {
}

type RestrictionParameters struct {

	// +kubebuilder:validation:Required
	EndHour *float64 `json:"endHour" tf:"end_hour,omitempty"`

	// +kubebuilder:validation:Required
	EndMin *float64 `json:"endMin" tf:"end_min,omitempty"`

	// +kubebuilder:validation:Required
	StartHour *float64 `json:"startHour" tf:"start_hour,omitempty"`

	// +kubebuilder:validation:Required
	StartMin *float64 `json:"startMin" tf:"start_min,omitempty"`
}

type RestrictionsObservation struct {
}

type RestrictionsParameters struct {

	// +kubebuilder:validation:Required
	EndDay *string `json:"endDay" tf:"end_day,omitempty"`

	// +kubebuilder:validation:Required
	EndHour *float64 `json:"endHour" tf:"end_hour,omitempty"`

	// +kubebuilder:validation:Required
	EndMin *float64 `json:"endMin" tf:"end_min,omitempty"`

	// +kubebuilder:validation:Required
	StartDay *string `json:"startDay" tf:"start_day,omitempty"`

	// +kubebuilder:validation:Required
	StartHour *float64 `json:"startHour" tf:"start_hour,omitempty"`

	// +kubebuilder:validation:Required
	StartMin *float64 `json:"startMin" tf:"start_min,omitempty"`
}

type TimeRestrictionObservation struct {
}

type TimeRestrictionParameters struct {

	// +kubebuilder:validation:Optional
	Restriction []RestrictionParameters `json:"restriction,omitempty" tf:"restriction,omitempty"`

	// +kubebuilder:validation:Optional
	Restrictions []RestrictionsParameters `json:"restrictions,omitempty" tf:"restrictions,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opsgenie-providerjet}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec"`
	Status            PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
