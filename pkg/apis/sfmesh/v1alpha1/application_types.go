/*
Copyright 2019 The SFmesh-Controller Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Resorcegroup string            `json:"resourceGroup"`
	Location     string            `json:"location"`
	Tags         map[string]string `json:"tags,omitempty"`
	Description  string            `json:"description,omitempty"`
	Services     []Service         `json:"services,omitempty"`
}

type Service struct {
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	OsType       string        `json:"osType"`
	ReplicaCount int32         `json:"replicaCount"`
	CodePackages []CodePackage `json:"codePackages,omitempty"`
	NetworkRefs  []string      `json:"networks,omitempty"`
}

type CodePackage struct {
	Name      string     `json:"name"`
	Image     string     `json:"image"`
	EndPoints []EndPoint `json:"endPoints,omitempty"`

	//TODO: Add More Service Options like commands
}

type EndPoint struct {
	Name string `json:"name"`
	Port int32  `json:"port"`
}

// ApplicationStatus defines the observed state of Application
type ApplicationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ProvisioningState   string   `json:"provisioningState"`
	HealthState         string   `json:"healthState"`
	UnHealthyEvaulation string   `json:"unhealthyEvaluation"`
	ServiceNames        []string `json:"serviceNames,omitempty"`
	Status              string   `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Application is the Schema for the applications API
// +k8s:openapi-gen=true
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApplicationList contains a list of Application
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
