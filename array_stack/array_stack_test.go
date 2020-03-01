package array_stack

import (
	"fmt"
	"testing"
)

func Example() {
	as := NewArrayStack(5)
	as.Add(3, "hoge")
	v := as.Get(3)
	fmt.Println(v)

	_ = as.Set(3, "fuge")
	v = as.Get(3)
	fmt.Println(v)

	as.Add(3, "piyo")
	v = as.Get(3)
	fmt.Println(v)

	// Output:
	// hoge
	// fuge
	// piyo
}

func TestArrayStack_Size(t *testing.T) {
	as := NewArrayStack(5)
	v := as.Size()
	if v != 0 {
		t.Errorf("expected 0 but god %d\n", v)
	}
}

func TestArrayStack_Get(t *testing.T) {
	as := NewArrayStack(5)
	v := as.Get(0)

	if v != "" {
		t.Errorf("expected blank but got %s\n", v)
	}
}

// Addの際に適切にresizeされることをテスト
func TestArrayStack_AddAndResize(t *testing.T) {
	// 大きさ2の配列で初期化
	as := NewArrayStack(2)
	as.Add(0, "hoge")
	as.Add(0, "fuge")
	// 配列が満タンなので、2x2の4に拡張される
	as.Add(0, "piyo")

	expected := 4
	got := len(as.a)
	if got != expected {
		t.Errorf("expected %v but got %v\n", expected, got)
	}
}

func TestArrayStack_Set(t *testing.T) {
	as := NewArrayStack(5)
	as.Add(3, "hoge")

	old := as.Set(3, "fuge")
	if old != "hoge" {
		t.Errorf("expected: hoge but got %v\n", old)
	}
	v := as.Get(3)

	if v != "fuge" {
		t.Errorf("expected: fuge but got %v\n", v)
	}

	as.Add(3, "piyo")
	v = as.Get(3)
	if v != "piyo" {
		t.Errorf("expected: piyo but got %v\n", v)
	}
	v = as.Get(4)
	if v != "fuge" {
		t.Errorf("expected: fuge but got %v\n", v)
	}
	size := as.Size()
	if size != 2 {
		t.Errorf("expected 2 but got %v\n", size)
	}
}

func TestArrayStack_Remove(t *testing.T) {
	stack := NewArrayStack(3)
	stack.Add(0, "a")
	stack.Add(1, "b")
	stack.Add(2, "c")
	got := stack.Remove(1)
	expected := "b"
	if got != expected {
		t.Errorf("expected: %v but got %v\n", expected, got)
	}

	got = stack.Get(0)
	expected = "a"
	if got != expected {
		t.Errorf("expected: %v but got %v\n", expected, got)
	}

	got = stack.Get(1)
	expected = "c"
	if got != expected {
		t.Errorf("expected: %v but got %v\n", expected, got)
	}

	got = stack.Get(2)
	expected = ""
	if got != expected {
		t.Errorf("expected: %v but got %v\n", expected, got)
	}

	size := stack.Size()
	if stack.Size() != size {
		t.Errorf("expected size: 2 but got %v\n", size)
	}
}

func TestArrayStack_RemoveAndResize(t *testing.T) {
	stack := NewArrayStack(10)
	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("hoge_%d", i)
		stack.Add(i, s)
	}

	for i := 0; i < 7; i++ {
		stack.Remove(i)
	}

	length := len(stack.a)
	if length != 6 {
		t.Errorf("expected 6 but got %v\n", length)
	}
}
