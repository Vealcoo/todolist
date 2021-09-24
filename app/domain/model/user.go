package model

type UserInfo struct {
	UserID   string `bson:"userid"`
	UserPW   string `bson:"userpassword"`
	UserName string `bson:"username"`
}

type UserLoginInfo struct {
	UserID   string `json:"userid"`
	UserName string `json:"username"`
}
