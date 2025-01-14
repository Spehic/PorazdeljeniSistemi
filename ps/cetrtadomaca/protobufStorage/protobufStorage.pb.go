// prevajanje datoteke *.proto na Arnes HPC:
//   - namestitev modulov
//      module load protobuf/23.0-GCCcore-12.2.0
//      module load binutils/2.39-GCCcore-12.2.0
//      go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0 # samo ob prvi uporabi
//      go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 # samo ob prvi uporabi
//      export PATH="$PATH:$(go env GOPATH)/bin"
//   - prevajanje
//      srun protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protobufStorage.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.0
// source: protobufStorage.proto

package protobufStorage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task      string `protobuf:"bytes,1,opt,name=Task,proto3" json:"Task,omitempty"`
	Completed bool   `protobuf:"varint,2,opt,name=Completed,proto3" json:"Completed,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufStorage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_protobufStorage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_protobufStorage_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *Todo) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

type TodoStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *TodoStorage) Reset() {
	*x = TodoStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufStorage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoStorage) ProtoMessage() {}

func (x *TodoStorage) ProtoReflect() protoreflect.Message {
	mi := &file_protobufStorage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoStorage.ProtoReflect.Descriptor instead.
func (*TodoStorage) Descriptor() ([]byte, []int) {
	return file_protobufStorage_proto_rawDescGZIP(), []int{1}
}

func (x *TodoStorage) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

var File_protobufStorage_proto protoreflect.FileDescriptor

var file_protobufStorage_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22,
	0x3a, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2b,
	0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x32, 0xb1, 0x01, 0x0a, 0x04,
	0x43, 0x52, 0x55, 0x44, 0x12, 0x34, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x12, 0x5a, 0x10, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobufStorage_proto_rawDescOnce sync.Once
	file_protobufStorage_proto_rawDescData = file_protobufStorage_proto_rawDesc
)

func file_protobufStorage_proto_rawDescGZIP() []byte {
	file_protobufStorage_proto_rawDescOnce.Do(func() {
		file_protobufStorage_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobufStorage_proto_rawDescData)
	})
	return file_protobufStorage_proto_rawDescData
}

var file_protobufStorage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobufStorage_proto_goTypes = []interface{}{
	(*Todo)(nil),          // 0: protobufStorage.Todo
	(*TodoStorage)(nil),   // 1: protobufStorage.TodoStorage
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_protobufStorage_proto_depIdxs = []int32{
	0, // 0: protobufStorage.TodoStorage.todos:type_name -> protobufStorage.Todo
	0, // 1: protobufStorage.CRUD.Put:input_type -> protobufStorage.Todo
	0, // 2: protobufStorage.CRUD.Get:input_type -> protobufStorage.Todo
	0, // 3: protobufStorage.CRUD.Commit:input_type -> protobufStorage.Todo
	2, // 4: protobufStorage.CRUD.Put:output_type -> google.protobuf.Empty
	1, // 5: protobufStorage.CRUD.Get:output_type -> protobufStorage.TodoStorage
	2, // 6: protobufStorage.CRUD.Commit:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobufStorage_proto_init() }
func file_protobufStorage_proto_init() {
	if File_protobufStorage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobufStorage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_protobufStorage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoStorage); i {
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
			RawDescriptor: file_protobufStorage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobufStorage_proto_goTypes,
		DependencyIndexes: file_protobufStorage_proto_depIdxs,
		MessageInfos:      file_protobufStorage_proto_msgTypes,
	}.Build()
	File_protobufStorage_proto = out.File
	file_protobufStorage_proto_rawDesc = nil
	file_protobufStorage_proto_goTypes = nil
	file_protobufStorage_proto_depIdxs = nil
}
