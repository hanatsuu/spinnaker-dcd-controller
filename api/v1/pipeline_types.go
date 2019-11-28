package v1

import (
	"encoding/json"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpinnakerPipelineResource defines the resource of Spinnaker
type SpinnakerPipelineResource struct {
	ApplicationName string `json:"applicationName,omitempty"`
	ID              string `json:"id,omitempty"`
}

// PipelineStatus defines the observed state of Pipeline
type PipelineStatus struct {
	SpinnakerResource SpinnakerPipelineResource `json:"spinnakerResource,omitempty"`
	Phase             string                    `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="PHASE",type=string,JSONPath=`.status.phase`

// Pipeline is the schema for Spinnaker Pipeline
type Pipeline struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	Spec   json.RawMessage `json:"spec,omitempty"`
	Status PipelineStatus  `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PipelineList contains a list of Pipeline
type PipelineList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata,omitempty"`
	Items           []Pipeline `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pipeline{}, &PipelineList{})
}
