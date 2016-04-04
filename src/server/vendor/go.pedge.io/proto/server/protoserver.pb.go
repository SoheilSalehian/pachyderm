// Code generated by protoc-gen-go.
// source: server/protoserver.proto
// DO NOT EDIT!

package protoserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ServerStarted struct {
	Port     uint32 `protobuf:"varint,1,opt,name=port" json:"port,omitempty"`
	HttpPort uint32 `protobuf:"varint,2,opt,name=http_port" json:"http_port,omitempty"`
}

func (m *ServerStarted) Reset()                    { *m = ServerStarted{} }
func (m *ServerStarted) String() string            { return proto.CompactTextString(m) }
func (*ServerStarted) ProtoMessage()               {}
func (*ServerStarted) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServerFinished struct {
	Error    string                    `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Duration *google_protobuf.Duration `protobuf:"bytes,2,opt,name=duration" json:"duration,omitempty"`
}

func (m *ServerFinished) Reset()                    { *m = ServerFinished{} }
func (m *ServerFinished) String() string            { return proto.CompactTextString(m) }
func (*ServerFinished) ProtoMessage()               {}
func (*ServerFinished) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServerFinished) GetDuration() *google_protobuf.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerStarted)(nil), "protoserver.ServerStarted")
	proto.RegisterType((*ServerFinished)(nil), "protoserver.ServerFinished")
}

var fileDescriptor0 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0xb0, 0xf5, 0xc0, 0x6c, 0x21, 0x6e, 0x24,
	0x21, 0x29, 0xb9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x88, 0xb2, 0xa4, 0xd2, 0x34, 0xfd, 0x94,
	0xd2, 0xa2, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0x88, 0x62, 0x25, 0x03, 0x2e, 0xde, 0x60, 0xb0, 0xca,
	0xe0, 0x92, 0xc4, 0xa2, 0x92, 0xd4, 0x14, 0x21, 0x1e, 0x2e, 0x96, 0x82, 0xfc, 0xa2, 0x12, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x5e, 0x21, 0x41, 0x2e, 0xce, 0x8c, 0x92, 0x92, 0x82, 0x78, 0xb0, 0x10,
	0x13, 0x48, 0x48, 0xc9, 0x87, 0x8b, 0x0f, 0xa2, 0xc3, 0x2d, 0x33, 0x2f, 0xb3, 0x38, 0x03, 0xa8,
	0x85, 0x97, 0x8b, 0x35, 0xb5, 0xa8, 0x28, 0xbf, 0x08, 0xac, 0x87, 0x53, 0x48, 0x9b, 0x8b, 0x03,
	0x66, 0x09, 0x58, 0x0b, 0xb7, 0x91, 0xa4, 0x1e, 0xc4, 0x15, 0x7a, 0x30, 0x57, 0xe8, 0xb9, 0x40,
	0x15, 0x24, 0xb1, 0x81, 0x85, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x70, 0x04, 0x1d,
	0xcf, 0x00, 0x00, 0x00,
}