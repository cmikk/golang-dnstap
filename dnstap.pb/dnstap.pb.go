// Code generated by protoc-gen-go.
// source: dnstap.pb/dnstap.proto
// DO NOT EDIT!

package dnstap

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type SocketFamily int32

const (
	SocketFamily_INET  SocketFamily = 1
	SocketFamily_INET6 SocketFamily = 2
)

var SocketFamily_name = map[int32]string{
	1: "INET",
	2: "INET6",
}
var SocketFamily_value = map[string]int32{
	"INET":  1,
	"INET6": 2,
}

func (x SocketFamily) Enum() *SocketFamily {
	p := new(SocketFamily)
	*p = x
	return p
}
func (x SocketFamily) String() string {
	return proto.EnumName(SocketFamily_name, int32(x))
}
func (x SocketFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *SocketFamily) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SocketFamily_value, data, "SocketFamily")
	if err != nil {
		return err
	}
	*x = SocketFamily(value)
	return nil
}

type SocketProtocol int32

const (
	SocketProtocol_UDP SocketProtocol = 1
	SocketProtocol_TCP SocketProtocol = 2
)

var SocketProtocol_name = map[int32]string{
	1: "UDP",
	2: "TCP",
}
var SocketProtocol_value = map[string]int32{
	"UDP": 1,
	"TCP": 2,
}

func (x SocketProtocol) Enum() *SocketProtocol {
	p := new(SocketProtocol)
	*p = x
	return p
}
func (x SocketProtocol) String() string {
	return proto.EnumName(SocketProtocol_name, int32(x))
}
func (x SocketProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *SocketProtocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SocketProtocol_value, data, "SocketProtocol")
	if err != nil {
		return err
	}
	*x = SocketProtocol(value)
	return nil
}

type Dnstap_Type int32

const (
	Dnstap_MESSAGE Dnstap_Type = 1
)

var Dnstap_Type_name = map[int32]string{
	1: "MESSAGE",
}
var Dnstap_Type_value = map[string]int32{
	"MESSAGE": 1,
}

func (x Dnstap_Type) Enum() *Dnstap_Type {
	p := new(Dnstap_Type)
	*p = x
	return p
}
func (x Dnstap_Type) String() string {
	return proto.EnumName(Dnstap_Type_name, int32(x))
}
func (x Dnstap_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Dnstap_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Dnstap_Type_value, data, "Dnstap_Type")
	if err != nil {
		return err
	}
	*x = Dnstap_Type(value)
	return nil
}

type Message_Type int32

const (
	Message_AUTH_QUERY         Message_Type = 1
	Message_AUTH_RESPONSE      Message_Type = 2
	Message_RESOLVER_QUERY     Message_Type = 3
	Message_RESOLVER_RESPONSE  Message_Type = 4
	Message_CLIENT_QUERY       Message_Type = 5
	Message_CLIENT_RESPONSE    Message_Type = 6
	Message_FORWARDER_QUERY    Message_Type = 7
	Message_FORWARDER_RESPONSE Message_Type = 8
)

var Message_Type_name = map[int32]string{
	1: "AUTH_QUERY",
	2: "AUTH_RESPONSE",
	3: "RESOLVER_QUERY",
	4: "RESOLVER_RESPONSE",
	5: "CLIENT_QUERY",
	6: "CLIENT_RESPONSE",
	7: "FORWARDER_QUERY",
	8: "FORWARDER_RESPONSE",
}
var Message_Type_value = map[string]int32{
	"AUTH_QUERY":         1,
	"AUTH_RESPONSE":      2,
	"RESOLVER_QUERY":     3,
	"RESOLVER_RESPONSE":  4,
	"CLIENT_QUERY":       5,
	"CLIENT_RESPONSE":    6,
	"FORWARDER_QUERY":    7,
	"FORWARDER_RESPONSE": 8,
}

func (x Message_Type) Enum() *Message_Type {
	p := new(Message_Type)
	*p = x
	return p
}
func (x Message_Type) String() string {
	return proto.EnumName(Message_Type_name, int32(x))
}
func (x Message_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Message_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Message_Type_value, data, "Message_Type")
	if err != nil {
		return err
	}
	*x = Message_Type(value)
	return nil
}

type Dnstap struct {
	Identity         []byte       `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	Version          []byte       `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Type             *Dnstap_Type `protobuf:"varint,15,req,name=type,enum=dnstap.Dnstap_Type" json:"type,omitempty"`
	Message          *Message     `protobuf:"bytes,14,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Dnstap) Reset()         { *m = Dnstap{} }
func (m *Dnstap) String() string { return proto.CompactTextString(m) }
func (*Dnstap) ProtoMessage()    {}

func (m *Dnstap) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Dnstap) GetVersion() []byte {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Dnstap) GetType() Dnstap_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Dnstap) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type Message struct {
	Type             *Message_Type   `protobuf:"varint,1,req,name=type,enum=dnstap.Message_Type" json:"type,omitempty"`
	SocketFamily     *SocketFamily   `protobuf:"varint,2,opt,name=socket_family,enum=dnstap.SocketFamily" json:"socket_family,omitempty"`
	SocketProtocol   *SocketProtocol `protobuf:"varint,3,opt,name=socket_protocol,enum=dnstap.SocketProtocol" json:"socket_protocol,omitempty"`
	QueryAddress     []byte          `protobuf:"bytes,4,opt,name=query_address" json:"query_address,omitempty"`
	ResponseAddress  []byte          `protobuf:"bytes,5,opt,name=response_address" json:"response_address,omitempty"`
	QueryPort        *uint32         `protobuf:"varint,6,opt,name=query_port" json:"query_port,omitempty"`
	ResponsePort     *uint32         `protobuf:"varint,7,opt,name=response_port" json:"response_port,omitempty"`
	MessageId        *uint32         `protobuf:"varint,8,opt,name=message_id" json:"message_id,omitempty"`
	QueryName        []byte          `protobuf:"bytes,9,opt,name=query_name" json:"query_name,omitempty"`
	QueryType        *uint32         `protobuf:"varint,10,opt,name=query_type" json:"query_type,omitempty"`
	QueryClass       *uint32         `protobuf:"varint,11,opt,name=query_class" json:"query_class,omitempty"`
	QueryTimeSec     *uint64         `protobuf:"varint,12,opt,name=query_time_sec" json:"query_time_sec,omitempty"`
	QueryTimeNsec    *uint32         `protobuf:"fixed32,13,opt,name=query_time_nsec" json:"query_time_nsec,omitempty"`
	QueryMessage     []byte          `protobuf:"bytes,14,opt,name=query_message" json:"query_message,omitempty"`
	QueryZone        []byte          `protobuf:"bytes,15,opt,name=query_zone" json:"query_zone,omitempty"`
	ResponseTimeSec  *uint64         `protobuf:"varint,16,opt,name=response_time_sec" json:"response_time_sec,omitempty"`
	ResponseTimeNsec *uint32         `protobuf:"fixed32,17,opt,name=response_time_nsec" json:"response_time_nsec,omitempty"`
	ResponseMessage  []byte          `protobuf:"bytes,18,opt,name=response_message" json:"response_message,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

func (m *Message) GetType() Message_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Message) GetSocketFamily() SocketFamily {
	if m != nil && m.SocketFamily != nil {
		return *m.SocketFamily
	}
	return 0
}

func (m *Message) GetSocketProtocol() SocketProtocol {
	if m != nil && m.SocketProtocol != nil {
		return *m.SocketProtocol
	}
	return 0
}

func (m *Message) GetQueryAddress() []byte {
	if m != nil {
		return m.QueryAddress
	}
	return nil
}

func (m *Message) GetResponseAddress() []byte {
	if m != nil {
		return m.ResponseAddress
	}
	return nil
}

func (m *Message) GetQueryPort() uint32 {
	if m != nil && m.QueryPort != nil {
		return *m.QueryPort
	}
	return 0
}

func (m *Message) GetResponsePort() uint32 {
	if m != nil && m.ResponsePort != nil {
		return *m.ResponsePort
	}
	return 0
}

func (m *Message) GetMessageId() uint32 {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return 0
}

func (m *Message) GetQueryName() []byte {
	if m != nil {
		return m.QueryName
	}
	return nil
}

func (m *Message) GetQueryType() uint32 {
	if m != nil && m.QueryType != nil {
		return *m.QueryType
	}
	return 0
}

func (m *Message) GetQueryClass() uint32 {
	if m != nil && m.QueryClass != nil {
		return *m.QueryClass
	}
	return 0
}

func (m *Message) GetQueryTimeSec() uint64 {
	if m != nil && m.QueryTimeSec != nil {
		return *m.QueryTimeSec
	}
	return 0
}

func (m *Message) GetQueryTimeNsec() uint32 {
	if m != nil && m.QueryTimeNsec != nil {
		return *m.QueryTimeNsec
	}
	return 0
}

func (m *Message) GetQueryMessage() []byte {
	if m != nil {
		return m.QueryMessage
	}
	return nil
}

func (m *Message) GetQueryZone() []byte {
	if m != nil {
		return m.QueryZone
	}
	return nil
}

func (m *Message) GetResponseTimeSec() uint64 {
	if m != nil && m.ResponseTimeSec != nil {
		return *m.ResponseTimeSec
	}
	return 0
}

func (m *Message) GetResponseTimeNsec() uint32 {
	if m != nil && m.ResponseTimeNsec != nil {
		return *m.ResponseTimeNsec
	}
	return 0
}

func (m *Message) GetResponseMessage() []byte {
	if m != nil {
		return m.ResponseMessage
	}
	return nil
}

func init() {
	proto.RegisterEnum("dnstap.SocketFamily", SocketFamily_name, SocketFamily_value)
	proto.RegisterEnum("dnstap.SocketProtocol", SocketProtocol_name, SocketProtocol_value)
	proto.RegisterEnum("dnstap.Dnstap_Type", Dnstap_Type_name, Dnstap_Type_value)
	proto.RegisterEnum("dnstap.Message_Type", Message_Type_name, Message_Type_value)
}
