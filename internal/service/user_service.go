package service

import "context"

// --- Interfaces ---
type (
	IUserLogin interface {
		Login(ctx context.Context) error
	}

	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
		GetAllUser(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

// --- Local service singletons ---
var (
	localUserInfo  IUserInfo
	localUserAdmin IUserAdmin
	localUserLogin IUserLogin
)

// --- Getters ---
func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("IUserAdmin not implemented")
	}
	return localUserAdmin
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("IUserInfo not implemented")
	}
	return localUserInfo
}

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("IUserLogin not implemented")
	}
	return localUserLogin
}

// --- Initializers ---
func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}
