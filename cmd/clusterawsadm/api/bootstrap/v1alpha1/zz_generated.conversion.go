//go:build !ignore_autogenerated_conversions
// +build !ignore_autogenerated_conversions

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

package v1alpha1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1beta1 "sigs.k8s.io/cluster-api-provider-aws/api/v1beta1"
	v1beta1 "sigs.k8s.io/cluster-api-provider-aws/cmd/clusterawsadm/api/bootstrap/v1beta1"
	iamapiv1beta1 "sigs.k8s.io/cluster-api-provider-aws/iam/api/v1beta1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AWSIAMConfiguration)(nil), (*v1beta1.AWSIAMConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AWSIAMConfiguration_To_v1beta1_AWSIAMConfiguration(a.(*AWSIAMConfiguration), b.(*v1beta1.AWSIAMConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSIAMConfiguration)(nil), (*AWSIAMConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSIAMConfiguration_To_v1alpha1_AWSIAMConfiguration(a.(*v1beta1.AWSIAMConfiguration), b.(*AWSIAMConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSIAMConfigurationSpec)(nil), (*v1beta1.AWSIAMConfigurationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AWSIAMConfigurationSpec_To_v1beta1_AWSIAMConfigurationSpec(a.(*AWSIAMConfigurationSpec), b.(*v1beta1.AWSIAMConfigurationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSIAMRoleSpec)(nil), (*v1beta1.AWSIAMRoleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AWSIAMRoleSpec_To_v1beta1_AWSIAMRoleSpec(a.(*AWSIAMRoleSpec), b.(*v1beta1.AWSIAMRoleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSIAMRoleSpec)(nil), (*AWSIAMRoleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSIAMRoleSpec_To_v1alpha1_AWSIAMRoleSpec(a.(*v1beta1.AWSIAMRoleSpec), b.(*AWSIAMRoleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BootstrapUser)(nil), (*v1beta1.BootstrapUser)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_BootstrapUser_To_v1beta1_BootstrapUser(a.(*BootstrapUser), b.(*v1beta1.BootstrapUser), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.BootstrapUser)(nil), (*BootstrapUser)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_BootstrapUser_To_v1alpha1_BootstrapUser(a.(*v1beta1.BootstrapUser), b.(*BootstrapUser), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterAPIControllers)(nil), (*v1beta1.ClusterAPIControllers)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ClusterAPIControllers_To_v1beta1_ClusterAPIControllers(a.(*ClusterAPIControllers), b.(*v1beta1.ClusterAPIControllers), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ClusterAPIControllers)(nil), (*ClusterAPIControllers)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterAPIControllers_To_v1alpha1_ClusterAPIControllers(a.(*v1beta1.ClusterAPIControllers), b.(*ClusterAPIControllers), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControlPlane)(nil), (*v1beta1.ControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControlPlane_To_v1beta1_ControlPlane(a.(*ControlPlane), b.(*v1beta1.ControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ControlPlane)(nil), (*ControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ControlPlane_To_v1alpha1_ControlPlane(a.(*v1beta1.ControlPlane), b.(*ControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfig)(nil), (*v1beta1.EKSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EKSConfig_To_v1beta1_EKSConfig(a.(*EKSConfig), b.(*v1beta1.EKSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EKSConfig)(nil), (*EKSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfig_To_v1alpha1_EKSConfig(a.(*v1beta1.EKSConfig), b.(*EKSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EventBridgeConfig)(nil), (*v1beta1.EventBridgeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EventBridgeConfig_To_v1beta1_EventBridgeConfig(a.(*EventBridgeConfig), b.(*v1beta1.EventBridgeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EventBridgeConfig)(nil), (*EventBridgeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EventBridgeConfig_To_v1alpha1_EventBridgeConfig(a.(*v1beta1.EventBridgeConfig), b.(*EventBridgeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Nodes)(nil), (*v1beta1.Nodes)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Nodes_To_v1beta1_Nodes(a.(*Nodes), b.(*v1beta1.Nodes), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.Nodes)(nil), (*Nodes)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Nodes_To_v1alpha1_Nodes(a.(*v1beta1.Nodes), b.(*Nodes), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.AWSIAMConfigurationSpec)(nil), (*AWSIAMConfigurationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSIAMConfigurationSpec_To_v1alpha1_AWSIAMConfigurationSpec(a.(*v1beta1.AWSIAMConfigurationSpec), b.(*AWSIAMConfigurationSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_AWSIAMConfiguration_To_v1beta1_AWSIAMConfiguration(in *AWSIAMConfiguration, out *v1beta1.AWSIAMConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_AWSIAMConfigurationSpec_To_v1beta1_AWSIAMConfigurationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_AWSIAMConfiguration_To_v1beta1_AWSIAMConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AWSIAMConfiguration_To_v1beta1_AWSIAMConfiguration(in *AWSIAMConfiguration, out *v1beta1.AWSIAMConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AWSIAMConfiguration_To_v1beta1_AWSIAMConfiguration(in, out, s)
}

func autoConvert_v1beta1_AWSIAMConfiguration_To_v1alpha1_AWSIAMConfiguration(in *v1beta1.AWSIAMConfiguration, out *AWSIAMConfiguration, s conversion.Scope) error {
	if err := Convert_v1beta1_AWSIAMConfigurationSpec_To_v1alpha1_AWSIAMConfigurationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_AWSIAMConfiguration_To_v1alpha1_AWSIAMConfiguration is an autogenerated conversion function.
func Convert_v1beta1_AWSIAMConfiguration_To_v1alpha1_AWSIAMConfiguration(in *v1beta1.AWSIAMConfiguration, out *AWSIAMConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSIAMConfiguration_To_v1alpha1_AWSIAMConfiguration(in, out, s)
}

func autoConvert_v1alpha1_AWSIAMConfigurationSpec_To_v1beta1_AWSIAMConfigurationSpec(in *AWSIAMConfigurationSpec, out *v1beta1.AWSIAMConfigurationSpec, s conversion.Scope) error {
	out.NamePrefix = in.NamePrefix
	out.NameSuffix = (*string)(unsafe.Pointer(in.NameSuffix))
	if err := Convert_v1alpha1_ControlPlane_To_v1beta1_ControlPlane(&in.ControlPlane, &out.ControlPlane, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ClusterAPIControllers_To_v1beta1_ClusterAPIControllers(&in.ClusterAPIControllers, &out.ClusterAPIControllers, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_Nodes_To_v1beta1_Nodes(&in.Nodes, &out.Nodes, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_BootstrapUser_To_v1beta1_BootstrapUser(&in.BootstrapUser, &out.BootstrapUser, s); err != nil {
		return err
	}
	out.StackName = in.StackName
	out.Region = in.Region
	out.EKS = (*v1beta1.EKSConfig)(unsafe.Pointer(in.EKS))
	out.EventBridge = (*v1beta1.EventBridgeConfig)(unsafe.Pointer(in.EventBridge))
	out.Partition = in.Partition
	out.SecureSecretsBackends = *(*[]apiv1beta1.SecretBackend)(unsafe.Pointer(&in.SecureSecretsBackends))
	return nil
}

// Convert_v1alpha1_AWSIAMConfigurationSpec_To_v1beta1_AWSIAMConfigurationSpec is an autogenerated conversion function.
func Convert_v1alpha1_AWSIAMConfigurationSpec_To_v1beta1_AWSIAMConfigurationSpec(in *AWSIAMConfigurationSpec, out *v1beta1.AWSIAMConfigurationSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_AWSIAMConfigurationSpec_To_v1beta1_AWSIAMConfigurationSpec(in, out, s)
}

func autoConvert_v1beta1_AWSIAMConfigurationSpec_To_v1alpha1_AWSIAMConfigurationSpec(in *v1beta1.AWSIAMConfigurationSpec, out *AWSIAMConfigurationSpec, s conversion.Scope) error {
	out.NamePrefix = in.NamePrefix
	out.NameSuffix = (*string)(unsafe.Pointer(in.NameSuffix))
	if err := Convert_v1beta1_ControlPlane_To_v1alpha1_ControlPlane(&in.ControlPlane, &out.ControlPlane, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ClusterAPIControllers_To_v1alpha1_ClusterAPIControllers(&in.ClusterAPIControllers, &out.ClusterAPIControllers, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_Nodes_To_v1alpha1_Nodes(&in.Nodes, &out.Nodes, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_BootstrapUser_To_v1alpha1_BootstrapUser(&in.BootstrapUser, &out.BootstrapUser, s); err != nil {
		return err
	}
	out.StackName = in.StackName
	// WARNING: in.StackTags requires manual conversion: does not exist in peer-type
	out.Region = in.Region
	out.EKS = (*EKSConfig)(unsafe.Pointer(in.EKS))
	out.EventBridge = (*EventBridgeConfig)(unsafe.Pointer(in.EventBridge))
	out.Partition = in.Partition
	out.SecureSecretsBackends = *(*[]apiv1beta1.SecretBackend)(unsafe.Pointer(&in.SecureSecretsBackends))
	return nil
}

func autoConvert_v1alpha1_AWSIAMRoleSpec_To_v1beta1_AWSIAMRoleSpec(in *AWSIAMRoleSpec, out *v1beta1.AWSIAMRoleSpec, s conversion.Scope) error {
	out.Disable = in.Disable
	out.ExtraPolicyAttachments = *(*[]string)(unsafe.Pointer(&in.ExtraPolicyAttachments))
	out.ExtraStatements = *(*[]iamapiv1beta1.StatementEntry)(unsafe.Pointer(&in.ExtraStatements))
	out.TrustStatements = *(*[]iamapiv1beta1.StatementEntry)(unsafe.Pointer(&in.TrustStatements))
	out.Tags = *(*apiv1beta1.Tags)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_v1alpha1_AWSIAMRoleSpec_To_v1beta1_AWSIAMRoleSpec is an autogenerated conversion function.
func Convert_v1alpha1_AWSIAMRoleSpec_To_v1beta1_AWSIAMRoleSpec(in *AWSIAMRoleSpec, out *v1beta1.AWSIAMRoleSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_AWSIAMRoleSpec_To_v1beta1_AWSIAMRoleSpec(in, out, s)
}

func autoConvert_v1beta1_AWSIAMRoleSpec_To_v1alpha1_AWSIAMRoleSpec(in *v1beta1.AWSIAMRoleSpec, out *AWSIAMRoleSpec, s conversion.Scope) error {
	out.Disable = in.Disable
	out.ExtraPolicyAttachments = *(*[]string)(unsafe.Pointer(&in.ExtraPolicyAttachments))
	out.ExtraStatements = *(*[]iamapiv1beta1.StatementEntry)(unsafe.Pointer(&in.ExtraStatements))
	out.TrustStatements = *(*[]iamapiv1beta1.StatementEntry)(unsafe.Pointer(&in.TrustStatements))
	out.Tags = *(*apiv1beta1.Tags)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_v1beta1_AWSIAMRoleSpec_To_v1alpha1_AWSIAMRoleSpec is an autogenerated conversion function.
func Convert_v1beta1_AWSIAMRoleSpec_To_v1alpha1_AWSIAMRoleSpec(in *v1beta1.AWSIAMRoleSpec, out *AWSIAMRoleSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSIAMRoleSpec_To_v1alpha1_AWSIAMRoleSpec(in, out, s)
}

func autoConvert_v1alpha1_BootstrapUser_To_v1beta1_BootstrapUser(in *BootstrapUser, out *v1beta1.BootstrapUser, s conversion.Scope) error {
	out.Enable = in.Enable
	out.UserName = in.UserName
	out.GroupName = in.GroupName
	out.ExtraPolicyAttachments = *(*[]string)(unsafe.Pointer(&in.ExtraPolicyAttachments))
	out.ExtraGroups = *(*[]string)(unsafe.Pointer(&in.ExtraGroups))
	out.ExtraStatements = *(*[]iamapiv1beta1.StatementEntry)(unsafe.Pointer(&in.ExtraStatements))
	out.Tags = *(*apiv1beta1.Tags)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_v1alpha1_BootstrapUser_To_v1beta1_BootstrapUser is an autogenerated conversion function.
func Convert_v1alpha1_BootstrapUser_To_v1beta1_BootstrapUser(in *BootstrapUser, out *v1beta1.BootstrapUser, s conversion.Scope) error {
	return autoConvert_v1alpha1_BootstrapUser_To_v1beta1_BootstrapUser(in, out, s)
}

func autoConvert_v1beta1_BootstrapUser_To_v1alpha1_BootstrapUser(in *v1beta1.BootstrapUser, out *BootstrapUser, s conversion.Scope) error {
	out.Enable = in.Enable
	out.UserName = in.UserName
	out.GroupName = in.GroupName
	out.ExtraPolicyAttachments = *(*[]string)(unsafe.Pointer(&in.ExtraPolicyAttachments))
	out.ExtraGroups = *(*[]string)(unsafe.Pointer(&in.ExtraGroups))
	out.ExtraStatements = *(*[]iamapiv1beta1.StatementEntry)(unsafe.Pointer(&in.ExtraStatements))
	out.Tags = *(*apiv1beta1.Tags)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_v1beta1_BootstrapUser_To_v1alpha1_BootstrapUser is an autogenerated conversion function.
func Convert_v1beta1_BootstrapUser_To_v1alpha1_BootstrapUser(in *v1beta1.BootstrapUser, out *BootstrapUser, s conversion.Scope) error {
	return autoConvert_v1beta1_BootstrapUser_To_v1alpha1_BootstrapUser(in, out, s)
}

func autoConvert_v1alpha1_ClusterAPIControllers_To_v1beta1_ClusterAPIControllers(in *ClusterAPIControllers, out *v1beta1.ClusterAPIControllers, s conversion.Scope) error {
	if err := Convert_v1alpha1_AWSIAMRoleSpec_To_v1beta1_AWSIAMRoleSpec(&in.AWSIAMRoleSpec, &out.AWSIAMRoleSpec, s); err != nil {
		return err
	}
	out.AllowedEC2InstanceProfiles = *(*[]string)(unsafe.Pointer(&in.AllowedEC2InstanceProfiles))
	return nil
}

// Convert_v1alpha1_ClusterAPIControllers_To_v1beta1_ClusterAPIControllers is an autogenerated conversion function.
func Convert_v1alpha1_ClusterAPIControllers_To_v1beta1_ClusterAPIControllers(in *ClusterAPIControllers, out *v1beta1.ClusterAPIControllers, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterAPIControllers_To_v1beta1_ClusterAPIControllers(in, out, s)
}

func autoConvert_v1beta1_ClusterAPIControllers_To_v1alpha1_ClusterAPIControllers(in *v1beta1.ClusterAPIControllers, out *ClusterAPIControllers, s conversion.Scope) error {
	if err := Convert_v1beta1_AWSIAMRoleSpec_To_v1alpha1_AWSIAMRoleSpec(&in.AWSIAMRoleSpec, &out.AWSIAMRoleSpec, s); err != nil {
		return err
	}
	out.AllowedEC2InstanceProfiles = *(*[]string)(unsafe.Pointer(&in.AllowedEC2InstanceProfiles))
	return nil
}

// Convert_v1beta1_ClusterAPIControllers_To_v1alpha1_ClusterAPIControllers is an autogenerated conversion function.
func Convert_v1beta1_ClusterAPIControllers_To_v1alpha1_ClusterAPIControllers(in *v1beta1.ClusterAPIControllers, out *ClusterAPIControllers, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterAPIControllers_To_v1alpha1_ClusterAPIControllers(in, out, s)
}

func autoConvert_v1alpha1_ControlPlane_To_v1beta1_ControlPlane(in *ControlPlane, out *v1beta1.ControlPlane, s conversion.Scope) error {
	if err := Convert_v1alpha1_AWSIAMRoleSpec_To_v1beta1_AWSIAMRoleSpec(&in.AWSIAMRoleSpec, &out.AWSIAMRoleSpec, s); err != nil {
		return err
	}
	out.DisableClusterAPIControllerPolicyAttachment = in.DisableClusterAPIControllerPolicyAttachment
	out.DisableCloudProviderPolicy = in.DisableCloudProviderPolicy
	out.EnableCSIPolicy = in.EnableCSIPolicy
	return nil
}

// Convert_v1alpha1_ControlPlane_To_v1beta1_ControlPlane is an autogenerated conversion function.
func Convert_v1alpha1_ControlPlane_To_v1beta1_ControlPlane(in *ControlPlane, out *v1beta1.ControlPlane, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControlPlane_To_v1beta1_ControlPlane(in, out, s)
}

func autoConvert_v1beta1_ControlPlane_To_v1alpha1_ControlPlane(in *v1beta1.ControlPlane, out *ControlPlane, s conversion.Scope) error {
	if err := Convert_v1beta1_AWSIAMRoleSpec_To_v1alpha1_AWSIAMRoleSpec(&in.AWSIAMRoleSpec, &out.AWSIAMRoleSpec, s); err != nil {
		return err
	}
	out.DisableClusterAPIControllerPolicyAttachment = in.DisableClusterAPIControllerPolicyAttachment
	out.DisableCloudProviderPolicy = in.DisableCloudProviderPolicy
	out.EnableCSIPolicy = in.EnableCSIPolicy
	return nil
}

// Convert_v1beta1_ControlPlane_To_v1alpha1_ControlPlane is an autogenerated conversion function.
func Convert_v1beta1_ControlPlane_To_v1alpha1_ControlPlane(in *v1beta1.ControlPlane, out *ControlPlane, s conversion.Scope) error {
	return autoConvert_v1beta1_ControlPlane_To_v1alpha1_ControlPlane(in, out, s)
}

func autoConvert_v1alpha1_EKSConfig_To_v1beta1_EKSConfig(in *EKSConfig, out *v1beta1.EKSConfig, s conversion.Scope) error {
	out.Disable = in.Disable
	out.AllowIAMRoleCreation = in.AllowIAMRoleCreation
	out.EnableUserEKSConsolePolicy = in.EnableUserEKSConsolePolicy
	if err := Convert_v1alpha1_AWSIAMRoleSpec_To_v1beta1_AWSIAMRoleSpec(&in.DefaultControlPlaneRole, &out.DefaultControlPlaneRole, s); err != nil {
		return err
	}
	out.ManagedMachinePool = (*v1beta1.AWSIAMRoleSpec)(unsafe.Pointer(in.ManagedMachinePool))
	out.Fargate = (*v1beta1.AWSIAMRoleSpec)(unsafe.Pointer(in.Fargate))
	out.KMSAliasPrefix = in.KMSAliasPrefix
	return nil
}

// Convert_v1alpha1_EKSConfig_To_v1beta1_EKSConfig is an autogenerated conversion function.
func Convert_v1alpha1_EKSConfig_To_v1beta1_EKSConfig(in *EKSConfig, out *v1beta1.EKSConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_EKSConfig_To_v1beta1_EKSConfig(in, out, s)
}

func autoConvert_v1beta1_EKSConfig_To_v1alpha1_EKSConfig(in *v1beta1.EKSConfig, out *EKSConfig, s conversion.Scope) error {
	out.Disable = in.Disable
	out.AllowIAMRoleCreation = in.AllowIAMRoleCreation
	out.EnableUserEKSConsolePolicy = in.EnableUserEKSConsolePolicy
	if err := Convert_v1beta1_AWSIAMRoleSpec_To_v1alpha1_AWSIAMRoleSpec(&in.DefaultControlPlaneRole, &out.DefaultControlPlaneRole, s); err != nil {
		return err
	}
	out.ManagedMachinePool = (*AWSIAMRoleSpec)(unsafe.Pointer(in.ManagedMachinePool))
	out.Fargate = (*AWSIAMRoleSpec)(unsafe.Pointer(in.Fargate))
	out.KMSAliasPrefix = in.KMSAliasPrefix
	return nil
}

// Convert_v1beta1_EKSConfig_To_v1alpha1_EKSConfig is an autogenerated conversion function.
func Convert_v1beta1_EKSConfig_To_v1alpha1_EKSConfig(in *v1beta1.EKSConfig, out *EKSConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfig_To_v1alpha1_EKSConfig(in, out, s)
}

func autoConvert_v1alpha1_EventBridgeConfig_To_v1beta1_EventBridgeConfig(in *EventBridgeConfig, out *v1beta1.EventBridgeConfig, s conversion.Scope) error {
	out.Enable = in.Enable
	return nil
}

// Convert_v1alpha1_EventBridgeConfig_To_v1beta1_EventBridgeConfig is an autogenerated conversion function.
func Convert_v1alpha1_EventBridgeConfig_To_v1beta1_EventBridgeConfig(in *EventBridgeConfig, out *v1beta1.EventBridgeConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_EventBridgeConfig_To_v1beta1_EventBridgeConfig(in, out, s)
}

func autoConvert_v1beta1_EventBridgeConfig_To_v1alpha1_EventBridgeConfig(in *v1beta1.EventBridgeConfig, out *EventBridgeConfig, s conversion.Scope) error {
	out.Enable = in.Enable
	return nil
}

// Convert_v1beta1_EventBridgeConfig_To_v1alpha1_EventBridgeConfig is an autogenerated conversion function.
func Convert_v1beta1_EventBridgeConfig_To_v1alpha1_EventBridgeConfig(in *v1beta1.EventBridgeConfig, out *EventBridgeConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_EventBridgeConfig_To_v1alpha1_EventBridgeConfig(in, out, s)
}

func autoConvert_v1alpha1_Nodes_To_v1beta1_Nodes(in *Nodes, out *v1beta1.Nodes, s conversion.Scope) error {
	if err := Convert_v1alpha1_AWSIAMRoleSpec_To_v1beta1_AWSIAMRoleSpec(&in.AWSIAMRoleSpec, &out.AWSIAMRoleSpec, s); err != nil {
		return err
	}
	out.DisableCloudProviderPolicy = in.DisableCloudProviderPolicy
	out.EC2ContainerRegistryReadOnly = in.EC2ContainerRegistryReadOnly
	return nil
}

// Convert_v1alpha1_Nodes_To_v1beta1_Nodes is an autogenerated conversion function.
func Convert_v1alpha1_Nodes_To_v1beta1_Nodes(in *Nodes, out *v1beta1.Nodes, s conversion.Scope) error {
	return autoConvert_v1alpha1_Nodes_To_v1beta1_Nodes(in, out, s)
}

func autoConvert_v1beta1_Nodes_To_v1alpha1_Nodes(in *v1beta1.Nodes, out *Nodes, s conversion.Scope) error {
	if err := Convert_v1beta1_AWSIAMRoleSpec_To_v1alpha1_AWSIAMRoleSpec(&in.AWSIAMRoleSpec, &out.AWSIAMRoleSpec, s); err != nil {
		return err
	}
	out.DisableCloudProviderPolicy = in.DisableCloudProviderPolicy
	out.EC2ContainerRegistryReadOnly = in.EC2ContainerRegistryReadOnly
	return nil
}

// Convert_v1beta1_Nodes_To_v1alpha1_Nodes is an autogenerated conversion function.
func Convert_v1beta1_Nodes_To_v1alpha1_Nodes(in *v1beta1.Nodes, out *Nodes, s conversion.Scope) error {
	return autoConvert_v1beta1_Nodes_To_v1alpha1_Nodes(in, out, s)
}
