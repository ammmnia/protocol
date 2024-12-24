// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.2
// source: openchat/common/common.proto

package common

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

type UserFullInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID           string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Password         string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Account          string `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	PhoneNumber      string `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	AreaCode         string `protobuf:"bytes,5,opt,name=areaCode,proto3" json:"areaCode,omitempty"`
	Email            string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Nickname         string `protobuf:"bytes,7,opt,name=nickname,proto3" json:"nickname,omitempty"`
	FaceURL          string `protobuf:"bytes,8,opt,name=faceURL,proto3" json:"faceURL,omitempty"`
	Gender           int32  `protobuf:"varint,9,opt,name=gender,proto3" json:"gender,omitempty"`
	Level            int32  `protobuf:"varint,10,opt,name=level,proto3" json:"level,omitempty"`
	Birth            int64  `protobuf:"varint,11,opt,name=birth,proto3" json:"birth,omitempty"`
	AllowAddFriend   int32  `protobuf:"varint,12,opt,name=allowAddFriend,proto3" json:"allowAddFriend,omitempty"`
	AllowBeep        int32  `protobuf:"varint,13,opt,name=allowBeep,proto3" json:"allowBeep,omitempty"`
	AllowVibration   int32  `protobuf:"varint,14,opt,name=allowVibration,proto3" json:"allowVibration,omitempty"`
	GlobalRecvMsgOpt int32  `protobuf:"varint,15,opt,name=globalRecvMsgOpt,proto3" json:"globalRecvMsgOpt,omitempty"`
	RegisterType     int32  `protobuf:"varint,16,opt,name=registerType,proto3" json:"registerType,omitempty"`
	CreateTime       int64  `protobuf:"varint,17,opt,name=createTime,proto3" json:"createTime,omitempty"`       //注册时间
	ChangeTime       int64  `protobuf:"varint,18,opt,name=changeTime,proto3" json:"changeTime,omitempty"`       //操作时间
	LastLoginTime    int64  `protobuf:"varint,19,opt,name=lastLoginTime,proto3" json:"lastLoginTime,omitempty"` //最近登录时间
	Operator         string `protobuf:"bytes,20,opt,name=operator,proto3" json:"operator,omitempty"`            //操作人
	UserStatus       int32  `protobuf:"varint,21,opt,name=userStatus,proto3" json:"userStatus,omitempty"`       //用户状态
	OnlineStatus     int32  `protobuf:"varint,22,opt,name=onlineStatus,proto3" json:"onlineStatus,omitempty"`   //在线状态
	AddFriendLimit   int32  `protobuf:"varint,23,opt,name=AddFriendLimit,proto3" json:"AddFriendLimit,omitempty"`
}

func (x *UserFullInfo) Reset() {
	*x = UserFullInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openchat_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFullInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFullInfo) ProtoMessage() {}

func (x *UserFullInfo) ProtoReflect() protoreflect.Message {
	mi := &file_openchat_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFullInfo.ProtoReflect.Descriptor instead.
func (*UserFullInfo) Descriptor() ([]byte, []int) {
	return file_openchat_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *UserFullInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserFullInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserFullInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserFullInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserFullInfo) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

func (x *UserFullInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserFullInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserFullInfo) GetFaceURL() string {
	if x != nil {
		return x.FaceURL
	}
	return ""
}

func (x *UserFullInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserFullInfo) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *UserFullInfo) GetBirth() int64 {
	if x != nil {
		return x.Birth
	}
	return 0
}

func (x *UserFullInfo) GetAllowAddFriend() int32 {
	if x != nil {
		return x.AllowAddFriend
	}
	return 0
}

func (x *UserFullInfo) GetAllowBeep() int32 {
	if x != nil {
		return x.AllowBeep
	}
	return 0
}

func (x *UserFullInfo) GetAllowVibration() int32 {
	if x != nil {
		return x.AllowVibration
	}
	return 0
}

func (x *UserFullInfo) GetGlobalRecvMsgOpt() int32 {
	if x != nil {
		return x.GlobalRecvMsgOpt
	}
	return 0
}

func (x *UserFullInfo) GetRegisterType() int32 {
	if x != nil {
		return x.RegisterType
	}
	return 0
}

func (x *UserFullInfo) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *UserFullInfo) GetChangeTime() int64 {
	if x != nil {
		return x.ChangeTime
	}
	return 0
}

func (x *UserFullInfo) GetLastLoginTime() int64 {
	if x != nil {
		return x.LastLoginTime
	}
	return 0
}

func (x *UserFullInfo) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *UserFullInfo) GetUserStatus() int32 {
	if x != nil {
		return x.UserStatus
	}
	return 0
}

func (x *UserFullInfo) GetOnlineStatus() int32 {
	if x != nil {
		return x.OnlineStatus
	}
	return 0
}

func (x *UserFullInfo) GetAddFriendLimit() int32 {
	if x != nil {
		return x.AddFriendLimit
	}
	return 0
}

type UserPublicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Account  string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Nickname string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	FaceURL  string `protobuf:"bytes,5,opt,name=faceURL,proto3" json:"faceURL,omitempty"`
	Gender   int32  `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Level    int32  `protobuf:"varint,7,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *UserPublicInfo) Reset() {
	*x = UserPublicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openchat_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPublicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPublicInfo) ProtoMessage() {}

func (x *UserPublicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_openchat_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPublicInfo.ProtoReflect.Descriptor instead.
func (*UserPublicInfo) Descriptor() ([]byte, []int) {
	return file_openchat_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *UserPublicInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserPublicInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserPublicInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserPublicInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserPublicInfo) GetFaceURL() string {
	if x != nil {
		return x.FaceURL
	}
	return ""
}

func (x *UserPublicInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserPublicInfo) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type UserIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	AreaCode    string `protobuf:"bytes,2,opt,name=areaCode,proto3" json:"areaCode,omitempty"`
	PhoneNumber string `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	DeviceID    string `protobuf:"bytes,4,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	Platform    int32  `protobuf:"varint,5,opt,name=platform,proto3" json:"platform,omitempty"`
	Account     string `protobuf:"bytes,6,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *UserIdentity) Reset() {
	*x = UserIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openchat_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdentity) ProtoMessage() {}

func (x *UserIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_openchat_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdentity.ProtoReflect.Descriptor instead.
func (*UserIdentity) Descriptor() ([]byte, []int) {
	return file_openchat_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *UserIdentity) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserIdentity) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

func (x *UserIdentity) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserIdentity) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *UserIdentity) GetPlatform() int32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *UserIdentity) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type AppletInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AppID      string `protobuf:"bytes,3,opt,name=appID,proto3" json:"appID,omitempty"`
	Icon       string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	Url        string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Md5        string `protobuf:"bytes,6,opt,name=md5,proto3" json:"md5,omitempty"`
	Size       int64  `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	Version    string `protobuf:"bytes,8,opt,name=version,proto3" json:"version,omitempty"`
	Priority   uint32 `protobuf:"varint,9,opt,name=priority,proto3" json:"priority,omitempty"`
	Status     uint32 `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime int64  `protobuf:"varint,11,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *AppletInfo) Reset() {
	*x = AppletInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openchat_common_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppletInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppletInfo) ProtoMessage() {}

func (x *AppletInfo) ProtoReflect() protoreflect.Message {
	mi := &file_openchat_common_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppletInfo.ProtoReflect.Descriptor instead.
func (*AppletInfo) Descriptor() ([]byte, []int) {
	return file_openchat_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *AppletInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppletInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppletInfo) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *AppletInfo) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *AppletInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AppletInfo) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *AppletInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *AppletInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AppletInfo) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *AppletInfo) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AppletInfo) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type LogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Platform   int32  `protobuf:"varint,2,opt,name=platform,proto3" json:"platform,omitempty"`
	Url        string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	CreateTime int64  `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Nickname   string `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	LogID      string `protobuf:"bytes,6,opt,name=logID,proto3" json:"logID,omitempty"`
	Filename   string `protobuf:"bytes,7,opt,name=filename,proto3" json:"filename,omitempty"`
	SystemType string `protobuf:"bytes,8,opt,name=systemType,proto3" json:"systemType,omitempty"`
	Ex         string `protobuf:"bytes,9,opt,name=ex,proto3" json:"ex,omitempty"`
	Version    string `protobuf:"bytes,10,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *LogInfo) Reset() {
	*x = LogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openchat_common_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogInfo) ProtoMessage() {}

func (x *LogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_openchat_common_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogInfo.ProtoReflect.Descriptor instead.
func (*LogInfo) Descriptor() ([]byte, []int) {
	return file_openchat_common_common_proto_rawDescGZIP(), []int{4}
}

func (x *LogInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *LogInfo) GetPlatform() int32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *LogInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *LogInfo) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *LogInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *LogInfo) GetLogID() string {
	if x != nil {
		return x.LogID
	}
	return ""
}

func (x *LogInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *LogInfo) GetSystemType() string {
	if x != nil {
		return x.SystemType
	}
	return ""
}

func (x *LogInfo) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

func (x *LogInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type UserLoginRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	LoginTime int64  `protobuf:"varint,2,opt,name=loginTime,proto3" json:"loginTime,omitempty"`
	IP        string `protobuf:"bytes,3,opt,name=iP,proto3" json:"iP,omitempty"`
	DeviceID  string `protobuf:"bytes,4,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	Platform  string `protobuf:"bytes,5,opt,name=platform,proto3" json:"platform,omitempty"`
}

func (x *UserLoginRecord) Reset() {
	*x = UserLoginRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openchat_common_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRecord) ProtoMessage() {}

func (x *UserLoginRecord) ProtoReflect() protoreflect.Message {
	mi := &file_openchat_common_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRecord.ProtoReflect.Descriptor instead.
func (*UserLoginRecord) Descriptor() ([]byte, []int) {
	return file_openchat_common_common_proto_rawDescGZIP(), []int{5}
}

func (x *UserLoginRecord) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserLoginRecord) GetLoginTime() int64 {
	if x != nil {
		return x.LoginTime
	}
	return 0
}

func (x *UserLoginRecord) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *UserLoginRecord) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *UserLoginRecord) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

var File_openchat_common_common_proto protoreflect.FileDescriptor

var file_openchat_common_common_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xd6, 0x05, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x64, 0x64, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x42, 0x65, 0x65, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x42, 0x65, 0x65, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x56, 0x69, 0x62, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x56, 0x69, 0x62, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x76, 0x4d,
	0x73, 0x67, 0x4f, 0x70, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x76, 0x4d, 0x73, 0x67, 0x4f, 0x70, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0xbc, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52,
	0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0xb4,
	0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x80, 0x02, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x87, 0x02, 0x0a, 0x07, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x44, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x78, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x50, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6d, 0x6d, 0x6d, 0x6e, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_openchat_common_common_proto_rawDescOnce sync.Once
	file_openchat_common_common_proto_rawDescData = file_openchat_common_common_proto_rawDesc
)

func file_openchat_common_common_proto_rawDescGZIP() []byte {
	file_openchat_common_common_proto_rawDescOnce.Do(func() {
		file_openchat_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_openchat_common_common_proto_rawDescData)
	})
	return file_openchat_common_common_proto_rawDescData
}

var file_openchat_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_openchat_common_common_proto_goTypes = []interface{}{
	(*UserFullInfo)(nil),    // 0: openim.openchat.common.UserFullInfo
	(*UserPublicInfo)(nil),  // 1: openim.openchat.common.UserPublicInfo
	(*UserIdentity)(nil),    // 2: openim.openchat.common.UserIdentity
	(*AppletInfo)(nil),      // 3: openim.openchat.common.AppletInfo
	(*LogInfo)(nil),         // 4: openim.openchat.common.LogInfo
	(*UserLoginRecord)(nil), // 5: openim.openchat.common.UserLoginRecord
}
var file_openchat_common_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_openchat_common_common_proto_init() }
func file_openchat_common_common_proto_init() {
	if File_openchat_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_openchat_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFullInfo); i {
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
		file_openchat_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPublicInfo); i {
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
		file_openchat_common_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdentity); i {
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
		file_openchat_common_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppletInfo); i {
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
		file_openchat_common_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogInfo); i {
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
		file_openchat_common_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginRecord); i {
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
			RawDescriptor: file_openchat_common_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_openchat_common_common_proto_goTypes,
		DependencyIndexes: file_openchat_common_common_proto_depIdxs,
		MessageInfos:      file_openchat_common_common_proto_msgTypes,
	}.Build()
	File_openchat_common_common_proto = out.File
	file_openchat_common_common_proto_rawDesc = nil
	file_openchat_common_common_proto_goTypes = nil
	file_openchat_common_common_proto_depIdxs = nil
}
