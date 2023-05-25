/*
Copyright 2023.

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

// MigrationPodSpec defines the desired state of MigrationPod
type MigrationPodSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of MigrationPod. Edit migrationpod_types.go to remove/update
	ContainerTemplates []ContainerTemplates `json:"containerTemplates"`
}

type ContainerTemplates struct {
	ContainerName  string `json:"containerName,omitempty"`
	ContainerImage string `json:"containerImage,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MigrationPod is the Schema for the migrationpods API
type MigrationPod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MigrationPodSpec `json:"spec"`
}

//+kubebuilder:object:root=true

// MigrationPodList contains a list of MigrationPod
type MigrationPodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MigrationPod `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MigrationPod{}, &MigrationPodList{})
}
