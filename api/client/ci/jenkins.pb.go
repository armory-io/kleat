// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: ci/jenkins.proto

package ci

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	client "github.com/spinnaker/kleat/api/client"
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

// Configuration for Jenkins.
type Jenkins struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether Jenkins is enabled.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured Jenkins accounts.
	Masters []*JenkinsAccount `protobuf:"bytes,2,rep,name=masters,proto3" json:"masters,omitempty"`
}

func (x *Jenkins) Reset() {
	*x = Jenkins{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_jenkins_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jenkins) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jenkins) ProtoMessage() {}

func (x *Jenkins) ProtoReflect() protoreflect.Message {
	mi := &file_ci_jenkins_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jenkins.ProtoReflect.Descriptor instead.
func (*Jenkins) Descriptor() ([]byte, []int) {
	return file_ci_jenkins_proto_rawDescGZIP(), []int{0}
}

func (x *Jenkins) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *Jenkins) GetMasters() []*JenkinsAccount {
	if x != nil {
		return x.Masters
	}
	return nil
}

// Configuration for a Jenkins account.
type JenkinsAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) The username of the Jenkins user to authenticate as.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// (Required) The password of the Jenkins user to authenticate as.
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// (Required) The address at which the Jenkins server is reachable.
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// Whether or not to negotiate CSRF tokens when calling Jenkins.
	Csrf *wrappers.BoolValue `protobuf:"bytes,5,opt,name=csrf,proto3" json:"csrf,omitempty"`
	// Fiat permissions configuration. A user must have at least one of the READ
	// roles in order to view this build account or use it as a trigger source.
	// A user must have at least one of the WRITE roles in order to run jobs on
	// this build account.
	Permissions *client.Permissions `protobuf:"bytes,6,opt,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *JenkinsAccount) Reset() {
	*x = JenkinsAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_jenkins_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JenkinsAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JenkinsAccount) ProtoMessage() {}

func (x *JenkinsAccount) ProtoReflect() protoreflect.Message {
	mi := &file_ci_jenkins_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JenkinsAccount.ProtoReflect.Descriptor instead.
func (*JenkinsAccount) Descriptor() ([]byte, []int) {
	return file_ci_jenkins_proto_rawDescGZIP(), []int{1}
}

func (x *JenkinsAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JenkinsAccount) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *JenkinsAccount) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *JenkinsAccount) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *JenkinsAccount) GetCsrf() *wrappers.BoolValue {
	if x != nil {
		return x.Csrf
	}
	return nil
}

func (x *JenkinsAccount) GetPermissions() *client.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_ci_jenkins_proto protoreflect.FileDescriptor

var file_ci_jenkins_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x69, 0x2f, 0x6a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x73, 0x0a, 0x07, 0x4a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x32, 0x0a, 0x07, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x2e, 0x4a, 0x65, 0x6e,
	0x6b, 0x69, 0x6e, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x22, 0xdc, 0x01, 0x0a, 0x0e, 0x4a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a,
	0x04, 0x63, 0x73, 0x72, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x63, 0x73, 0x72, 0x66, 0x12, 0x34, 0x0a,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ci_jenkins_proto_rawDescOnce sync.Once
	file_ci_jenkins_proto_rawDescData = file_ci_jenkins_proto_rawDesc
)

func file_ci_jenkins_proto_rawDescGZIP() []byte {
	file_ci_jenkins_proto_rawDescOnce.Do(func() {
		file_ci_jenkins_proto_rawDescData = protoimpl.X.CompressGZIP(file_ci_jenkins_proto_rawDescData)
	})
	return file_ci_jenkins_proto_rawDescData
}

var file_ci_jenkins_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ci_jenkins_proto_goTypes = []interface{}{
	(*Jenkins)(nil),            // 0: proto.ci.Jenkins
	(*JenkinsAccount)(nil),     // 1: proto.ci.JenkinsAccount
	(*wrappers.BoolValue)(nil), // 2: google.protobuf.BoolValue
	(*client.Permissions)(nil), // 3: proto.Permissions
}
var file_ci_jenkins_proto_depIdxs = []int32{
	2, // 0: proto.ci.Jenkins.enabled:type_name -> google.protobuf.BoolValue
	1, // 1: proto.ci.Jenkins.masters:type_name -> proto.ci.JenkinsAccount
	2, // 2: proto.ci.JenkinsAccount.csrf:type_name -> google.protobuf.BoolValue
	3, // 3: proto.ci.JenkinsAccount.permissions:type_name -> proto.Permissions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ci_jenkins_proto_init() }
func file_ci_jenkins_proto_init() {
	if File_ci_jenkins_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ci_jenkins_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jenkins); i {
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
		file_ci_jenkins_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JenkinsAccount); i {
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
			RawDescriptor: file_ci_jenkins_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ci_jenkins_proto_goTypes,
		DependencyIndexes: file_ci_jenkins_proto_depIdxs,
		MessageInfos:      file_ci_jenkins_proto_msgTypes,
	}.Build()
	File_ci_jenkins_proto = out.File
	file_ci_jenkins_proto_rawDesc = nil
	file_ci_jenkins_proto_goTypes = nil
	file_ci_jenkins_proto_depIdxs = nil
}
