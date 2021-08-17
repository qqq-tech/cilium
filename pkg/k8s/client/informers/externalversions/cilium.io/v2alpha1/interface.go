// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by informer-gen. DO NOT EDIT.

package v2alpha1

import (
	internalinterfaces "github.com/cilium/cilium/pkg/k8s/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CiliumEgressNATPolicies returns a CiliumEgressNATPolicyInformer.
	CiliumEgressNATPolicies() CiliumEgressNATPolicyInformer
	// CiliumEndpointSlices returns a CiliumEndpointSliceInformer.
	CiliumEndpointSlices() CiliumEndpointSliceInformer
	// CiliumEnvoyConfigs returns a CiliumEnvoyConfigInformer.
	CiliumEnvoyConfigs() CiliumEnvoyConfigInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CiliumEgressNATPolicies returns a CiliumEgressNATPolicyInformer.
func (v *version) CiliumEgressNATPolicies() CiliumEgressNATPolicyInformer {
	return &ciliumEgressNATPolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CiliumEndpointSlices returns a CiliumEndpointSliceInformer.
func (v *version) CiliumEndpointSlices() CiliumEndpointSliceInformer {
	return &ciliumEndpointSliceInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CiliumEnvoyConfigs returns a CiliumEnvoyConfigInformer.
func (v *version) CiliumEnvoyConfigs() CiliumEnvoyConfigInformer {
	return &ciliumEnvoyConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
