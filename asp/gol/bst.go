package gol

type Node struct {
	key   int
	par   *Node
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func NewBST() *BST {
	return &BST{root: nil}
}

// TREE-INSERT (T, z)
// 1   x = T.root          // node being compared with z
// 2   y = NIL             // y will be parent of z
// 3   while x != NIL      // descend until reaching a leaf
// 4     y = x
// 5     if z.key < x.key
// 6       x = x.left
// 7     else x = x.right
// 8   z.par = y           // found the location - insert z with parent y
// 9   if y == NIL
// 10    T.root = z        // tree T was empty
// 11  elseif z.key < y.key
// 12    y.left = z
// 13  else y.right = z

func (T *BST) Insert(key int) {
	var z = &Node{key: key}
	var x = T.root
	var y *Node = nil
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.par = y
	if y == nil {
		T.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
}

// TRANSPLANT (T, u, v)
// 1  if u.par == NIL
// 2    T.root = v
// 3  elseif u == u.par.left
// 4    u.p.left = v
// 5  else u.par.right = v
// 6  if v != NIL
// 7    v.par = u.par

func (T *BST) Transplant(u, v *Node) {
	if u.par == nil {
		T.root = v
	} else if u == u.par.left {
		u.par.left = v
	} else {
		u.par.right = v
	}
	if v != nil {
		v.par = u.par
	}
}

// TREE-DELETE (T, z)
// 1 if z.left == NIL
// 2   TRANSPLANT (T, z, z.right)        // replace z by its right child
// 3 elseif z.right == NIL
// 4   TRANSPLANT (T, z, z.left)         // replace z by its left child
// 5 else y = TREE-MINIMUM(z.right)      // y is z’s successor
// 6      if y != z.right                // is y farther down the tree?
// 7          TRANSPLANT (T, y, y.right) // replace y by its right child
// 8          y.right = z.right          // z’s right child becomes
// 9          y.right.par = y            // y’s right child
// 10     TRANSPLANT (T, z, y)           // replace z by its successor y
// 11     y.left = z.left                // and give z’s left child to y,
// 12     y.left.par = y                 // which had no left child

func (T *BST) Delete(value int) bool {
	var z = T.Search(value)
	if z == nil {
		return false
	}
	if z.left == nil {
		T.Transplant(z, z.right)
	} else if z.right == nil {
		T.Transplant(z, z.left)
	} else {
		var y = minimum(z.right)
		if y != z.right {
			T.Transplant(y, y.right)
			y.right = z.right
			y.right.par = y
		}
		T.Transplant(z, y)
		y.left = z.left
		y.left.par = y
	}
	return true
}

func (T *BST) Search(value int) *Node {
	var x = T.root
	for x != nil {
		if value == x.key {
			return x
		} else if value < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return nil
}

func minimum(z *Node) *Node {
	var x = z
	for x.left != nil {
		x = x.left
	}
	return x
}

// Print the tree in ascii with indentation
func (T *BST) Print(title string) {
	println(title)
	println("-------------")
	printNode(T.root, 0, false)
	println("-------------")
}

func printNode(node *Node, level int, left bool) {
	if node == nil {
		return
	}
	for i := 1; i < level; i++ {
		print("    ")
	}
	if level > 0 {
		if left {
			print("-L- ")
		} else {
			print("-R- ")
		}
	}
	print("(")
	print(node.key)
	print(")\n")
	printNode(node.left, level+1, true)
	printNode(node.right, level+1, false)
}
