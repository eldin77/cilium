// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/grpc_credential/v2alpha/file_based_metadata.proto

package v2alpha

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FileBasedMetadataConfig struct {
	// Location or inline data of secret to use for authentication of the Google gRPC connection
	// this secret will be attached to a header of the gRPC connection
	SecretData *core.DataSource `protobuf:"bytes,1,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	// Metadata header key to use for sending the secret data
	// if no header key is set, "authorization" header will be used
	HeaderKey string `protobuf:"bytes,2,opt,name=header_key,json=headerKey,proto3" json:"header_key,omitempty"`
	// Prefix to prepend to the secret in the metadata header
	// if no prefix is set, the default is to use no prefix
	HeaderPrefix         string   `protobuf:"bytes,3,opt,name=header_prefix,json=headerPrefix,proto3" json:"header_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileBasedMetadataConfig) Reset()         { *m = FileBasedMetadataConfig{} }
func (m *FileBasedMetadataConfig) String() string { return proto.CompactTextString(m) }
func (*FileBasedMetadataConfig) ProtoMessage()    {}
func (*FileBasedMetadataConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_based_metadata_ebf64960d046d79e, []int{0}
}
func (m *FileBasedMetadataConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileBasedMetadataConfig.Unmarshal(m, b)
}
func (m *FileBasedMetadataConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileBasedMetadataConfig.Marshal(b, m, deterministic)
}
func (dst *FileBasedMetadataConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileBasedMetadataConfig.Merge(dst, src)
}
func (m *FileBasedMetadataConfig) XXX_Size() int {
	return xxx_messageInfo_FileBasedMetadataConfig.Size(m)
}
func (m *FileBasedMetadataConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FileBasedMetadataConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FileBasedMetadataConfig proto.InternalMessageInfo

func (m *FileBasedMetadataConfig) GetSecretData() *core.DataSource {
	if m != nil {
		return m.SecretData
	}
	return nil
}

func (m *FileBasedMetadataConfig) GetHeaderKey() string {
	if m != nil {
		return m.HeaderKey
	}
	return ""
}

func (m *FileBasedMetadataConfig) GetHeaderPrefix() string {
	if m != nil {
		return m.HeaderPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*FileBasedMetadataConfig)(nil), "envoy.config.grpc_credential.v2alpha.FileBasedMetadataConfig")
}

func init() {
	proto.RegisterFile("envoy/config/grpc_credential/v2alpha/file_based_metadata.proto", fileDescriptor_file_based_metadata_ebf64960d046d79e)
}

var fileDescriptor_file_based_metadata_ebf64960d046d79e = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x3f, 0x4b, 0xc4, 0x30,
	0x1c, 0x86, 0xa9, 0x82, 0xd2, 0x9c, 0x2e, 0x5d, 0x2c, 0xe2, 0xc1, 0xa1, 0x0e, 0x37, 0x25, 0x50,
	0xf7, 0x1b, 0x4e, 0x71, 0x11, 0x41, 0xce, 0xcd, 0x25, 0xfc, 0x2e, 0x7d, 0x7b, 0x17, 0xac, 0x4d,
	0x48, 0x63, 0xb1, 0x1f, 0xc6, 0xef, 0x2a, 0xf9, 0x33, 0xb9, 0x3e, 0xc9, 0xf3, 0xbc, 0x09, 0xdb,
	0x60, 0x98, 0xcc, 0x2c, 0x94, 0x19, 0x3a, 0x7d, 0x10, 0x07, 0x67, 0x95, 0x54, 0x0e, 0x2d, 0x06,
	0xaf, 0xa9, 0x17, 0x53, 0x43, 0xbd, 0x3d, 0x92, 0xe8, 0x74, 0x0f, 0xb9, 0xa7, 0x11, 0xad, 0xfc,
	0x82, 0xa7, 0x96, 0x3c, 0x71, 0xeb, 0x8c, 0x37, 0xd5, 0x7d, 0xf4, 0x79, 0xf2, 0xf9, 0x3f, 0x9f,
	0x67, 0xff, 0xfa, 0x26, 0xad, 0x90, 0xd5, 0x62, 0x6a, 0x84, 0x32, 0x0e, 0x22, 0xd4, 0x52, 0xe3,
	0xf6, 0xb7, 0x60, 0x57, 0xcf, 0xba, 0xc7, 0x36, 0x0c, 0xbc, 0xe6, 0xfe, 0x63, 0x2c, 0x56, 0x1b,
	0xb6, 0x18, 0xa1, 0x1c, 0xbc, 0x0c, 0xb0, 0x2e, 0x56, 0xc5, 0x7a, 0xd1, 0x2c, 0x79, 0x5a, 0x25,
	0xab, 0xf9, 0xd4, 0xf0, 0xd0, 0xe3, 0x4f, 0xe4, 0xe9, 0xdd, 0x7c, 0x3b, 0x85, 0x1d, 0x4b, 0x46,
	0x20, 0xd5, 0x92, 0xb1, 0x23, 0xa8, 0x85, 0x93, 0x9f, 0x98, 0xeb, 0x93, 0x55, 0xb1, 0x2e, 0x77,
	0x65, 0x22, 0x2f, 0x98, 0xab, 0x3b, 0x76, 0x99, 0x8f, 0xad, 0x43, 0xa7, 0x7f, 0xea, 0xd3, 0x78,
	0xe3, 0x22, 0xc1, 0xb7, 0xc8, 0xb6, 0xe5, 0xc7, 0x79, 0xfe, 0xc8, 0xfe, 0x2c, 0xbe, 0xf8, 0xe1,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xa5, 0x3b, 0xd0, 0x37, 0x01, 0x00, 0x00,
}
