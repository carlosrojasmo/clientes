// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: proto/proto.proto

package proto

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

type SendToOrden struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPaquete string `protobuf:"bytes,1,opt,name=idPaquete,proto3" json:"idPaquete,omitempty"`
	Tipo      string `protobuf:"bytes,2,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Nombre    string `protobuf:"bytes,3,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Valor     int64  `protobuf:"varint,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Origen    string `protobuf:"bytes,5,opt,name=origen,proto3" json:"origen,omitempty"`
	Destino   string `protobuf:"bytes,6,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *SendToOrden) Reset() {
	*x = SendToOrden{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendToOrden) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendToOrden) ProtoMessage() {}

func (x *SendToOrden) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendToOrden.ProtoReflect.Descriptor instead.
func (*SendToOrden) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{0}
}

func (x *SendToOrden) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *SendToOrden) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *SendToOrden) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *SendToOrden) GetValor() int64 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *SendToOrden) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *SendToOrden) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

type ReplyFromOrden struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seguimiento int64 `protobuf:"varint,1,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
}

func (x *ReplyFromOrden) Reset() {
	*x = ReplyFromOrden{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyFromOrden) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyFromOrden) ProtoMessage() {}

func (x *ReplyFromOrden) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyFromOrden.ProtoReflect.Descriptor instead.
func (*ReplyFromOrden) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyFromOrden) GetSeguimiento() int64 {
	if x != nil {
		return x.Seguimiento
	}
	return 0
}

type InfoSeguimiento struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado string `protobuf:"bytes,1,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *InfoSeguimiento) Reset() {
	*x = InfoSeguimiento{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoSeguimiento) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoSeguimiento) ProtoMessage() {}

func (x *InfoSeguimiento) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoSeguimiento.ProtoReflect.Descriptor instead.
func (*InfoSeguimiento) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{2}
}

func (x *InfoSeguimiento) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

var File_proto_proto_proto protoreflect.FileDescriptor

var file_proto_proto_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x0b, 0x73,
	0x65, 0x6e, 0x64, 0x54, 0x6f, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64,
	0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x22, 0x32, 0x0a, 0x0e,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f,
	0x22, 0x29, 0x0a, 0x0f, 0x69, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65,
	0x6e, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x32, 0x84, 0x01, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0c,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x4f, 0x72, 0x64, 0x65, 0x6e,
	0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72,
	0x6f, 0x6d, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x12, 0x39, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e,
	0x74, 0x6f, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_proto_proto_rawDescOnce sync.Once
	file_proto_proto_proto_rawDescData = file_proto_proto_proto_rawDesc
)

func file_proto_proto_proto_rawDescGZIP() []byte {
	file_proto_proto_proto_rawDescOnce.Do(func() {
		file_proto_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proto_proto_rawDescData)
	})
	return file_proto_proto_proto_rawDescData
}

var file_proto_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_proto_proto_goTypes = []interface{}{
	(*SendToOrden)(nil),     // 0: proto.sendToOrden
	(*ReplyFromOrden)(nil),  // 1: proto.replyFromOrden
	(*InfoSeguimiento)(nil), // 2: proto.infoSeguimiento
}
var file_proto_proto_proto_depIdxs = []int32{
	0, // 0: proto.OrdenService.replyToOrder:input_type -> proto.sendToOrden
	1, // 1: proto.OrdenService.getState:input_type -> proto.replyFromOrden
	1, // 2: proto.OrdenService.replyToOrder:output_type -> proto.replyFromOrden
	2, // 3: proto.OrdenService.getState:output_type -> proto.infoSeguimiento
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_proto_proto_init() }
func file_proto_proto_proto_init() {
	if File_proto_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendToOrden); i {
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
		file_proto_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyFromOrden); i {
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
		file_proto_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoSeguimiento); i {
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
			RawDescriptor: file_proto_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_proto_proto_goTypes,
		DependencyIndexes: file_proto_proto_proto_depIdxs,
		MessageInfos:      file_proto_proto_proto_msgTypes,
	}.Build()
	File_proto_proto_proto = out.File
	file_proto_proto_proto_rawDesc = nil
	file_proto_proto_proto_goTypes = nil
	file_proto_proto_proto_depIdxs = nil
}
