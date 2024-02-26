// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	ipamv1alpha1 "github.com/ironcore-dev/ipam/api/ipam/v1alpha1"
	internal "github.com/ironcore-dev/ipam/clientgo/applyconfiguration/internal"
	v1 "github.com/ironcore-dev/ipam/clientgo/applyconfiguration/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
)

// NetworkCounterApplyConfiguration represents an declarative configuration of the NetworkCounter type for use
// with apply.
type NetworkCounterApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *NetworkCounterSpecApplyConfiguration `json:"spec,omitempty"`
	Status                           *ipamv1alpha1.NetworkCounterStatus    `json:"status,omitempty"`
}

// NetworkCounter constructs an declarative configuration of the NetworkCounter type for use with
// apply.
func NetworkCounter(name, namespace string) *NetworkCounterApplyConfiguration {
	b := &NetworkCounterApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("NetworkCounter")
	b.WithAPIVersion("ipam.onmetal.de/v1alpha1")
	return b
}

// ExtractNetworkCounter extracts the applied configuration owned by fieldManager from
// networkCounter. If no managedFields are found in networkCounter for fieldManager, a
// NetworkCounterApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// networkCounter must be a unmodified NetworkCounter API object that was retrieved from the Kubernetes API.
// ExtractNetworkCounter provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractNetworkCounter(networkCounter *ipamv1alpha1.NetworkCounter, fieldManager string) (*NetworkCounterApplyConfiguration, error) {
	return extractNetworkCounter(networkCounter, fieldManager, "")
}

// ExtractNetworkCounterStatus is the same as ExtractNetworkCounter except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractNetworkCounterStatus(networkCounter *ipamv1alpha1.NetworkCounter, fieldManager string) (*NetworkCounterApplyConfiguration, error) {
	return extractNetworkCounter(networkCounter, fieldManager, "status")
}

func extractNetworkCounter(networkCounter *ipamv1alpha1.NetworkCounter, fieldManager string, subresource string) (*NetworkCounterApplyConfiguration, error) {
	b := &NetworkCounterApplyConfiguration{}
	err := managedfields.ExtractInto(networkCounter, internal.Parser().Type("com.github.onmetal.ipam.api.ipam.v1alpha1.NetworkCounter"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(networkCounter.Name)
	b.WithNamespace(networkCounter.Namespace)

	b.WithKind("NetworkCounter")
	b.WithAPIVersion("ipam.onmetal.de/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithKind(value string) *NetworkCounterApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithAPIVersion(value string) *NetworkCounterApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithName(value string) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithGenerateName(value string) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithNamespace(value string) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithUID(value types.UID) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithResourceVersion(value string) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithGeneration(value int64) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithCreationTimestamp(value metav1.Time) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *NetworkCounterApplyConfiguration) WithLabels(entries map[string]string) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *NetworkCounterApplyConfiguration) WithAnnotations(entries map[string]string) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *NetworkCounterApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *NetworkCounterApplyConfiguration) WithFinalizers(values ...string) *NetworkCounterApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *NetworkCounterApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithSpec(value *NetworkCounterSpecApplyConfiguration) *NetworkCounterApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *NetworkCounterApplyConfiguration) WithStatus(value ipamv1alpha1.NetworkCounterStatus) *NetworkCounterApplyConfiguration {
	b.Status = &value
	return b
}
