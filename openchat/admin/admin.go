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

package admin

import (
	"github.com/ammmnia/protocol/openchat/common/constant"
	"github.com/ammmnia/tools/errs"
)

func (x *LoginReq) Check() error {
	if x.Account == "" {
		return errs.ErrArgs.WrapMsg("account is empty")
	}
	if x.Password == "" {
		return errs.ErrArgs.WrapMsg("password is empty")
	}
	return nil
}

func (x *ChangePasswordReq) Check() error {
	if x.Password == "" {
		return errs.ErrArgs.WrapMsg("password is empty")
	}
	return nil
}

func (x *CreateTokenReq) Check() error {
	if x.UserID == "" {
		return errs.ErrArgs.WrapMsg("userID is empty")
	}
	if x.UserType > constant.AdminUser || x.UserType < constant.NormalUser {
		return errs.ErrArgs.WrapMsg("userType is invalid")
	}
	return nil
}

func (x *ParseTokenReq) Check() error {
	if x.Token == "" {
		return errs.ErrArgs.WrapMsg("token is empty")
	}
	return nil
}

func (x *ChangeAdminPasswordReq) Check() error {
	if x.UserID == "" {
		return errs.ErrArgs.WrapMsg("userID is empty")
	}
	if x.CurrentPassword == "" {
		return errs.ErrArgs.WrapMsg("currentPassword is empty")
	}
	if x.NewPassword == "" {
		return errs.ErrArgs.WrapMsg("newPassword is empty")
	}
	if x.CurrentPassword == x.NewPassword {
		return errs.ErrArgs.WrapMsg("currentPassword is equal to newPassword")
	}
	return nil
}

func (x *AddAdminAccountReq) Check() error {
	if x.Account == "" {
		return errs.ErrArgs.WrapMsg("account is empty")
	}
	if x.Password == "" {
		return errs.ErrArgs.WrapMsg("password is empty")
	}
	return nil
}

//func (x *DelAdminAccountReq) Check() error {
//	if len(x.UserIDs) == 0 {
//		return errs.ErrArgs.WrapMsg("userIDs is empty")
//	}
//	return nil
//}

func (x *SearchAdminAccountReq) Check() error {
	if x.Pagination.ShowNumber == 0 {
		return errs.ErrArgs.WrapMsg("showNumber is empty")
	}
	if x.Pagination.PageNumber == 0 {
		return errs.ErrArgs.WrapMsg("pageNumber is empty")
	}
	return nil
}
