package constant

import "github.com/ammmnia/tools/constant"

const (
	// verificationCode used for.
	VerificationCodeForRegister      = 1 // Register
	VerificationCodeForResetPassword = 2 // Reset password
	VerificationCodeForLogin         = 3 // Login
)

const LogFileName = "chat.log"

// block unblock.
const (
	BlockUser   = 1
	UnblockUser = 2
)

// Mode.
const (
	UserMode  = "user"
	AdminMode = "admin"
)

const DefaultAdminLevel = 100

// user level.
const (
	NormalAdmin       = 80
	AdvancedUserLevel = 100
)

// AddFriendCtrl.
const (
	OrdinaryUserAddFriendEnable  = 1  // Allow ordinary users to add friends
	OrdinaryUserAddFriendDisable = -1 // Do not allow ordinary users to add friends
)

const (
	NormalUser = 1
	AdminUser  = 2
)

// mini-app
const (
	StatusOnShelf = 1 // OnShelf
	StatusUnShelf = 2 // UnShelf
)

const (
	LimitNil             = 0 // None
	LimitEmpty           = 1 // Neither are restricted
	LimitOnlyLoginIP     = 2 // Only login is restricted
	LimitOnlyRegisterIP  = 3 // Only registration is restricted
	LimitLoginIP         = 4 // Restrict login
	LimitRegisterIP      = 5 // Restrict registration
	LimitLoginRegisterIP = 6 // Restrict both login and registration
)

const (
	InvitationCodeAll    = 0 // All
	InvitationCodeUsed   = 1 // Used
	InvitationCodeUnused = 2 // Unused
)

const (
	RpcOpUserID                      = constant.OpUserID
	RpcCustomHeader                  = constant.RpcCustomHeader
	RpcOpUserType   constant.MetaKey = "opUserType"
	RpcOpAccount    constant.MetaKey = "opAccount"
)

const NeedInvitationCodeRegisterConfigKey = "needInvitationCodeRegister"

const (
	DefaultAllowVibration = 1
	DefaultAllowBeep      = 1
	DefaultAllowAddFriend = 1
	DefaultAddFriendLimit = 10
)

const (
	FinDAllUser    = 0
	FindNormalUser = 1
)

const CtxApiToken = "api-token"
