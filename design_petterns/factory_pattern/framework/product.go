package framework

// 最終的にクライアントサイドが使うメソッドを定義
// Product 抽象メソッドのみ定義されているクラス
type Product interface {
	Use()
}
