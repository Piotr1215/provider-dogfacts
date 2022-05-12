/*
Copyright 2020 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// DogFactParameters are the configurable fields of a DogFact.
type DogFactParameters struct {
	Fact string `json:"fact"`
}

// DogFactObservation are the observable fields of a DogFact.
type DogFactObservation struct {
	DogFactsNumber string `json:"dogFactsNumber,omitempty"`
}

// A DogFactSpec defines the desired state of a DogFact.
type DogFactSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DogFactParameters `json:"forProvider"`
}

// A DogFactStatus represents the observed state of a DogFact.
type DogFactStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DogFactObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A DogFact is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
type DogFact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DogFactSpec   `json:"spec"`
	Status DogFactStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DogFactList contains a list of DogFact
type DogFactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DogFact `json:"items"`
}

// DogFact type metadata.
var (
	DogFactKind             = reflect.TypeOf(DogFact{}).Name()
	DogFactGroupKind        = schema.GroupKind{Group: Group, Kind: DogFactKind}.String()
	DogFactKindAPIVersion   = DogFactKind + "." + SchemeGroupVersion.String()
	DogFactGroupVersionKind = SchemeGroupVersion.WithKind(DogFactKind)
)

func init() {
	SchemeBuilder.Register(&DogFact{}, &DogFactList{})
}
