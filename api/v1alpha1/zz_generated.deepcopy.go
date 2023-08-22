//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimplePostgreSQL) DeepCopyInto(out *SimplePostgreSQL) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimplePostgreSQL.
func (in *SimplePostgreSQL) DeepCopy() *SimplePostgreSQL {
	if in == nil {
		return nil
	}
	out := new(SimplePostgreSQL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SimplePostgreSQL) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimplePostgreSQLList) DeepCopyInto(out *SimplePostgreSQLList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SimplePostgreSQL, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimplePostgreSQLList.
func (in *SimplePostgreSQLList) DeepCopy() *SimplePostgreSQLList {
	if in == nil {
		return nil
	}
	out := new(SimplePostgreSQLList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SimplePostgreSQLList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimplePostgreSQLSpec) DeepCopyInto(out *SimplePostgreSQLSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimplePostgreSQLSpec.
func (in *SimplePostgreSQLSpec) DeepCopy() *SimplePostgreSQLSpec {
	if in == nil {
		return nil
	}
	out := new(SimplePostgreSQLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimplePostgreSQLStatus) DeepCopyInto(out *SimplePostgreSQLStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimplePostgreSQLStatus.
func (in *SimplePostgreSQLStatus) DeepCopy() *SimplePostgreSQLStatus {
	if in == nil {
		return nil
	}
	out := new(SimplePostgreSQLStatus)
	in.DeepCopyInto(out)
	return out
}
