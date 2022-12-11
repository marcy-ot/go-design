package idcard

import (
	"develop/gotraining/factory_pattern/framework"
	"log"
)

// IdCard メソッドuseを実装しているクラス
type IdCard struct {
	owner string
}

func NewIDCard(owner string) framework.Product {
	log.Println(owner, "のカードを作ります")
	return &IdCard{owner: owner}
}

func (c IdCard) Use() {
	log.Println(c.owner, "のカードを使います")
}

func (c IdCard) GetOwner() string {
	return c.owner
}
