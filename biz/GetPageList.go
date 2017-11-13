package biz

type FuncGetPageList struct {
	input  interface{}
	output interface{}
}

func NewFuncGetPageList() *FuncGetPageList {
	return &FuncGetPageList{}
}

func (this *FuncGetPageList) GetHtml() (err error) {
	return nil
}
