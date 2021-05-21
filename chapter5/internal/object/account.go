package object

// 账号资产
type AccountWithProperty struct {
	Id int64 `json:"id"`
	// 昵称
	Nickname string `json:"nickname"`
	// 头像地址
	Avatar string `json:"avatar"`
	// 性别
	Gender int8 `json:"gender"`
	// 手机号
	Mobile string `json:"mobile"`
	//  余额
	Balance int64 `json:"balance"`
}

// 账号 simple
type AccountSimple struct {
	Id int64 `json:"id"`
	// 昵称
	Nickname string `json:"nickname"`
	// 头像地址
	Avatar string `json:"avatar"`
	// 性别
	Gender int8 `json:"gender"`
	// 手机号
	Mobile string `json:"mobile"`
}
