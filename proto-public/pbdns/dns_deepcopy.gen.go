// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package pbdns

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using QueryRequest within kubernetes types, where deepcopy-gen is used.
func (in *QueryRequest) DeepCopyInto(out *QueryRequest) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryRequest. Required by controller-gen.
func (in *QueryRequest) DeepCopy() *QueryRequest {
	if in == nil {
		return nil
	}
	out := new(QueryRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new QueryRequest. Required by controller-gen.
func (in *QueryRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using QueryResponse within kubernetes types, where deepcopy-gen is used.
func (in *QueryResponse) DeepCopyInto(out *QueryResponse) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryResponse. Required by controller-gen.
func (in *QueryResponse) DeepCopy() *QueryResponse {
	if in == nil {
		return nil
	}
	out := new(QueryResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new QueryResponse. Required by controller-gen.
func (in *QueryResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
