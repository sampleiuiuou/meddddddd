package types

// login request structure
type LoginReqStruct struct {
	UserId   string `json:"userid"`
	Password string `json:"password"`
}

// store the login user details in struture
type UserStruct struct {
	LoginId int    `json:"login_id"`
	Role    string `json:"role"`
	UserId  string `json:"userid"`
	LHId    int    `json:"lh_id"`
}

// Login Response structure
type LoginRespStruct struct {
	UserDetails UserStruct `json:"userdetails"`
	ErrMsg      string     `json:"errmsg"`
	Status      string     `json:"status"`
	Msg         bool       `json:"msg"`
}

// Add New User Request Structure
type AddUserReqStruct struct {
	UserId        string `json:"userid"`
	Password      string `json:"password"`
	Role          string `json:"role"`
	CreaterUserId string `json:"cuser_id"`
}

// Common Response Structure
type CommonRespStruct struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    bool   `json:"msg"`
}
type LoginHistory struct {
	Entry  string `json:"entry"`
	Date   string `json:"date"`
	Time   string `json:"time"`
	UserId string `json:"userid"`
}
type ListHistoryRespStruct struct {
	LoginArr []LoginHistory `json:"history"`
	ErrMsg   string         `json:"errmsg"`
	Status   string         `json:"status"`
	Msg      bool           `json:"msg"`
}
