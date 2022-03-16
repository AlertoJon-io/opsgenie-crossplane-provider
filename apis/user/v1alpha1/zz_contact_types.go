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

type ContactObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ContactParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Method *string `json:"method" tf:"method,omitempty"`

	// +kubebuilder:validation:Required
	To *string `json:"to" tf:"to,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

// ContactSpec defines the desired state of Contact
type ContactSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContactParameters `json:"forProvider"`
}

// ContactStatus defines the observed state of Contact.
type ContactStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContactObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Contact is the Schema for the Contacts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opsgenie-providerjet}
type Contact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContactSpec   `json:"spec"`
	Status            ContactStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContactList contains a list of Contacts
type ContactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Contact `json:"items"`
}

// Repository type metadata.
var (
	Contact_Kind             = "Contact"
	Contact_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Contact_Kind}.String()
	Contact_KindAPIVersion   = Contact_Kind + "." + CRDGroupVersion.String()
	Contact_GroupVersionKind = CRDGroupVersion.WithKind(Contact_Kind)
)

func init() {
	SchemeBuilder.Register(&Contact{}, &ContactList{})
}
