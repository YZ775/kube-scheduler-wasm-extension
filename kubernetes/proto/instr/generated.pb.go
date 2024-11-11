//
//Copyright The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// This file was autogenerated by go-to-protobuf. Do not edit it manually!

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v5.28.3
// source: k8s.io/apimachinery/pkg/util/intstr/generated.proto

package intstr

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// IntOrString is a type that can hold an int32 or a string.  When used in
// JSON or YAML marshalling and unmarshalling, it produces or consumes the
// inner type.  This allows you to have, for example, a JSON field that can
// accept a name or number.
// TODO: Rename to Int32OrString
//
// +protobuf=true
// +protobuf.options.(gogoproto.goproto_stringer)=false
// +k8s:openapi-gen=true
type IntOrString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   *int64  `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	IntVal *int32  `protobuf:"varint,2,opt,name=intVal" json:"intVal,omitempty"`
	StrVal *string `protobuf:"bytes,3,opt,name=strVal" json:"strVal,omitempty"`
}

func (x *IntOrString) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *IntOrString) GetType() int64 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *IntOrString) GetIntVal() int32 {
	if x != nil && x.IntVal != nil {
		return *x.IntVal
	}
	return 0
}

func (x *IntOrString) GetStrVal() string {
	if x != nil && x.StrVal != nil {
		return *x.StrVal
	}
	return ""
}
