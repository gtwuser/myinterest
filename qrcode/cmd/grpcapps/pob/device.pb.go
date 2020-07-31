// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: qrcode/cmd/grpcapps/pob/device.proto

package pob

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceInfo *DeviceInfo `protobuf:"bytes,1,opt,name=deviceInfo,proto3" json:"deviceInfo,omitempty"`
}

func (x *DeviceRequest) Reset() {
	*x = DeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRequest) ProtoMessage() {}

func (x *DeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRequest.ProtoReflect.Descriptor instead.
func (*DeviceRequest) Descriptor() ([]byte, []int) {
	return file_qrcode_cmd_grpcapps_pob_device_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceRequest) GetDeviceInfo() *DeviceInfo {
	if x != nil {
		return x.DeviceInfo
	}
	return nil
}

type DeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *DeviceInfo) Reset() {
	*x = DeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfo) ProtoMessage() {}

func (x *DeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfo.ProtoReflect.Descriptor instead.
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return file_qrcode_cmd_grpcapps_pob_device_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *DeviceInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type DeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DevDetails *DeviceDetails `protobuf:"bytes,1,opt,name=devDetails,proto3" json:"devDetails,omitempty"`
}

func (x *DeviceResponse) Reset() {
	*x = DeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceResponse) ProtoMessage() {}

func (x *DeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceResponse.ProtoReflect.Descriptor instead.
func (*DeviceResponse) Descriptor() ([]byte, []int) {
	return file_qrcode_cmd_grpcapps_pob_device_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceResponse) GetDevDetails() *DeviceDetails {
	if x != nil {
		return x.DevDetails
	}
	return nil
}

type DeviceDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *DeviceDetails) Reset() {
	*x = DeviceDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceDetails) ProtoMessage() {}

func (x *DeviceDetails) ProtoReflect() protoreflect.Message {
	mi := &file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceDetails.ProtoReflect.Descriptor instead.
func (*DeviceDetails) Descriptor() ([]byte, []int) {
	return file_qrcode_cmd_grpcapps_pob_device_proto_rawDescGZIP(), []int{3}
}

func (x *DeviceDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceDetails) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_qrcode_cmd_grpcapps_pob_device_proto protoreflect.FileDescriptor

var file_qrcode_cmd_grpcapps_pob_device_proto_rawDesc = []byte{
	0x0a, 0x24, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6f, 0x62, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0d,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x30, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x0e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x0a, 0x64, 0x65, 0x76, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x37, 0x0a, 0x0d,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x32, 0x6d, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x3a, 0x01, 0x2a, 0x42, 0x19, 0x5a, 0x17, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x63,
	0x6d, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6f, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qrcode_cmd_grpcapps_pob_device_proto_rawDescOnce sync.Once
	file_qrcode_cmd_grpcapps_pob_device_proto_rawDescData = file_qrcode_cmd_grpcapps_pob_device_proto_rawDesc
)

func file_qrcode_cmd_grpcapps_pob_device_proto_rawDescGZIP() []byte {
	file_qrcode_cmd_grpcapps_pob_device_proto_rawDescOnce.Do(func() {
		file_qrcode_cmd_grpcapps_pob_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_qrcode_cmd_grpcapps_pob_device_proto_rawDescData)
	})
	return file_qrcode_cmd_grpcapps_pob_device_proto_rawDescData
}

var file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_qrcode_cmd_grpcapps_pob_device_proto_goTypes = []interface{}{
	(*DeviceRequest)(nil),  // 0: device.DeviceRequest
	(*DeviceInfo)(nil),     // 1: device.DeviceInfo
	(*DeviceResponse)(nil), // 2: device.DeviceResponse
	(*DeviceDetails)(nil),  // 3: device.DeviceDetails
}
var file_qrcode_cmd_grpcapps_pob_device_proto_depIdxs = []int32{
	1, // 0: device.DeviceRequest.deviceInfo:type_name -> device.DeviceInfo
	3, // 1: device.DeviceResponse.devDetails:type_name -> device.DeviceDetails
	0, // 2: device.ManageDevice.ManagedDevice:input_type -> device.DeviceRequest
	2, // 3: device.ManageDevice.ManagedDevice:output_type -> device.DeviceResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_qrcode_cmd_grpcapps_pob_device_proto_init() }
func file_qrcode_cmd_grpcapps_pob_device_proto_init() {
	if File_qrcode_cmd_grpcapps_pob_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_qrcode_cmd_grpcapps_pob_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qrcode_cmd_grpcapps_pob_device_proto_goTypes,
		DependencyIndexes: file_qrcode_cmd_grpcapps_pob_device_proto_depIdxs,
		MessageInfos:      file_qrcode_cmd_grpcapps_pob_device_proto_msgTypes,
	}.Build()
	File_qrcode_cmd_grpcapps_pob_device_proto = out.File
	file_qrcode_cmd_grpcapps_pob_device_proto_rawDesc = nil
	file_qrcode_cmd_grpcapps_pob_device_proto_goTypes = nil
	file_qrcode_cmd_grpcapps_pob_device_proto_depIdxs = nil
}
