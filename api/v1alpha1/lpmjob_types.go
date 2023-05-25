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

type MigrationStatus string

const (
	Processing    MigrationStatus = "Processing"
	Checkpointing MigrationStatus = "Checkpointing"
	Transfering   MigrationStatus = "Transfering"
	Restoring     MigrationStatus = "Restoring"
	Cleaning      MigrationStatus = "Cleaning"
	Done          MigrationStatus = "Done"
	Error         MigrationStatus = "Error"
)

// LPMJobSpec defines the desired state of LPMJob
type LPMJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// LPMJobStatus defines the observed state of LPMJob
type LPMJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	SourcePodName     string `json:"sourcePodName,omitempty"`
	SourceNodeName    string `json:"sourceNodeName,omitempty"`
	SourceNodeAddress string `json:"sourceNodeAddress,omitempty"`

	DestionationPodName     string `json:"destionationPodName,omitempty"`
	DestionationNodeName    string `json:"destionationNodeName,omitempty"`
	DestionationNodeAddress string `json:"destionationNodeAddress,omitempty"`

	MigrationStatus MigrationStatus `json:"migrationStatus,omitempty"`

	NumberOfContainers int               `json:"numberOfContainers,omitempty"`
	ContainerStatuses  ContainerStatuses `json:"containerStatuses,omitempty"`
}

type ContainerStatuses struct {
	ContainerIdentifier      string          `json:"containerIdentifier,omitempty"`
	ContainerMigrationStatus MigrationStatus `json:"containerMigrationStatus,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// LPMJob is the Schema for the lpmjobs API
type LPMJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LPMJobSpec   `json:"spec,omitempty"`
	Status LPMJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LPMJobList contains a list of LPMJob
type LPMJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LPMJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LPMJob{}, &LPMJobList{})
}
