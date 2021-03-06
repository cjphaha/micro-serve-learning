// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ProdService.proto

package Services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProdModel struct {
	ProdID               int32    `protobuf:"varint,1,opt,name=ProdID,proto3" json:"ProdID,omitempty"`
	ProdName             string   `protobuf:"bytes,2,opt,name=ProdName,proto3" json:"ProdName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdModel) Reset()         { *m = ProdModel{} }
func (m *ProdModel) String() string { return proto.CompactTextString(m) }
func (*ProdModel) ProtoMessage()    {}
func (*ProdModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_50db98fd6a3e2ab5, []int{0}
}

func (m *ProdModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdModel.Unmarshal(m, b)
}
func (m *ProdModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdModel.Marshal(b, m, deterministic)
}
func (m *ProdModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdModel.Merge(m, src)
}
func (m *ProdModel) XXX_Size() int {
	return xxx_messageInfo_ProdModel.Size(m)
}
func (m *ProdModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdModel.DiscardUnknown(m)
}

var xxx_messageInfo_ProdModel proto.InternalMessageInfo

func (m *ProdModel) GetProdID() int32 {
	if m != nil {
		return m.ProdID
	}
	return 0
}

func (m *ProdModel) GetProdName() string {
	if m != nil {
		return m.ProdName
	}
	return ""
}

type ProdsRequest struct {
	Size                 int32    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	ProdId               int32    `protobuf:"varint,2,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdsRequest) Reset()         { *m = ProdsRequest{} }
func (m *ProdsRequest) String() string { return proto.CompactTextString(m) }
func (*ProdsRequest) ProtoMessage()    {}
func (*ProdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_50db98fd6a3e2ab5, []int{1}
}

func (m *ProdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdsRequest.Unmarshal(m, b)
}
func (m *ProdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdsRequest.Marshal(b, m, deterministic)
}
func (m *ProdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdsRequest.Merge(m, src)
}
func (m *ProdsRequest) XXX_Size() int {
	return xxx_messageInfo_ProdsRequest.Size(m)
}
func (m *ProdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProdsRequest proto.InternalMessageInfo

func (m *ProdsRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ProdsRequest) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

type ProdListResponse struct {
	Data                 []*ProdModel `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProdListResponse) Reset()         { *m = ProdListResponse{} }
func (m *ProdListResponse) String() string { return proto.CompactTextString(m) }
func (*ProdListResponse) ProtoMessage()    {}
func (*ProdListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50db98fd6a3e2ab5, []int{2}
}

func (m *ProdListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdListResponse.Unmarshal(m, b)
}
func (m *ProdListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdListResponse.Marshal(b, m, deterministic)
}
func (m *ProdListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdListResponse.Merge(m, src)
}
func (m *ProdListResponse) XXX_Size() int {
	return xxx_messageInfo_ProdListResponse.Size(m)
}
func (m *ProdListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProdListResponse proto.InternalMessageInfo

func (m *ProdListResponse) GetData() []*ProdModel {
	if m != nil {
		return m.Data
	}
	return nil
}

type ProdDetailResponse struct {
	Data                 *ProdModel `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProdDetailResponse) Reset()         { *m = ProdDetailResponse{} }
func (m *ProdDetailResponse) String() string { return proto.CompactTextString(m) }
func (*ProdDetailResponse) ProtoMessage()    {}
func (*ProdDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50db98fd6a3e2ab5, []int{3}
}

func (m *ProdDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdDetailResponse.Unmarshal(m, b)
}
func (m *ProdDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdDetailResponse.Marshal(b, m, deterministic)
}
func (m *ProdDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdDetailResponse.Merge(m, src)
}
func (m *ProdDetailResponse) XXX_Size() int {
	return xxx_messageInfo_ProdDetailResponse.Size(m)
}
func (m *ProdDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProdDetailResponse proto.InternalMessageInfo

func (m *ProdDetailResponse) GetData() *ProdModel {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ProdModel)(nil), "Services.ProdModel")
	proto.RegisterType((*ProdsRequest)(nil), "Services.ProdsRequest")
	proto.RegisterType((*ProdListResponse)(nil), "Services.ProdListResponse")
	proto.RegisterType((*ProdDetailResponse)(nil), "Services.ProdDetailResponse")
}

func init() { proto.RegisterFile("ProdService.proto", fileDescriptor_50db98fd6a3e2ab5) }

var fileDescriptor_50db98fd6a3e2ab5 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0x24, 0xda, 0xae, 0xed, 0x6b, 0x11, 0x7d, 0x42, 0x5d, 0x16, 0x0f, 0xcb, 0x5e, 0xdc, 0xd3,
	0x1e, 0xea, 0xb1, 0x88, 0x20, 0x45, 0x29, 0xa8, 0x48, 0xfc, 0x00, 0x89, 0xe6, 0x1d, 0x02, 0xd5,
	0xac, 0x49, 0xf4, 0xe0, 0x9f, 0xf8, 0xb7, 0xf2, 0xe2, 0x6e, 0x31, 0xa0, 0xbd, 0xcd, 0x4c, 0x32,
	0xc3, 0x4c, 0x02, 0x87, 0xf7, 0xce, 0xea, 0x07, 0x72, 0x1f, 0xe6, 0x99, 0x9a, 0xd6, 0xd9, 0x60,
	0x71, 0xd4, 0x51, 0x5f, 0x5d, 0xc0, 0x98, 0x8f, 0x6f, 0xad, 0xa6, 0x35, 0xce, 0x20, 0x63, 0xb2,
	0x5a, 0xe6, 0xa2, 0x14, 0xf5, 0x50, 0x76, 0x0c, 0x0b, 0x18, 0x31, 0xba, 0x53, 0x2f, 0x94, 0xef,
	0x94, 0xa2, 0x1e, 0xcb, 0x0d, 0xaf, 0x16, 0x30, 0x65, 0xec, 0x25, 0xbd, 0xbd, 0x93, 0x0f, 0x88,
	0x30, 0xf0, 0xe6, 0x93, 0xba, 0x84, 0x88, 0xf1, 0x18, 0xf6, 0x5a, 0x67, 0xf5, 0xa3, 0xd1, 0xd1,
	0x3e, 0x94, 0x19, 0xd3, 0x95, 0xae, 0x16, 0x70, 0xc0, 0xe6, 0x1b, 0xe3, 0x83, 0x24, 0xdf, 0xda,
	0x57, 0x4f, 0x78, 0x0a, 0x03, 0xad, 0x82, 0xca, 0x45, 0xb9, 0x5b, 0x4f, 0xe6, 0x47, 0x4d, 0x5f,
	0xb5, 0xd9, 0xf4, 0x94, 0xf1, 0x42, 0x75, 0x0e, 0xc8, 0xd2, 0x92, 0x82, 0x32, 0xeb, 0x3f, 0xec,
	0x62, 0xab, 0x7d, 0xfe, 0x25, 0x60, 0xf2, 0xeb, 0x65, 0xf0, 0x12, 0xa6, 0xd7, 0x14, 0xe2, 0x16,
	0xee, 0x83, 0xb3, 0xd4, 0xda, 0x0f, 0x2c, 0x8a, 0x54, 0x4f, 0xba, 0x5f, 0xc1, 0x7e, 0x9f, 0xf1,
	0x53, 0xeb, 0xdf, 0x94, 0x93, 0x54, 0x4f, 0x47, 0x3c, 0x65, 0xf1, 0x9b, 0xce, 0xbe, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xef, 0x65, 0xd7, 0xa4, 0xbb, 0x01, 0x00, 0x00,
}
