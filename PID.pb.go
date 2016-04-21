// Code generated by protoc-gen-go.
// source: PID.proto
// DO NOT EDIT!

/*
Package gam is a generated protocol buffer package.

It is generated from these files:
	PID.proto

It has these top-level messages:
	PID
*/
package gam

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type PID struct {
	Node string `protobuf:"bytes,1,opt,name=Node,json=node" json:"Node,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=Host,json=host" json:"Host,omitempty"`
	Id   uint64 `protobuf:"varint,3,opt,name=Id,json=id" json:"Id,omitempty"`
}

func (m *PID) Reset()                    { *m = PID{} }
func (m *PID) String() string            { return proto.CompactTextString(m) }
func (*PID) ProtoMessage()               {}
func (*PID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*PID)(nil), "gam.PID")
}

var fileDescriptor0 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x0c, 0xf0, 0x74, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x4f, 0xcc, 0x55, 0xb2, 0xe5, 0x62, 0x06, 0x8a,
	0x08, 0x09, 0x71, 0xb1, 0xf8, 0xe5, 0xa7, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1,
	0xe4, 0x01, 0xd9, 0x20, 0x31, 0x8f, 0xfc, 0xe2, 0x12, 0x09, 0x26, 0x88, 0x58, 0x06, 0x90, 0x2d,
	0xc4, 0xc7, 0xc5, 0xe4, 0x99, 0x22, 0xc1, 0x0c, 0x14, 0x61, 0x09, 0x62, 0xca, 0x4c, 0x49, 0x62,
	0x03, 0x1b, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xec, 0x95, 0x41, 0xf2, 0x57, 0x00, 0x00,
	0x00,
}
