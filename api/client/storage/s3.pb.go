// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: storage/s3.proto

package storage

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Configuration for S3 server-side encryption; values correspond to values of
// the 'x-amz-server-side-encryption' header.
type S3ServerSideEncryption int32

const (
	// Amazon S3-managed encryption keys, equivalent to a header value of 'AES256'.
	S3ServerSideEncryption_AES256 S3ServerSideEncryption = 0
	// AWS KMS-managed encryption keys, equivalent to a header value of 'aws:kms'.
	S3ServerSideEncryption_AWSKMS S3ServerSideEncryption = 1
)

// Enum value maps for S3ServerSideEncryption.
var (
	S3ServerSideEncryption_name = map[int32]string{
		0: "AES256",
		1: "AWSKMS",
	}
	S3ServerSideEncryption_value = map[string]int32{
		"AES256": 0,
		"AWSKMS": 1,
	}
)

func (x S3ServerSideEncryption) Enum() *S3ServerSideEncryption {
	p := new(S3ServerSideEncryption)
	*p = x
	return p
}

func (x S3ServerSideEncryption) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (S3ServerSideEncryption) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_s3_proto_enumTypes[0].Descriptor()
}

func (S3ServerSideEncryption) Type() protoreflect.EnumType {
	return &file_storage_s3_proto_enumTypes[0]
}

func (x S3ServerSideEncryption) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use S3ServerSideEncryption.Descriptor instead.
func (S3ServerSideEncryption) EnumDescriptor() ([]byte, []int) {
	return file_storage_s3_proto_rawDescGZIP(), []int{0}
}

// Configuration for an Amazon S3 persistent store.
type S3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether this persistent store is enabled.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The name of a storage bucket that your specified account has access to.
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// The root folder in the chosen bucket to place all of Spinnaker's persistent
	// data in.
	RootFolder string `protobuf:"bytes,3,opt,name=rootFolder,proto3" json:"rootFolder,omitempty"`
	// This is only required if the bucket you specify doesn't exist yet. In that
	// case, the bucket will be created in that region.
	// See http://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region.
	Region string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	// When true, use path-style to access bucket; when false, use virtual hosted-style
	// to access bucket.
	// See https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html#VirtualHostingExamples.
	PathStyleAccess *wrappers.BoolValue `protobuf:"bytes,5,opt,name=pathStyleAccess,proto3" json:"pathStyleAccess,omitempty"`
	// An alternate endpoint that your S3-compatible storage can be found at. This is
	// intended for self-hosted storage services with S3-compatible APIs, e.g. Minio.
	Endpoint string `protobuf:"bytes,6,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Your AWS Access Key ID. If not provided, Spinnaker will try to find AWS credentials
	// as described at http://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html#credentials-default
	AccessKeyId string `protobuf:"bytes,7,opt,name=accessKeyId,proto3" json:"accessKeyId,omitempty"`
	// Configuration for S3 server-size encryption.
	ServerSideEncryption S3ServerSideEncryption `protobuf:"varint,8,opt,name=serverSideEncryption,proto3,enum=proto.storage.S3ServerSideEncryption" json:"serverSideEncryption,omitempty"`
	// Your AWS Secret Key.
	SecretAccessKey string `protobuf:"bytes,9,opt,name=secretAccessKey,proto3" json:"secretAccessKey,omitempty"`
}

func (x *S3) Reset() {
	*x = S3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_s3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3) ProtoMessage() {}

func (x *S3) ProtoReflect() protoreflect.Message {
	mi := &file_storage_s3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3.ProtoReflect.Descriptor instead.
func (*S3) Descriptor() ([]byte, []int) {
	return file_storage_s3_proto_rawDescGZIP(), []int{0}
}

func (x *S3) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *S3) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *S3) GetRootFolder() string {
	if x != nil {
		return x.RootFolder
	}
	return ""
}

func (x *S3) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *S3) GetPathStyleAccess() *wrappers.BoolValue {
	if x != nil {
		return x.PathStyleAccess
	}
	return nil
}

func (x *S3) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *S3) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *S3) GetServerSideEncryption() S3ServerSideEncryption {
	if x != nil {
		return x.ServerSideEncryption
	}
	return S3ServerSideEncryption_AES256
}

func (x *S3) GetSecretAccessKey() string {
	if x != nil {
		return x.SecretAccessKey
	}
	return ""
}

var File_storage_s3_proto protoreflect.FileDescriptor

var file_storage_s3_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x93, 0x03, 0x0a, 0x02, 0x53, 0x33, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x74, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x74,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x44,
	0x0a, 0x0f, 0x70, 0x61, 0x74, 0x68, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0f, 0x70, 0x61, 0x74, 0x68, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x49, 0x64, 0x12, 0x59, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x53, 0x33, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x69, 0x64, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x0f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x2a, 0x30, 0x0a, 0x16, 0x53, 0x33, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x45, 0x53, 0x32, 0x35, 0x36, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x41, 0x57, 0x53, 0x4b, 0x4d, 0x53, 0x10, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65,
	0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_storage_s3_proto_rawDescOnce sync.Once
	file_storage_s3_proto_rawDescData = file_storage_s3_proto_rawDesc
)

func file_storage_s3_proto_rawDescGZIP() []byte {
	file_storage_s3_proto_rawDescOnce.Do(func() {
		file_storage_s3_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_s3_proto_rawDescData)
	})
	return file_storage_s3_proto_rawDescData
}

var file_storage_s3_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_s3_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_s3_proto_goTypes = []interface{}{
	(S3ServerSideEncryption)(0), // 0: proto.storage.S3ServerSideEncryption
	(*S3)(nil),                  // 1: proto.storage.S3
	(*wrappers.BoolValue)(nil),  // 2: google.protobuf.BoolValue
}
var file_storage_s3_proto_depIdxs = []int32{
	2, // 0: proto.storage.S3.enabled:type_name -> google.protobuf.BoolValue
	2, // 1: proto.storage.S3.pathStyleAccess:type_name -> google.protobuf.BoolValue
	0, // 2: proto.storage.S3.serverSideEncryption:type_name -> proto.storage.S3ServerSideEncryption
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_storage_s3_proto_init() }
func file_storage_s3_proto_init() {
	if File_storage_s3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_s3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3); i {
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
			RawDescriptor: file_storage_s3_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_s3_proto_goTypes,
		DependencyIndexes: file_storage_s3_proto_depIdxs,
		EnumInfos:         file_storage_s3_proto_enumTypes,
		MessageInfos:      file_storage_s3_proto_msgTypes,
	}.Build()
	File_storage_s3_proto = out.File
	file_storage_s3_proto_rawDesc = nil
	file_storage_s3_proto_goTypes = nil
	file_storage_s3_proto_depIdxs = nil
}
