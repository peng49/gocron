// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: internal/modules/rpc/proto/task.proto

package rpc

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

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`  // 命令
	Timeout int32  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"` // 任务执行超时时间
	Id      int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`           // 执行任务唯一ID
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_modules_rpc_proto_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_modules_rpc_proto_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_internal_modules_rpc_proto_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *TaskRequest) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *TaskRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"` // 命令标准输出
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`   // 命令错误
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_modules_rpc_proto_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_modules_rpc_proto_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_internal_modules_rpc_proto_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *TaskResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_internal_modules_rpc_proto_task_proto protoreflect.FileDescriptor

var file_internal_modules_rpc_proto_task_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0c, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x2c, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x24, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x0c, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_modules_rpc_proto_task_proto_rawDescOnce sync.Once
	file_internal_modules_rpc_proto_task_proto_rawDescData = file_internal_modules_rpc_proto_task_proto_rawDesc
)

func file_internal_modules_rpc_proto_task_proto_rawDescGZIP() []byte {
	file_internal_modules_rpc_proto_task_proto_rawDescOnce.Do(func() {
		file_internal_modules_rpc_proto_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_modules_rpc_proto_task_proto_rawDescData)
	})
	return file_internal_modules_rpc_proto_task_proto_rawDescData
}

var file_internal_modules_rpc_proto_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_modules_rpc_proto_task_proto_goTypes = []interface{}{
	(*TaskRequest)(nil),  // 0: TaskRequest
	(*TaskResponse)(nil), // 1: TaskResponse
}
var file_internal_modules_rpc_proto_task_proto_depIdxs = []int32{
	0, // 0: Task.Run:input_type -> TaskRequest
	1, // 1: Task.Run:output_type -> TaskResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_modules_rpc_proto_task_proto_init() }
func file_internal_modules_rpc_proto_task_proto_init() {
	if File_internal_modules_rpc_proto_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_modules_rpc_proto_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
		file_internal_modules_rpc_proto_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResponse); i {
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
			RawDescriptor: file_internal_modules_rpc_proto_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_modules_rpc_proto_task_proto_goTypes,
		DependencyIndexes: file_internal_modules_rpc_proto_task_proto_depIdxs,
		MessageInfos:      file_internal_modules_rpc_proto_task_proto_msgTypes,
	}.Build()
	File_internal_modules_rpc_proto_task_proto = out.File
	file_internal_modules_rpc_proto_task_proto_rawDesc = nil
	file_internal_modules_rpc_proto_task_proto_goTypes = nil
	file_internal_modules_rpc_proto_task_proto_depIdxs = nil
}
