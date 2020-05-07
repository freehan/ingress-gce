/*
Copyright 2018 The Kubernetes Authors.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
// +k8s:openapi-gen=true
type NetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec              NetworkEndpointGroupSpec   `json:"spec,omitempty"`
	Status            NetworkEndpointGroupStatus `json:"status,omitempty"`
}

// NetworkEndpointGroupSpec is the spec for a NetworkEndpointGroup resource
// +k8s:openapi-gen=true
type NetworkEndpointGroupSpec struct {}

// NetworkEndpointGroupStatus is the status for a NetworkEndpointGroup resource
type NetworkEndpointGroupStatus struct {
	NetworkEndpointGroups []NegObjectReference `json:"networkEndpointGroups,omitempty"`
	Condition []NegCondition `json:"conditions,omitempty"`
}

// NegObjectReference is the object reference to the NEG resource in GCE
type NegObjectReference struct {
	// The unique identifier for the NEG resource in GCE API.
	Id uint64 `json:"id,omitempty,string"`

	// SelfLink is the GCE Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// NetworkEndpointType: Type of network endpoints in this network
	// endpoint group.
	NetworkEndpointType NetworkEndpointType `json:"networkEndpointType,omitempty"`

	// Zone is where the network endpoint group is located.
	Zone string `json:"zone,omitempty"`
}

type NetworkEndpointType string

const (
	VmIpPortEndpointType      = NetworkEndpointType("GCE_VM_IP_PORT")
	VmIpEndpointType          = NetworkEndpointType("GCE_VM_IP")
	NonGCPPrivateEndpointType = NetworkEndpointType("NON_GCP_PRIVATE_IP_PORT")
)

// NegCondition contains details for the current condition of this NEG.
type NegCondition struct {
	// Type is the type of the condition.
	Type NegConditionType `json:"type"`
	// Status is the status of the condition.
	// Can be True, False, Unknown.
	Status NegConditionStatus `json:"status"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

type NegConditionType string

// These are valid conditions of NEG.
const (
	Initialized NegConditionType = "Initialized"
)

type NegConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition.
// "ConditionFalse" means a resource is not in the condition. "ConditionUnknown" means NEG controller
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
const (
	ConditionTrue    NegConditionStatus = "True"
	ConditionFalse   NegConditionStatus = "False"
	ConditionUnknown NegConditionStatus = "Unknown"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkEndpointGroupList is a list of NetworkEndpointGroup resources
type NetworkEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NetworkEndpointGroup `json:"items"`
}
