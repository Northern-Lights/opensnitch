// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opensnitch/rules/rules.proto

package rules // import "github.com/evilsocket/opensnitch/rules"

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

type Action int32

const (
	Action_ACTION_UNKNOWN Action = 0
	Action_ALLOW          Action = 1
	Action_DENY           Action = 2
)

var Action_name = map[int32]string{
	0: "ACTION_UNKNOWN",
	1: "ALLOW",
	2: "DENY",
}
var Action_value = map[string]int32{
	"ACTION_UNKNOWN": 0,
	"ALLOW":          1,
	"DENY":           2,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}
func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rules_d1f7ce15b89ebdec, []int{0}
}

// Duration tells how long a rule should apply
type Duration int32

const (
	Duration_DURATION_UNKNOWN Duration = 0
	Duration_ONCE             Duration = 1
	Duration_FIREWALL_SESSION Duration = 2
	Duration_ALWAYS           Duration = 3
)

var Duration_name = map[int32]string{
	0: "DURATION_UNKNOWN",
	1: "ONCE",
	2: "FIREWALL_SESSION",
	3: "ALWAYS",
}
var Duration_value = map[string]int32{
	"DURATION_UNKNOWN": 0,
	"ONCE":             1,
	"FIREWALL_SESSION": 2,
	"ALWAYS":           3,
}

func (x Duration) String() string {
	return proto.EnumName(Duration_name, int32(x))
}
func (Duration) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rules_d1f7ce15b89ebdec, []int{1}
}

// Operation denotes the target to operate upon and how to operate on it
type Operation int32

const (
	Operation_OPERATION_UNKNOWN Operation = 0
	Operation_TRUE              Operation = 9
	Operation_FALSE             Operation = 10
	Operation_AND               Operation = 1
	Operation_OR                Operation = 2
	Operation_NOT               Operation = 3
	Operation_DST_IP            Operation = 4
	Operation_DST_HOST          Operation = 5
	Operation_DST_PORT          Operation = 6
	Operation_PROC_PATH         Operation = 7
	Operation_PID               Operation = 8
)

var Operation_name = map[int32]string{
	0:  "OPERATION_UNKNOWN",
	9:  "TRUE",
	10: "FALSE",
	1:  "AND",
	2:  "OR",
	3:  "NOT",
	4:  "DST_IP",
	5:  "DST_HOST",
	6:  "DST_PORT",
	7:  "PROC_PATH",
	8:  "PID",
}
var Operation_value = map[string]int32{
	"OPERATION_UNKNOWN": 0,
	"TRUE":              9,
	"FALSE":             10,
	"AND":               1,
	"OR":                2,
	"NOT":               3,
	"DST_IP":            4,
	"DST_HOST":          5,
	"DST_PORT":          6,
	"PROC_PATH":         7,
	"PID":               8,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}
func (Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rules_d1f7ce15b89ebdec, []int{2}
}

type Rule struct {
	Name                 string      `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Action               Action      `protobuf:"varint,1,opt,name=action,proto3,enum=opensnitch.rules.Action" json:"action,omitempty"`
	Duration             Duration    `protobuf:"varint,2,opt,name=duration,proto3,enum=opensnitch.rules.Duration" json:"duration,omitempty"`
	Condition            *Expression `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_rules_d1f7ce15b89ebdec, []int{0}
}
func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (dst *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(dst, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rule) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_ACTION_UNKNOWN
}

func (m *Rule) GetDuration() Duration {
	if m != nil {
		return m.Duration
	}
	return Duration_DURATION_UNKNOWN
}

func (m *Rule) GetCondition() *Expression {
	if m != nil {
		return m.Condition
	}
	return nil
}

// Expression is used to decide whether a connection matches a condition.
// It can be made up of 1 other expression (e.g. for NOT) or 2 other expressions
// (e.g. for AND or OR,) or it can be a "leaf" expression specifying a primitive
// type upon which to operate. The language-specific implementation should
// select the correct field based on the operation
type Expression struct {
	Operation            Operation   `protobuf:"varint,1,opt,name=operation,proto3,enum=opensnitch.rules.Operation" json:"operation,omitempty"`
	Left                 *Expression `protobuf:"bytes,14,opt,name=left,proto3" json:"left,omitempty"`
	Right                *Expression `protobuf:"bytes,15,opt,name=right,proto3" json:"right,omitempty"`
	Strings              []string    `protobuf:"bytes,2,rep,name=strings,proto3" json:"strings,omitempty"`
	Uint32S              []uint32    `protobuf:"varint,3,rep,packed,name=uint32s,proto3" json:"uint32s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Expression) Reset()         { *m = Expression{} }
func (m *Expression) String() string { return proto.CompactTextString(m) }
func (*Expression) ProtoMessage()    {}
func (*Expression) Descriptor() ([]byte, []int) {
	return fileDescriptor_rules_d1f7ce15b89ebdec, []int{1}
}
func (m *Expression) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expression.Unmarshal(m, b)
}
func (m *Expression) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expression.Marshal(b, m, deterministic)
}
func (dst *Expression) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expression.Merge(dst, src)
}
func (m *Expression) XXX_Size() int {
	return xxx_messageInfo_Expression.Size(m)
}
func (m *Expression) XXX_DiscardUnknown() {
	xxx_messageInfo_Expression.DiscardUnknown(m)
}

var xxx_messageInfo_Expression proto.InternalMessageInfo

func (m *Expression) GetOperation() Operation {
	if m != nil {
		return m.Operation
	}
	return Operation_OPERATION_UNKNOWN
}

func (m *Expression) GetLeft() *Expression {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *Expression) GetRight() *Expression {
	if m != nil {
		return m.Right
	}
	return nil
}

func (m *Expression) GetStrings() []string {
	if m != nil {
		return m.Strings
	}
	return nil
}

func (m *Expression) GetUint32S() []uint32 {
	if m != nil {
		return m.Uint32S
	}
	return nil
}

func init() {
	proto.RegisterType((*Rule)(nil), "opensnitch.rules.Rule")
	proto.RegisterType((*Expression)(nil), "opensnitch.rules.Expression")
	proto.RegisterEnum("opensnitch.rules.Action", Action_name, Action_value)
	proto.RegisterEnum("opensnitch.rules.Duration", Duration_name, Duration_value)
	proto.RegisterEnum("opensnitch.rules.Operation", Operation_name, Operation_value)
}

func init() { proto.RegisterFile("opensnitch/rules/rules.proto", fileDescriptor_rules_d1f7ce15b89ebdec) }

var fileDescriptor_rules_d1f7ce15b89ebdec = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x8a, 0x9b, 0x40,
	0x14, 0x87, 0x3b, 0x6a, 0x8c, 0x9e, 0x76, 0xd3, 0xe9, 0xa1, 0x05, 0x69, 0xf7, 0x22, 0xec, 0x45,
	0x91, 0x5c, 0x98, 0x6d, 0x16, 0x0a, 0xed, 0xdd, 0x34, 0xba, 0xac, 0x54, 0x1c, 0x19, 0x0d, 0x61,
	0x7b, 0x13, 0xb2, 0xae, 0x4d, 0xa4, 0x59, 0x0d, 0xfe, 0x29, 0x7d, 0x8d, 0xbe, 0x51, 0x5f, 0xa5,
	0x6f, 0x52, 0x34, 0x71, 0x03, 0x49, 0x2f, 0xf6, 0x46, 0xe6, 0xcc, 0xef, 0xfb, 0x0e, 0x87, 0x83,
	0x03, 0xe7, 0xf9, 0x36, 0xc9, 0xca, 0x2c, 0xad, 0xe2, 0xf5, 0xb8, 0xa8, 0x37, 0x49, 0xb9, 0xfb,
	0x5a, 0xdb, 0x22, 0xaf, 0x72, 0xa4, 0x87, 0xd4, 0x6a, 0xef, 0x2f, 0xfe, 0x10, 0x50, 0x44, 0xbd,
	0x49, 0x10, 0x41, 0xc9, 0x96, 0x0f, 0x89, 0xa1, 0x0c, 0x89, 0xa9, 0x8b, 0xf6, 0x8c, 0x97, 0xa0,
	0x2e, 0xe3, 0x2a, 0xcd, 0x33, 0x83, 0x0c, 0x89, 0x39, 0x98, 0x18, 0xd6, 0xb1, 0x6f, 0xb1, 0x36,
	0x17, 0x7b, 0x0e, 0x3f, 0x82, 0x76, 0x5f, 0x17, 0xcb, 0xd6, 0x91, 0x5a, 0xe7, 0xed, 0xa9, 0x63,
	0xef, 0x09, 0xf1, 0xc8, 0xe2, 0x67, 0xd0, 0xe3, 0x3c, 0xbb, 0x4f, 0x5b, 0x51, 0x1e, 0x12, 0xf3,
	0xf9, 0xe4, 0xfc, 0x54, 0x74, 0x7e, 0x6d, 0x8b, 0xa4, 0x2c, 0x1b, 0xf5, 0x80, 0x5f, 0xfc, 0x25,
	0x00, 0x87, 0x04, 0x3f, 0x81, 0x9e, 0x6f, 0x93, 0xfd, 0x0c, 0xbb, 0xb9, 0xdf, 0x9d, 0xb6, 0xe2,
	0x1d, 0x22, 0x0e, 0x34, 0x5e, 0x82, 0xb2, 0x49, 0xbe, 0x57, 0xc6, 0xe0, 0x09, 0x03, 0xb4, 0x24,
	0x4e, 0xa0, 0x57, 0xa4, 0xab, 0x75, 0x65, 0xbc, 0x7c, 0x82, 0xb2, 0x43, 0xd1, 0x80, 0x7e, 0x59,
	0x15, 0x69, 0xb6, 0x2a, 0x0d, 0x69, 0x28, 0x9b, 0xba, 0xe8, 0xca, 0x26, 0xa9, 0xd3, 0xac, 0xba,
	0x9a, 0x94, 0x86, 0x3c, 0x94, 0xcd, 0x33, 0xd1, 0x95, 0xa3, 0x0f, 0xa0, 0xee, 0x36, 0x8d, 0x08,
	0x03, 0x36, 0x8d, 0x5c, 0xee, 0x2f, 0x66, 0xfe, 0x57, 0x9f, 0xcf, 0x7d, 0xfa, 0x0c, 0x75, 0xe8,
	0x31, 0xcf, 0xe3, 0x73, 0x4a, 0x50, 0x03, 0xc5, 0x76, 0xfc, 0x5b, 0x2a, 0x8d, 0x3c, 0xd0, 0xba,
	0x45, 0xe3, 0x6b, 0xa0, 0xf6, 0x4c, 0xb0, 0x23, 0x4d, 0x03, 0x85, 0xfb, 0x53, 0x87, 0x92, 0x26,
	0xbf, 0x76, 0x85, 0x33, 0x67, 0x9e, 0xb7, 0x08, 0x9d, 0x30, 0x74, 0xb9, 0x4f, 0x25, 0x04, 0x50,
	0x99, 0x37, 0x67, 0xb7, 0x21, 0x95, 0x47, 0xbf, 0x09, 0xe8, 0x8f, 0x3b, 0xc3, 0x37, 0xf0, 0x8a,
	0x07, 0xce, 0xff, 0x1a, 0x46, 0x62, 0xe6, 0x50, 0xbd, 0x99, 0xe8, 0x9a, 0x79, 0xa1, 0x43, 0x01,
	0xfb, 0x20, 0x33, 0xdf, 0xa6, 0x04, 0x55, 0x90, 0xb8, 0xa0, 0x52, 0x73, 0xe1, 0xf3, 0x88, 0xca,
	0x4d, 0x7f, 0x3b, 0x8c, 0x16, 0x6e, 0x40, 0x15, 0x7c, 0x01, 0x5a, 0x73, 0xbe, 0xe1, 0x61, 0x44,
	0x7b, 0x5d, 0x15, 0x70, 0x11, 0x51, 0x15, 0xcf, 0x40, 0x0f, 0x04, 0x9f, 0x2e, 0x02, 0x16, 0xdd,
	0xd0, 0x7e, 0xe3, 0x07, 0xae, 0x4d, 0xb5, 0x2f, 0xe6, 0xb7, 0xf7, 0xab, 0xb4, 0x5a, 0xd7, 0x77,
	0x56, 0x9c, 0x3f, 0x8c, 0x93, 0x9f, 0xe9, 0xa6, 0xcc, 0xe3, 0x1f, 0x49, 0x35, 0x3e, 0x7e, 0x03,
	0x77, 0x6a, 0xfb, 0xfb, 0x5f, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xad, 0x21, 0xb8, 0x17, 0x1e,
	0x03, 0x00, 0x00,
}
