package main

import (
	"bytes"
	"testing"
)

func TestFibonacciSeriesAndNumberAgree(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for n, w := range want {
		if got := Series(n); got != w {
			t.Errorf("Series(%d) = %d, want %d", n, got, w)
		}
		if got := FibonacciNumber(n); got != w {
			t.Errorf("FibonacciNumber(%d) = %d, want %d", n, got, w)
		}
	}
}

func TestLookSay(t *testing.T) {
	tests := []struct{ in, want string }{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
	}
	for _, tt := range tests {
		if got := look_say(tt.in); got != tt.want {
			t.Errorf("look_say(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestThueMorseSequence(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteByte('0')
	ThueMorseSequence(&buf)
	got := buf.String()
	want := "01"
	if got != want {
		t.Fatalf("after one extension = %q, want %q", got, want)
	}
	ThueMorseSequence(&buf)
	if buf.String() != "0110" {
		t.Errorf("after two extensions = %q, want 0110", buf.String())
	}
}

func TestDictionaryOperations(t *testing.T) {
	dict := &Dictionary{}
	dict.Put("go", "golang")
	dict.Put("py", "python")

	if !dict.Contains("go") || dict.Contains("java") {
		t.Error("Contains returned wrong results")
	}
	if got := dict.Find("py"); got != "python" {
		t.Errorf("Find(py) = %q, want python", got)
	}
	if n := dict.NumberOfElements(); n != 2 {
		t.Errorf("NumberOfElements = %d, want 2", n)
	}
	keys := dict.GetKeys()
	if len(keys) != 2 {
		t.Errorf("GetKeys length = %d, want 2", len(keys))
	}
	values := dict.GetValues()
	if len(values) != 2 {
		t.Errorf("GetValues length = %d, want 2", len(values))
	}
	if ok := dict.Remove("go"); !ok || dict.Contains("go") {
		t.Error("Remove failed to delete key")
	}
	dict.Reset()
	if n := dict.NumberOfElements(); n != 0 {
		t.Errorf("after Reset NumberOfElements = %d, want 0", n)
	}
}

func TestTreeSetInsertSearchDelete(t *testing.T) {
	treeset := &TreeSet{}
	treeset.bst = new(BinarySearchTree)

	n1 := TreeNode{key: 1, value: 10}
	n2 := TreeNode{key: 2, value: 20}
	treeset.InsertTreeNode(n1, n2)

	if !treeset.Search(n1) {
		t.Error("Search should find inserted node")
	}
	treeset.Delete(n1)
	if treeset.Search(n1) {
		t.Error("Search should not find deleted node")
	}
	if !treeset.Search(n2) {
		t.Error("other node should remain after delete")
	}
}

func TestBSTRemoveRootNode(t *testing.T) {
	tree := new(BinarySearchTree)
	tree.InsertElement(1, 10)
	tree.InsertElement(2, 20)

	tree.RemoveNode(1)

	if tree.SearchNode(1) {
		t.Error("root node still searchable after removal")
	}
	if !tree.SearchNode(2) {
		t.Error("remaining node lost after root removal")
	}
}

func TestChapter07MainSmoke(t *testing.T) {
	main()
}
