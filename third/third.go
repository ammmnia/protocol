// Copyright © 2023 OpenIM. All rights reserved.
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

package third

import (
	"errors"
	"github.com/ammmnia/tools/constant"
)

func (x *FcmUpdateTokenReq) Check() error {
	if x.PlatformID > int32(constant.AdminPlatformID) || x.PlatformID < int32(constant.IOSPlatformID) {
		return errors.New("platformID is invalidate")
	}
	if x.FcmToken == "" {
		return errors.New("FcmToken is empty")
	}
	if x.Account == "" {
		return errors.New("account is empty")
	}
	return nil
}

func (x *SetAppBadgeReq) Check() error {
	if x.UserID == "" {
		return errors.New("UserID is empty")
	}
	return nil
}

func (x *InitiateMultipartUploadReq) Check() error {
	if x.UrlPrefix == "" {
		return errors.New("UrlPrefix is empty")
	}
	return nil
}

func (x *CompleteMultipartUploadReq) Check() error {
	if x.UrlPrefix == "" {
		return errors.New("UrlPrefix is empty")
	}
	return nil
}

func (x *CompleteFormDataReq) Check() error {
	if x.UrlPrefix == "" {
		return errors.New("UrlPrefix is empty")
	}
	return nil
}

func (x *LatestApplicationVersionReq) Check() error {
	if x.Platform == "" {
		return errors.New("platform is empty")
	}
	return nil
}

func (x *PageApplicationVersionReq) Check() error {
	if x.Pagination == nil {
		return errors.New("pagination is empty")
	}
	if x.Pagination.PageNumber < 1 {
		return errors.New("pageNumber is invalid")
	}
	if x.Pagination.ShowNumber < 1 {
		return errors.New("showNumber is invalid")
	}
	return nil
}
