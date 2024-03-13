// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1alpha1 "github.com/ironcore-dev/ipam/api/ipam/v1alpha1"
	ipamv1alpha1 "github.com/ironcore-dev/ipam/clientgo/applyconfiguration/ipam/v1alpha1"
	metav1 "github.com/ironcore-dev/ipam/clientgo/applyconfiguration/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=ipam.metal.ironcore.dev, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("IP"):
		return &ipamv1alpha1.IPApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("IPSpec"):
		return &ipamv1alpha1.IPSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("IPStatus"):
		return &ipamv1alpha1.IPStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Network"):
		return &ipamv1alpha1.NetworkApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NetworkCounter"):
		return &ipamv1alpha1.NetworkCounterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NetworkCounterSpec"):
		return &ipamv1alpha1.NetworkCounterSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NetworkIDInterval"):
		return &ipamv1alpha1.NetworkIDIntervalApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NetworkSpec"):
		return &ipamv1alpha1.NetworkSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NetworkStatus"):
		return &ipamv1alpha1.NetworkStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Region"):
		return &ipamv1alpha1.RegionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ResourceReference"):
		return &ipamv1alpha1.ResourceReferenceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Subnet"):
		return &ipamv1alpha1.SubnetApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubnetSpec"):
		return &ipamv1alpha1.SubnetSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubnetStatus"):
		return &ipamv1alpha1.SubnetStatusApplyConfiguration{}

		// Group=meta.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("ManagedFieldsEntry"):
		return &metav1.ManagedFieldsEntryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ObjectMeta"):
		return &metav1.ObjectMetaApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OwnerReference"):
		return &metav1.OwnerReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TypeMeta"):
		return &metav1.TypeMetaApplyConfiguration{}

	}
	return nil
}
