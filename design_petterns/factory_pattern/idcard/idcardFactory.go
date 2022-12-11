package idcard

import (
	"develop/gotraining/factory_pattern/framework"
)

// インスタンスを生成する具体的な処理を記載していく
// IdcardFactory メソッドcreateProduct, registerProductを実装しているクラス
// FactoryIFを満たすようにメソッドを実装している
type IdcardFactory struct {
	owners []string
}

func NewIdCardFactory() framework.Factory {
	return framework.Factory{
		FactoryIF: &IdcardFactory{
			owners: []string{},
		},
	}
}

func (c IdcardFactory) CreateProduct(owner string) framework.Product {
	return NewIDCard(owner)
}

func (c IdcardFactory) RegisterProduct(product framework.Product) {

	idCard, ok := product.(IdCard)
	if ok {
		c.owners = append(c.owners, idCard.GetOwner())
	}
}
