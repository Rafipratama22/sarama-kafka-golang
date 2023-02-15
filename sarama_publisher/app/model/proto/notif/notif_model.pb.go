// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: notif_model.proto

package notif

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

type NotifRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title            string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body             string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	To               string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Image            string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Channel          string `protobuf:"bytes,5,opt,name=channel,proto3" json:"channel,omitempty"`
	NotificationId   int32  `protobuf:"varint,6,opt,name=notificationId,proto3" json:"notificationId,omitempty"`
	NotificationType string `protobuf:"bytes,7,opt,name=notificationType,proto3" json:"notificationType,omitempty"`
}

func (x *NotifRequest) Reset() {
	*x = NotifRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notif_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifRequest) ProtoMessage() {}

func (x *NotifRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notif_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifRequest.ProtoReflect.Descriptor instead.
func (*NotifRequest) Descriptor() ([]byte, []int) {
	return file_notif_model_proto_rawDescGZIP(), []int{0}
}

func (x *NotifRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NotifRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *NotifRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *NotifRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *NotifRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *NotifRequest) GetNotificationId() int32 {
	if x != nil {
		return x.NotificationId
	}
	return 0
}

func (x *NotifRequest) GetNotificationType() string {
	if x != nil {
		return x.NotificationType
	}
	return ""
}

var File_notif_model_proto protoreflect.FileDescriptor

var file_notif_model_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xcc, 0x01, 0x0a, 0x0c, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notif_model_proto_rawDescOnce sync.Once
	file_notif_model_proto_rawDescData = file_notif_model_proto_rawDesc
)

func file_notif_model_proto_rawDescGZIP() []byte {
	file_notif_model_proto_rawDescOnce.Do(func() {
		file_notif_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_notif_model_proto_rawDescData)
	})
	return file_notif_model_proto_rawDescData
}

var file_notif_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_notif_model_proto_goTypes = []interface{}{
	(*NotifRequest)(nil), // 0: model.NotifRequest
}
var file_notif_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notif_model_proto_init() }
func file_notif_model_proto_init() {
	if File_notif_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notif_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifRequest); i {
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
			RawDescriptor: file_notif_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notif_model_proto_goTypes,
		DependencyIndexes: file_notif_model_proto_depIdxs,
		MessageInfos:      file_notif_model_proto_msgTypes,
	}.Build()
	File_notif_model_proto = out.File
	file_notif_model_proto_rawDesc = nil
	file_notif_model_proto_goTypes = nil
	file_notif_model_proto_depIdxs = nil
}
