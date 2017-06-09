// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&NetworkPolicy{}, func(obj interface{}) { SetObjectDefaults_NetworkPolicy(obj.(*NetworkPolicy)) })
	scheme.AddTypeDefaultingFunc(&NetworkPolicyList{}, func(obj interface{}) { SetObjectDefaults_NetworkPolicyList(obj.(*NetworkPolicyList)) })
	return nil
}

func SetObjectDefaults_NetworkPolicy(in *NetworkPolicy) {
	if in.ObjectMeta.CreationTimestamp.Time.loc != nil {
		for i := range in.ObjectMeta.CreationTimestamp.Time.loc.zone {
			a := &in.ObjectMeta.CreationTimestamp.Time.loc.zone[i]
		}
		for i := range in.ObjectMeta.CreationTimestamp.Time.loc.tx {
			a := &in.ObjectMeta.CreationTimestamp.Time.loc.tx[i]
		}
		if in.ObjectMeta.CreationTimestamp.Time.loc.cacheZone != nil {
		}
	}
	if in.ObjectMeta.DeletionTimestamp != nil {
	}
	if in.ObjectMeta.DeletionGracePeriodSeconds != nil {
	}
	for i := range in.ObjectMeta.OwnerReferences {
		a := &in.ObjectMeta.OwnerReferences[i]
		if a.Controller != nil {
		}
	}
	if in.ObjectMeta.Initializers != nil {
		for i := range in.ObjectMeta.Initializers.Pending {
			a := &in.ObjectMeta.Initializers.Pending[i]
		}
		if in.ObjectMeta.Initializers.Result != nil {
			if in.ObjectMeta.Initializers.Result.Details != nil {
				for i := range in.ObjectMeta.Initializers.Result.Details.Causes {
					a := &in.ObjectMeta.Initializers.Result.Details.Causes[i]
				}
			}
		}
	}
	for i := range in.ObjectMeta.Finalizers {
		a := &in.ObjectMeta.Finalizers[i]
	}
	for i := range in.Spec.PodSelector.MatchExpressions {
		a := &in.Spec.PodSelector.MatchExpressions[i]
	}
	for i := range in.Spec.Ingress {
		a := &in.Spec.Ingress[i]
		for j := range a.Ports {
			b := &a.Ports[j]
			SetDefaults_NetworkPolicyPort(b)
			if b.Protocol != nil {
			}
			if b.Port != nil {
			}
		}
		for j := range a.From {
			b := &a.From[j]
			if b.PodSelector != nil {
			}
		}
	}
}

func SetObjectDefaults_NetworkPolicyList(in *NetworkPolicyList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_NetworkPolicy(a)
	}
}
