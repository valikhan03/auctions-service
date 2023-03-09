// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protobuf/auctions_run_service.proto

package auctions_run_service

import (
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

type SuggestPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    int64  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	AuctionID string `protobuf:"bytes,2,opt,name=AuctionID,proto3" json:"AuctionID,omitempty"`
	LotID     string `protobuf:"bytes,3,opt,name=LotID,proto3" json:"LotID,omitempty"`
	NewPrice  int64  `protobuf:"varint,4,opt,name=NewPrice,proto3" json:"NewPrice,omitempty"`
}

func (x *SuggestPriceRequest) Reset() {
	*x = SuggestPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_auctions_run_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestPriceRequest) ProtoMessage() {}

func (x *SuggestPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_auctions_run_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestPriceRequest.ProtoReflect.Descriptor instead.
func (*SuggestPriceRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_auctions_run_service_proto_rawDescGZIP(), []int{0}
}

func (x *SuggestPriceRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *SuggestPriceRequest) GetAuctionID() string {
	if x != nil {
		return x.AuctionID
	}
	return ""
}

func (x *SuggestPriceRequest) GetLotID() string {
	if x != nil {
		return x.LotID
	}
	return ""
}

func (x *SuggestPriceRequest) GetNewPrice() int64 {
	if x != nil {
		return x.NewPrice
	}
	return 0
}

type SuggestPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *SuggestPriceResponse) Reset() {
	*x = SuggestPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_auctions_run_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestPriceResponse) ProtoMessage() {}

func (x *SuggestPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_auctions_run_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestPriceResponse.ProtoReflect.Descriptor instead.
func (*SuggestPriceResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_auctions_run_service_proto_rawDescGZIP(), []int{1}
}

func (x *SuggestPriceResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SuggestPriceResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetCurrentPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionID string `protobuf:"bytes,1,opt,name=AuctionID,proto3" json:"AuctionID,omitempty"`
	LotID     string `protobuf:"bytes,2,opt,name=LotID,proto3" json:"LotID,omitempty"`
	UserID    string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetCurrentPriceRequest) Reset() {
	*x = GetCurrentPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_auctions_run_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentPriceRequest) ProtoMessage() {}

func (x *GetCurrentPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_auctions_run_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentPriceRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentPriceRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_auctions_run_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCurrentPriceRequest) GetAuctionID() string {
	if x != nil {
		return x.AuctionID
	}
	return ""
}

func (x *GetCurrentPriceRequest) GetLotID() string {
	if x != nil {
		return x.LotID
	}
	return ""
}

func (x *GetCurrentPriceRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetCurrentPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPrice map[string]string `protobuf:"bytes,1,rep,name=UserPrice,proto3" json:"UserPrice,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetCurrentPriceResponse) Reset() {
	*x = GetCurrentPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_auctions_run_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentPriceResponse) ProtoMessage() {}

func (x *GetCurrentPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_auctions_run_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentPriceResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentPriceResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_auctions_run_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCurrentPriceResponse) GetUserPrice() map[string]string {
	if x != nil {
		return x.UserPrice
	}
	return nil
}

var File_protobuf_auctions_run_service_proto protoreflect.FileDescriptor

var file_protobuf_auctions_run_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22,
	0x7d, 0x0a, 0x13, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x6f, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x74,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x44,
	0x0a, 0x14, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x64, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x6f, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x74,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0xa7, 0x01, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x3c, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0xc1, 0x01, 0x0a, 0x12, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x1a, 0x5a, 0x18, 0x2f, 0x70, 0x62, 0x2f,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_auctions_run_service_proto_rawDescOnce sync.Once
	file_protobuf_auctions_run_service_proto_rawDescData = file_protobuf_auctions_run_service_proto_rawDesc
)

func file_protobuf_auctions_run_service_proto_rawDescGZIP() []byte {
	file_protobuf_auctions_run_service_proto_rawDescOnce.Do(func() {
		file_protobuf_auctions_run_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_auctions_run_service_proto_rawDescData)
	})
	return file_protobuf_auctions_run_service_proto_rawDescData
}

var file_protobuf_auctions_run_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protobuf_auctions_run_service_proto_goTypes = []interface{}{
	(*SuggestPriceRequest)(nil),     // 0: protobuf.SuggestPriceRequest
	(*SuggestPriceResponse)(nil),    // 1: protobuf.SuggestPriceResponse
	(*GetCurrentPriceRequest)(nil),  // 2: protobuf.GetCurrentPriceRequest
	(*GetCurrentPriceResponse)(nil), // 3: protobuf.GetCurrentPriceResponse
	nil,                             // 4: protobuf.GetCurrentPriceResponse.UserPriceEntry
}
var file_protobuf_auctions_run_service_proto_depIdxs = []int32{
	4, // 0: protobuf.GetCurrentPriceResponse.UserPrice:type_name -> protobuf.GetCurrentPriceResponse.UserPriceEntry
	0, // 1: protobuf.AuctionsRunService.SuggestPrice:input_type -> protobuf.SuggestPriceRequest
	2, // 2: protobuf.AuctionsRunService.GetCurrentPrice:input_type -> protobuf.GetCurrentPriceRequest
	1, // 3: protobuf.AuctionsRunService.SuggestPrice:output_type -> protobuf.SuggestPriceResponse
	3, // 4: protobuf.AuctionsRunService.GetCurrentPrice:output_type -> protobuf.GetCurrentPriceResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_auctions_run_service_proto_init() }
func file_protobuf_auctions_run_service_proto_init() {
	if File_protobuf_auctions_run_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_auctions_run_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestPriceRequest); i {
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
		file_protobuf_auctions_run_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestPriceResponse); i {
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
		file_protobuf_auctions_run_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentPriceRequest); i {
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
		file_protobuf_auctions_run_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentPriceResponse); i {
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
			RawDescriptor: file_protobuf_auctions_run_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_auctions_run_service_proto_goTypes,
		DependencyIndexes: file_protobuf_auctions_run_service_proto_depIdxs,
		MessageInfos:      file_protobuf_auctions_run_service_proto_msgTypes,
	}.Build()
	File_protobuf_auctions_run_service_proto = out.File
	file_protobuf_auctions_run_service_proto_rawDesc = nil
	file_protobuf_auctions_run_service_proto_goTypes = nil
	file_protobuf_auctions_run_service_proto_depIdxs = nil
}
