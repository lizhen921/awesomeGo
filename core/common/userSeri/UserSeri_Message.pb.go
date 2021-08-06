// Code generated by protoc-gen-go. DO NOT EDIT.
// source: main.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	main.proto

It has these top-level messages:
	UserSeri
*/
package main

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

type UserSeri struct {
	SeriList         []*UserSeri_Seri `protobuf:"bytes,1,rep,name=SeriList" json:"SeriList,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *UserSeri) Reset()                    { *m = UserSeri{} }
func (m *UserSeri) String() string            { return proto.CompactTextString(m) }
func (*UserSeri) ProtoMessage()               {}
func (*UserSeri) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserSeri) GetSeriList() []*UserSeri_Seri {
	if m != nil {
		return m.SeriList
	}
	return nil
}

type UserSeri_Seri struct {
	SeriId           *int32   `protobuf:"varint,1,opt,name=SeriId" json:"SeriId,omitempty"`
	Score            *float32 `protobuf:"fixed32,2,opt,name=Score" json:"Score,omitempty"`
	SegId            *int32   `protobuf:"varint,3,opt,name=SegId" json:"SegId,omitempty"`
	SegScore         *float32 `protobuf:"fixed32,4,opt,name=SegScore" json:"SegScore,omitempty"`
	Pack             *string  `protobuf:"bytes,5,opt,name=Pack" json:"Pack,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *UserSeri_Seri) Reset()                    { *m = UserSeri_Seri{} }
func (m *UserSeri_Seri) String() string            { return proto.CompactTextString(m) }
func (*UserSeri_Seri) ProtoMessage()               {}
func (*UserSeri_Seri) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *UserSeri_Seri) GetSeriId() int32 {
	if m != nil && m.SeriId != nil {
		return *m.SeriId
	}
	return 0
}

func (m *UserSeri_Seri) GetScore() float32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *UserSeri_Seri) GetSegId() int32 {
	if m != nil && m.SegId != nil {
		return *m.SegId
	}
	return 0
}

func (m *UserSeri_Seri) GetSegScore() float32 {
	if m != nil && m.SegScore != nil {
		return *m.SegScore
	}
	return 0
}

func (m *UserSeri_Seri) GetPack() string {
	if m != nil && m.Pack != nil {
		return *m.Pack
	}
	return ""
}

func init() {
	proto.RegisterType((*UserSeri)(nil), "UserSeri")
	proto.RegisterType((*UserSeri_Seri)(nil), "UserSeri.Seri")
}

func init() { proto.RegisterFile("main.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0b, 0x2d, 0x4e, 0x2d,
	0x0a, 0x4e, 0x2d, 0xca, 0x8c, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x57, 0x6a, 0x62, 0xe4, 0xe2, 0x80, 0x49, 0x09, 0x29, 0x70, 0x71, 0x80, 0x68, 0x9f,
	0xcc, 0xe2, 0x12, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x3e, 0x3d, 0x98, 0xa4, 0x1e, 0x88,
	0x90, 0x0a, 0xe1, 0x62, 0x01, 0xab, 0xe4, 0xe3, 0x62, 0x03, 0xd1, 0x9e, 0x29, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0xac, 0x42, 0xbc, 0x5c, 0xac, 0xc1, 0xc9, 0xf9, 0x45, 0xa9, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0x4c, 0x60, 0x6e, 0x6a, 0xba, 0x67, 0x8a, 0x04, 0x33, 0x58, 0x56, 0x00, 0x64, 0x6e, 0x3a,
	0x44, 0x01, 0x0b, 0x58, 0x01, 0x0f, 0x17, 0x4b, 0x40, 0x62, 0x72, 0xb6, 0x04, 0xab, 0x02, 0xa3,
	0x06, 0x27, 0x20, 0x00, 0x00, 0xff, 0xff, 0x61, 0x6a, 0x1a, 0x81, 0x9d, 0x00, 0x00, 0x00,
}