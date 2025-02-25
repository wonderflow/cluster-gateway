// +build !ignore_autogenerated

/*
Copyright 2021 The KubeVela Authors.
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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAccess) DeepCopyInto(out *ClusterAccess) {
	*out = *in
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(bool)
		**out = **in
	}
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(ClusterAccessCredential)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAccess.
func (in *ClusterAccess) DeepCopy() *ClusterAccess {
	if in == nil {
		return nil
	}
	out := new(ClusterAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAccessCredential) DeepCopyInto(out *ClusterAccessCredential) {
	*out = *in
	if in.X509 != nil {
		in, out := &in.X509, &out.X509
		*out = new(X509)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAccessCredential.
func (in *ClusterAccessCredential) DeepCopy() *ClusterAccessCredential {
	if in == nil {
		return nil
	}
	out := new(ClusterAccessCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterExtension) DeepCopyInto(out *ClusterExtension) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterExtension.
func (in *ClusterExtension) DeepCopy() *ClusterExtension {
	if in == nil {
		return nil
	}
	out := new(ClusterExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterExtension) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterExtensionFinalize) DeepCopyInto(out *ClusterExtensionFinalize) {
	*out = *in
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterExtensionFinalize.
func (in *ClusterExtensionFinalize) DeepCopy() *ClusterExtensionFinalize {
	if in == nil {
		return nil
	}
	out := new(ClusterExtensionFinalize)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterExtensionFinalize) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterExtensionList) DeepCopyInto(out *ClusterExtensionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterExtensionList.
func (in *ClusterExtensionList) DeepCopy() *ClusterExtensionList {
	if in == nil {
		return nil
	}
	out := new(ClusterExtensionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterExtensionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterExtensionProxy) DeepCopyInto(out *ClusterExtensionProxy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterExtensionProxy.
func (in *ClusterExtensionProxy) DeepCopy() *ClusterExtensionProxy {
	if in == nil {
		return nil
	}
	out := new(ClusterExtensionProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterExtensionProxyOptions) DeepCopyInto(out *ClusterExtensionProxyOptions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterExtensionProxyOptions.
func (in *ClusterExtensionProxyOptions) DeepCopy() *ClusterExtensionProxyOptions {
	if in == nil {
		return nil
	}
	out := new(ClusterExtensionProxyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterExtensionProxyOptions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterExtensionSpec) DeepCopyInto(out *ClusterExtensionSpec) {
	*out = *in
	in.Access.DeepCopyInto(&out.Access)
	if in.Finalized != nil {
		in, out := &in.Finalized, &out.Finalized
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterExtensionSpec.
func (in *ClusterExtensionSpec) DeepCopy() *ClusterExtensionSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterExtensionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterExtensionStatus) DeepCopyInto(out *ClusterExtensionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterExtensionStatus.
func (in *ClusterExtensionStatus) DeepCopy() *ClusterExtensionStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterExtensionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *X509) DeepCopyInto(out *X509) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new X509.
func (in *X509) DeepCopy() *X509 {
	if in == nil {
		return nil
	}
	out := new(X509)
	in.DeepCopyInto(out)
	return out
}
