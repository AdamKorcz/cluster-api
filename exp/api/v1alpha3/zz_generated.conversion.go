// +build !ignore_autogenerated_core_v1alpha3

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha3

import (
	unsafe "unsafe"

	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	apiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
	errors "sigs.k8s.io/cluster-api/errors"
	v1alpha4 "sigs.k8s.io/cluster-api/exp/api/v1alpha4"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*MachinePool)(nil), (*v1alpha4.MachinePool)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MachinePool_To_v1alpha4_MachinePool(a.(*MachinePool), b.(*v1alpha4.MachinePool), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.MachinePool)(nil), (*MachinePool)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_MachinePool_To_v1alpha3_MachinePool(a.(*v1alpha4.MachinePool), b.(*MachinePool), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachinePoolList)(nil), (*v1alpha4.MachinePoolList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MachinePoolList_To_v1alpha4_MachinePoolList(a.(*MachinePoolList), b.(*v1alpha4.MachinePoolList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.MachinePoolList)(nil), (*MachinePoolList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_MachinePoolList_To_v1alpha3_MachinePoolList(a.(*v1alpha4.MachinePoolList), b.(*MachinePoolList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachinePoolSpec)(nil), (*v1alpha4.MachinePoolSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MachinePoolSpec_To_v1alpha4_MachinePoolSpec(a.(*MachinePoolSpec), b.(*v1alpha4.MachinePoolSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.MachinePoolSpec)(nil), (*MachinePoolSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_MachinePoolSpec_To_v1alpha3_MachinePoolSpec(a.(*v1alpha4.MachinePoolSpec), b.(*MachinePoolSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachinePoolStatus)(nil), (*v1alpha4.MachinePoolStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MachinePoolStatus_To_v1alpha4_MachinePoolStatus(a.(*MachinePoolStatus), b.(*v1alpha4.MachinePoolStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha4.MachinePoolStatus)(nil), (*MachinePoolStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_MachinePoolStatus_To_v1alpha3_MachinePoolStatus(a.(*v1alpha4.MachinePoolStatus), b.(*MachinePoolStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha3_MachinePool_To_v1alpha4_MachinePool(in *MachinePool, out *v1alpha4.MachinePool, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_MachinePoolSpec_To_v1alpha4_MachinePoolSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_MachinePoolStatus_To_v1alpha4_MachinePoolStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_MachinePool_To_v1alpha4_MachinePool is an autogenerated conversion function.
func Convert_v1alpha3_MachinePool_To_v1alpha4_MachinePool(in *MachinePool, out *v1alpha4.MachinePool, s conversion.Scope) error {
	return autoConvert_v1alpha3_MachinePool_To_v1alpha4_MachinePool(in, out, s)
}

func autoConvert_v1alpha4_MachinePool_To_v1alpha3_MachinePool(in *v1alpha4.MachinePool, out *MachinePool, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha4_MachinePoolSpec_To_v1alpha3_MachinePoolSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha4_MachinePoolStatus_To_v1alpha3_MachinePoolStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_MachinePool_To_v1alpha3_MachinePool is an autogenerated conversion function.
func Convert_v1alpha4_MachinePool_To_v1alpha3_MachinePool(in *v1alpha4.MachinePool, out *MachinePool, s conversion.Scope) error {
	return autoConvert_v1alpha4_MachinePool_To_v1alpha3_MachinePool(in, out, s)
}

func autoConvert_v1alpha3_MachinePoolList_To_v1alpha4_MachinePoolList(in *MachinePoolList, out *v1alpha4.MachinePoolList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha4.MachinePool)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha3_MachinePoolList_To_v1alpha4_MachinePoolList is an autogenerated conversion function.
func Convert_v1alpha3_MachinePoolList_To_v1alpha4_MachinePoolList(in *MachinePoolList, out *v1alpha4.MachinePoolList, s conversion.Scope) error {
	return autoConvert_v1alpha3_MachinePoolList_To_v1alpha4_MachinePoolList(in, out, s)
}

func autoConvert_v1alpha4_MachinePoolList_To_v1alpha3_MachinePoolList(in *v1alpha4.MachinePoolList, out *MachinePoolList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]MachinePool)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha4_MachinePoolList_To_v1alpha3_MachinePoolList is an autogenerated conversion function.
func Convert_v1alpha4_MachinePoolList_To_v1alpha3_MachinePoolList(in *v1alpha4.MachinePoolList, out *MachinePoolList, s conversion.Scope) error {
	return autoConvert_v1alpha4_MachinePoolList_To_v1alpha3_MachinePoolList(in, out, s)
}

func autoConvert_v1alpha3_MachinePoolSpec_To_v1alpha4_MachinePoolSpec(in *MachinePoolSpec, out *v1alpha4.MachinePoolSpec, s conversion.Scope) error {
	out.ClusterName = in.ClusterName
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.Template, &out.Template, 0); err != nil {
		return err
	}
	out.Strategy = (*apiv1alpha4.MachineDeploymentStrategy)(unsafe.Pointer(in.Strategy))
	out.MinReadySeconds = (*int32)(unsafe.Pointer(in.MinReadySeconds))
	out.ProviderIDList = *(*[]string)(unsafe.Pointer(&in.ProviderIDList))
	out.FailureDomains = *(*[]string)(unsafe.Pointer(&in.FailureDomains))
	return nil
}

// Convert_v1alpha3_MachinePoolSpec_To_v1alpha4_MachinePoolSpec is an autogenerated conversion function.
func Convert_v1alpha3_MachinePoolSpec_To_v1alpha4_MachinePoolSpec(in *MachinePoolSpec, out *v1alpha4.MachinePoolSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_MachinePoolSpec_To_v1alpha4_MachinePoolSpec(in, out, s)
}

func autoConvert_v1alpha4_MachinePoolSpec_To_v1alpha3_MachinePoolSpec(in *v1alpha4.MachinePoolSpec, out *MachinePoolSpec, s conversion.Scope) error {
	out.ClusterName = in.ClusterName
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.Template, &out.Template, 0); err != nil {
		return err
	}
	out.Strategy = (*apiv1alpha3.MachineDeploymentStrategy)(unsafe.Pointer(in.Strategy))
	out.MinReadySeconds = (*int32)(unsafe.Pointer(in.MinReadySeconds))
	out.ProviderIDList = *(*[]string)(unsafe.Pointer(&in.ProviderIDList))
	out.FailureDomains = *(*[]string)(unsafe.Pointer(&in.FailureDomains))
	return nil
}

// Convert_v1alpha4_MachinePoolSpec_To_v1alpha3_MachinePoolSpec is an autogenerated conversion function.
func Convert_v1alpha4_MachinePoolSpec_To_v1alpha3_MachinePoolSpec(in *v1alpha4.MachinePoolSpec, out *MachinePoolSpec, s conversion.Scope) error {
	return autoConvert_v1alpha4_MachinePoolSpec_To_v1alpha3_MachinePoolSpec(in, out, s)
}

func autoConvert_v1alpha3_MachinePoolStatus_To_v1alpha4_MachinePoolStatus(in *MachinePoolStatus, out *v1alpha4.MachinePoolStatus, s conversion.Scope) error {
	out.NodeRefs = *(*[]v1.ObjectReference)(unsafe.Pointer(&in.NodeRefs))
	out.Replicas = in.Replicas
	out.ReadyReplicas = in.ReadyReplicas
	out.AvailableReplicas = in.AvailableReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	out.FailureReason = (*errors.MachinePoolStatusFailure)(unsafe.Pointer(in.FailureReason))
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	out.Phase = in.Phase
	out.BootstrapReady = in.BootstrapReady
	out.InfrastructureReady = in.InfrastructureReady
	out.ObservedGeneration = in.ObservedGeneration
	out.Conditions = *(*apiv1alpha4.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha3_MachinePoolStatus_To_v1alpha4_MachinePoolStatus is an autogenerated conversion function.
func Convert_v1alpha3_MachinePoolStatus_To_v1alpha4_MachinePoolStatus(in *MachinePoolStatus, out *v1alpha4.MachinePoolStatus, s conversion.Scope) error {
	return autoConvert_v1alpha3_MachinePoolStatus_To_v1alpha4_MachinePoolStatus(in, out, s)
}

func autoConvert_v1alpha4_MachinePoolStatus_To_v1alpha3_MachinePoolStatus(in *v1alpha4.MachinePoolStatus, out *MachinePoolStatus, s conversion.Scope) error {
	out.NodeRefs = *(*[]v1.ObjectReference)(unsafe.Pointer(&in.NodeRefs))
	out.Replicas = in.Replicas
	out.ReadyReplicas = in.ReadyReplicas
	out.AvailableReplicas = in.AvailableReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	out.FailureReason = (*errors.MachinePoolStatusFailure)(unsafe.Pointer(in.FailureReason))
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	out.Phase = in.Phase
	out.BootstrapReady = in.BootstrapReady
	out.InfrastructureReady = in.InfrastructureReady
	out.ObservedGeneration = in.ObservedGeneration
	out.Conditions = *(*apiv1alpha3.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha4_MachinePoolStatus_To_v1alpha3_MachinePoolStatus is an autogenerated conversion function.
func Convert_v1alpha4_MachinePoolStatus_To_v1alpha3_MachinePoolStatus(in *v1alpha4.MachinePoolStatus, out *MachinePoolStatus, s conversion.Scope) error {
	return autoConvert_v1alpha4_MachinePoolStatus_To_v1alpha3_MachinePoolStatus(in, out, s)
}
