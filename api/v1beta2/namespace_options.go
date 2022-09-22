package v1beta2

type NamespaceOptions struct {
	//+kubebuilder:validation:Minimum=1
	// Specifies the maximum number of namespaces allowed for that Tenant. Once the namespace quota assigned to the Tenant has been reached, the Tenant owner cannot create further namespaces. Optional.
	Quota *int32 `json:"quota,omitempty"`
	// Specifies additional labels and annotations the Capsule operator places on any Namespace resource in the Tenant. Optional.
	AdditionalMetadata *AdditionalMetadataSpec `json:"additionalMetadata,omitempty"`
	// Define the labels that a Tenant Owner cannot set for their Namespace resources.
	ForbiddenLabels ForbiddenListSpec `json:"forbiddenLabels"`
	// Define the annotations that a Tenant Owner cannot set for their Namespace resources.
	ForbiddenAnnotations ForbiddenListSpec `json:"forbiddenAnnotations"`
}
