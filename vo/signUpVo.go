package vo

type SignUpVo struct {
	Id          uint32 `form:"-"`
	Name        string `form:"name" valid:"Required"`
	Age         uint32 `form:"age"`
	Sex         uint32 `form:"sex"`
	Phone       string `form:"phone" valid:"Required"`
	Create_time string
}
