package framework

// FactoryIF 製品の作成と登録の処理を宣言
type FactoryIF interface {
	// CreateProduct 製品を作る
	CreateProduct(owner string) Product
	// RegisterProduct 作った製品を登録する
	RegisterProduct(product Product)
}

// Factory メソッドcreateを実装している抽象クラス
type Factory struct {
	FactoryIF
}

// Create インスタンスの生成手順を記述
func (f Factory) Create(owner string) Product {
	var p Product = f.CreateProduct(owner)
	f.RegisterProduct(p)

	return p
}
