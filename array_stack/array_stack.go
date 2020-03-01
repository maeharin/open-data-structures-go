package array_stack

// listインターフェース
// size()
// get(i)
// set(i,x)
// add(i,x)
// remove(i)

type ArrayStack struct {
	// リスト
	a []string

	// 現在の要素数
	n int
}

func NewArrayStack(cap int) ArrayStack {
	// cap分の領域を持つ配列を作成
	// 要素数は0
	a := make([]string, cap)
	return ArrayStack{a: a, n: 0}
}

// 現在の要素数を返す
func (as *ArrayStack) Size() int {
	return as.n
}

// インデックスの要素を返す
func (as *ArrayStack) Get(i int) string {
	return as.a[i]
}

// インデックスに要素をセットする
func (as *ArrayStack) Set(i int, new string) string {
	old := as.a[i]
	as.a[i] = new
	return old
}

// 所定のインデックスに要素を追加する
// 要素を追加することで配列の上限を超えるようであれば、配列を二倍にする
func (as *ArrayStack) Add(i int, v string) {
	// 配列が満タンなら、resizeする
	if as.n == len(as.a) {
		as.resize()
	}

	// 所定のインデックス以降の要素を
	for j := i; j < len(as.a)-1; j++ {
		as.a[j+1] = as.a[j]
	}

	// インデックスに値をセット
	as.a[i] = v

	// 要素数をインクリメント
	as.n++
}

// 所定のインデックスの要素を削除する
func (as *ArrayStack) Remove(i int) string {
	// 削除前の要素
	v := as.a[i]

	// 所定のインデックスから最後の一つ前まで要素を左にずらす
	for j := i; j < as.n-1; j++ {
		as.a[j] = as.a[j+1]
	}

	// 最後の要素を空にする
	// todo: サンプルコードにはこの処理がなかったがなぜ？この処理が無いと末尾に要素が残ってしまう
	as.a[as.n-1] = ""

	// 要素数をデクリメント
	as.n--

	// backingArrayに対して要素数が小さくなりすぎた場合（nがlengthの1/3以下）
	// backingArrayを縮小する（nがlengthの1/2になる）
	if as.n*3 <= len(as.a) {
		as.resize()
	}

	// 削除前の要素を返す
	return v
}

// backingArrayのサイズを現在の要素数の2倍に拡張する
func (as *ArrayStack) resize() {
	var newLen int
	if as.n == 0 {
		newLen = 1
	} else {
		newLen = as.n * 2
	}
	newA := make([]string, newLen)
	for i := 0; i < as.n; i++ {
		newA[i] = as.a[i]
	}
	as.a = newA
}
