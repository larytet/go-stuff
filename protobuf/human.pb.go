// Code generated by protoc-gen-go.
// source: human.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	human.proto

It has these top-level messages:
	Person
	Persons
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
const _ = proto.ProtoPackageIsVersion1

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Person struct {
	Name   string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id     int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email  string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=main.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Persons struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *Persons) Reset()                    { *m = Persons{} }
func (m *Persons) String() string            { return proto.CompactTextString(m) }
func (*Persons) ProtoMessage()               {}
func (*Persons) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Persons) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "main.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "main.Person.PhoneNumber")
	proto.RegisterType((*Persons)(nil), "main.Persons")
	proto.RegisterEnum("main.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

var fileDescriptor0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x28, 0xcd, 0x4d,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0xfa, 0xc2,
	0xc8, 0xc5, 0x16, 0x90, 0x5a, 0x54, 0x9c, 0x9f, 0x27, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b,
	0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48,
	0x30, 0x01, 0x45, 0x58, 0x83, 0x80, 0x2c, 0x21, 0x11, 0x2e, 0xd6, 0x54, 0xa0, 0xbe, 0x1c, 0x09,
	0x66, 0xb0, 0x22, 0x08, 0x47, 0xc8, 0x80, 0x8b, 0xad, 0x20, 0x23, 0x3f, 0x2f, 0xb5, 0x58, 0x82,
	0x45, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x42, 0x0f, 0x64, 0xb6, 0x1e, 0xc4, 0x5c, 0xbd, 0x00, 0x90,
	0x94, 0x5f, 0x69, 0x6e, 0x52, 0x6a, 0x51, 0x10, 0x54, 0x9d, 0x54, 0x20, 0x17, 0x37, 0x92, 0xb0,
	0x90, 0x18, 0x17, 0x5b, 0x1e, 0x98, 0x05, 0xb5, 0x1c, 0xca, 0x13, 0xd2, 0xe2, 0x62, 0x29, 0xa9,
	0x2c, 0x48, 0x05, 0x3b, 0x80, 0xcf, 0x48, 0x0c, 0xd3, 0xd8, 0x10, 0xa0, 0x6c, 0x10, 0x58, 0x8d,
	0x92, 0x36, 0x17, 0x27, 0x5c, 0x48, 0x88, 0x8b, 0x8b, 0xcd, 0xd7, 0xdf, 0xc9, 0xd3, 0xc7, 0x55,
	0x80, 0x41, 0x88, 0x83, 0x8b, 0xc5, 0xc3, 0xdf, 0xd7, 0x55, 0x80, 0x11, 0xc4, 0x0a, 0xf7, 0x0f,
	0xf2, 0x16, 0x60, 0x52, 0xd2, 0xe7, 0x62, 0x87, 0x18, 0x53, 0x2c, 0xa4, 0x02, 0x74, 0x7c, 0x6a,
	0x7e, 0x41, 0x0e, 0xc8, 0xe3, 0x20, 0xc7, 0xf3, 0x20, 0xdb, 0x12, 0x04, 0x95, 0x4b, 0x62, 0x03,
	0x07, 0x9a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xf1, 0xd4, 0x81, 0x43, 0x01, 0x00, 0x00,
}