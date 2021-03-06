// Code generated by protoc-gen-go.
// source: InRoom.proto
// DO NOT EDIT!

/*
Package logicserver is a generated protocol buffer package.

It is generated from these files:
	InRoom.proto

It has these top-level messages:
	InRoom
*/
package logicserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InRoom struct {
	UserName         *string `protobuf:"bytes,1,req,name=userName" json:"userName,omitempty"`
	State            *int32  `protobuf:"varint,2,req,name=state" json:"state,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InRoom) Reset()                    { *m = InRoom{} }
func (m *InRoom) String() string            { return proto.CompactTextString(m) }
func (*InRoom) ProtoMessage()               {}
func (*InRoom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InRoom) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *InRoom) GetState() int32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func init() {
	proto.RegisterType((*InRoom)(nil), "logicserver.InRoom")
}

func init() { proto.RegisterFile("InRoom.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 88 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xf1, 0xcc, 0x0b, 0xca,
	0xcf, 0xcf, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0xc9, 0x4f, 0xcf, 0x4c, 0x2e,
	0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x52, 0xd2, 0xe4, 0x62, 0x83, 0x48, 0x0a, 0x09, 0x70, 0x71, 0x94,
	0x02, 0x05, 0xfd, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x98, 0x34, 0x38, 0x85, 0x78, 0xb9, 0x58,
	0x8b, 0x4b, 0x12, 0x4b, 0x52, 0x25, 0x98, 0x80, 0x5c, 0x56, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb0, 0x0d, 0xeb, 0x41, 0x46, 0x00, 0x00, 0x00,
}
