// +build !ignore_autogenerated

/*
Copyright 2021 Cisco Systems, Inc. and/or its affiliates.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioControlPlane) DeepCopyInto(out *IstioControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = (*in).DeepCopy()
	}
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioControlPlane.
func (in *IstioControlPlane) DeepCopy() *IstioControlPlane {
	if in == nil {
		return nil
	}
	out := new(IstioControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IstioControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioControlPlaneList) DeepCopyInto(out *IstioControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IstioControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioControlPlaneList.
func (in *IstioControlPlaneList) DeepCopy() *IstioControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(IstioControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IstioControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioControlPlaneProperties) DeepCopyInto(out *IstioControlPlaneProperties) {
	*out = *in
	if in.Mesh != nil {
		in, out := &in.Mesh, &out.Mesh
		*out = new(IstioMesh)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioControlPlaneProperties.
func (in *IstioControlPlaneProperties) DeepCopy() *IstioControlPlaneProperties {
	if in == nil {
		return nil
	}
	out := new(IstioControlPlaneProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioControlPlaneWithProperties) DeepCopyInto(out *IstioControlPlaneWithProperties) {
	*out = *in
	if in.IstioControlPlane != nil {
		in, out := &in.IstioControlPlane, &out.IstioControlPlane
		*out = new(IstioControlPlane)
		(*in).DeepCopyInto(*out)
	}
	in.Properties.DeepCopyInto(&out.Properties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioControlPlaneWithProperties.
func (in *IstioControlPlaneWithProperties) DeepCopy() *IstioControlPlaneWithProperties {
	if in == nil {
		return nil
	}
	out := new(IstioControlPlaneWithProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioMesh) DeepCopyInto(out *IstioMesh) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = (*in).DeepCopy()
	}
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioMesh.
func (in *IstioMesh) DeepCopy() *IstioMesh {
	if in == nil {
		return nil
	}
	out := new(IstioMesh)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IstioMesh) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioMeshGateway) DeepCopyInto(out *IstioMeshGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = (*in).DeepCopy()
	}
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioMeshGateway.
func (in *IstioMeshGateway) DeepCopy() *IstioMeshGateway {
	if in == nil {
		return nil
	}
	out := new(IstioMeshGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IstioMeshGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioMeshGatewayList) DeepCopyInto(out *IstioMeshGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IstioMeshGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioMeshGatewayList.
func (in *IstioMeshGatewayList) DeepCopy() *IstioMeshGatewayList {
	if in == nil {
		return nil
	}
	out := new(IstioMeshGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IstioMeshGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioMeshGatewayProperties) DeepCopyInto(out *IstioMeshGatewayProperties) {
	*out = *in
	if in.EnablePrometheusMerge != nil {
		in, out := &in.EnablePrometheusMerge, &out.EnablePrometheusMerge
		*out = new(bool)
		**out = **in
	}
	if in.IstioControlPlane != nil {
		in, out := &in.IstioControlPlane, &out.IstioControlPlane
		*out = new(IstioControlPlane)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioMeshGatewayProperties.
func (in *IstioMeshGatewayProperties) DeepCopy() *IstioMeshGatewayProperties {
	if in == nil {
		return nil
	}
	out := new(IstioMeshGatewayProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioMeshGatewayWithProperties) DeepCopyInto(out *IstioMeshGatewayWithProperties) {
	*out = *in
	if in.IstioMeshGateway != nil {
		in, out := &in.IstioMeshGateway, &out.IstioMeshGateway
		*out = new(IstioMeshGateway)
		(*in).DeepCopyInto(*out)
	}
	in.Properties.DeepCopyInto(&out.Properties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioMeshGatewayWithProperties.
func (in *IstioMeshGatewayWithProperties) DeepCopy() *IstioMeshGatewayWithProperties {
	if in == nil {
		return nil
	}
	out := new(IstioMeshGatewayWithProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioMeshList) DeepCopyInto(out *IstioMeshList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IstioMesh, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioMeshList.
func (in *IstioMeshList) DeepCopy() *IstioMeshList {
	if in == nil {
		return nil
	}
	out := new(IstioMeshList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IstioMeshList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerIstioControlPlane) DeepCopyInto(out *PeerIstioControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = (*in).DeepCopy()
	}
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerIstioControlPlane.
func (in *PeerIstioControlPlane) DeepCopy() *PeerIstioControlPlane {
	if in == nil {
		return nil
	}
	out := new(PeerIstioControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PeerIstioControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerIstioControlPlaneList) DeepCopyInto(out *PeerIstioControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PeerIstioControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerIstioControlPlaneList.
func (in *PeerIstioControlPlaneList) DeepCopy() *PeerIstioControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(PeerIstioControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PeerIstioControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in SortableIstioControlPlaneItems) DeepCopyInto(out *SortableIstioControlPlaneItems) {
	{
		in := &in
		*out = make(SortableIstioControlPlaneItems, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SortableIstioControlPlaneItems.
func (in SortableIstioControlPlaneItems) DeepCopy() SortableIstioControlPlaneItems {
	if in == nil {
		return nil
	}
	out := new(SortableIstioControlPlaneItems)
	in.DeepCopyInto(out)
	return *out
}
