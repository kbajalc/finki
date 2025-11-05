package gol

import "testing"

func TestBST(t *testing.T) {
	bst := NewBST()
	values := []int{5, 3, 7, 2, 4, 6, 8, 1, 9}

	for _, v := range values {
		bst.Insert(v)
	}
	bst.Print("Insertion")

	bst.Delete(3)
	bst.Print("Delete 3 ...")
	if bst.Search(3) != nil {
		t.Errorf("Did not expect to find value 3 in BST after deletion")
	}

	bst.Delete(5)
	bst.Print("Delete 5 ...")
	if bst.Search(5) != nil {
		t.Errorf("Did not expect to find value 5 in BST after deletion")
	}
}
