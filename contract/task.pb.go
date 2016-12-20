// Code generated by protoc-gen-go.
// source: task.proto
// DO NOT EDIT!

/*
Package contract is a generated protocol buffer package.

It is generated from these files:
	task.proto

It has these top-level messages:
	Task
	Response
	User
*/
package contract

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

type Task struct {
	Task             *string `protobuf:"bytes,1,opt,name=task" json:"task,omitempty"`
	Id               *int32  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetTask() string {
	if m != nil && m.Task != nil {
		return *m.Task
	}
	return ""
}

func (m *Task) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type Response struct {
	Response         []byte `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

type User struct {
	UserName         *string `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	EmailId          *string `protobuf:"bytes,2,opt,name=emailId" json:"emailId,omitempty"`
	Password         *string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *User) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *User) GetEmailId() string {
	if m != nil && m.EmailId != nil {
		return *m.EmailId
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "contract.Task")
	proto.RegisterType((*Response)(nil), "contract.Response")
	proto.RegisterType((*User)(nil), "contract.User")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2c, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e,
	0x51, 0x52, 0xe0, 0x62, 0x09, 0x49, 0x2c, 0xce, 0x16, 0xe2, 0xe1, 0x62, 0x01, 0xc9, 0x4b, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x0a, 0x71, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xb0,
	0x2a, 0xc9, 0x70, 0x71, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x09, 0x70, 0x71,
	0x14, 0x41, 0xd9, 0x60, 0x95, 0x3c, 0x4a, 0xd6, 0x5c, 0x2c, 0xa1, 0xc5, 0xa9, 0x45, 0x20, 0x99,
	0xd2, 0xe2, 0xd4, 0x22, 0xbf, 0xc4, 0xdc, 0x54, 0xa8, 0x19, 0xfc, 0x5c, 0xec, 0xa9, 0xb9, 0x89,
	0x99, 0x39, 0x9e, 0x10, 0x83, 0x38, 0x41, 0x4a, 0x0a, 0x12, 0x8b, 0x8b, 0xcb, 0xf3, 0x8b, 0x52,
	0x24, 0x98, 0x41, 0x22, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x4b, 0xd4, 0x66, 0x93, 0x00,
	0x00, 0x00,
}
