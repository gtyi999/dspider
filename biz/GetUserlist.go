package biz

type FuncGetUserlist struct {
	input  interface{}
	output interface{}
}

func NewFuncGetUserlist() *FuncGetUserlist {
	return &FuncGetUserlist{}
}

func (this *FuncGetUserlist) GetHtml() (err error) {
	return nil
}
