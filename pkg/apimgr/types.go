package apimgr

type Api struct {
	Id        int    `json:"id"`
	Aid       int    `json:"aid"`
	Num       string `json:"num"`
	Url       string `json:"url"`
	Name      string `json:"name"`
	Des       string `json:"des"`
	Parameter string `json:"parameter"`
	Memo      string `json:"memo"`
	Re        string `json:"re"`
	St        string `json:"st"`
	Lasttime  int64  `json:"lasttime"`
	Lastuid   int    `json:"lastuid"`
	Isdel     int    `json:"isdel"`
	Type      string `json:"type"`
	Ord       int    `json:"ord"`
	OpUser    string `json:"op_user"`
}

type Cateogry struct {
	Aid     int    `json:"aid"`
	Cname   string `json:"cname"`
	Cdesc   string `json:"cdesc"`
	Csort   int    `json:"csort"`
	IsDel   int    `json:"isdel"`
	Addtime int    `json:"addtime"`
}

type User struct {
	Id       int    `json:"id"`
	PassWord string `json:"login_pwd"`
	UserBase
}

type UserBase struct {
	Id        int    `json:"id"`
	LoginName string `json:"login_name"`
	NiceName  string `json:"nice_name"`
	Role      int    `json:"role"`
	IsDel     int    `json:"isdel"`
}
