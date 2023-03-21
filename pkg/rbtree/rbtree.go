package rbtree

import (
	"fmt"
	"strings"
)

type Index interface {
	// Compare : 和 index 比较, 小于/等于/大于 分别返回 -1,0,1
	Compare(index Index) int8
}

func less(i, j Index) bool {
	return i.Compare(j) < 0
}

// 性质1：每个节点要么是黑色，要么是红色。
// 性质2：根节点是黑色。
// 性质3：每个叶子节点（NIL）是黑色。
// 性质4：每个红色结点的两个子结点一定都是黑色。
// 性质5：任意一结点到每个叶子结点的路径都包含数量相同的黑结点。（保证这棵树尽量是平衡的。）
type RBTree struct {
	root *RBNode
}

type nodeDisplayBlock struct {
	parent  *nodeDisplayBlock
	display string
}

func (t *RBTree) Find(target Index) *RBNode {
	if t.root == nil {
		return nil
	}
	return t.findNode(t.root, target)
}

// Add 添加节点
func (t *RBTree) Add(n *RBNode) {
	if t.root == nil {
		t.root = n
		t.root.isblack = true
		return
	}
	t.addNode(t.root, n)
	t.fixup(n)
}

// Remove 移除节点
func (t *RBTree) Remove(target Index) {
	n := t.findNode(t.root, target)
	if n == nil {
		return
	}
	t.delNode(n)
}

// 递归添加节点, 保证有序, 不保证平衡
func (t *RBTree) addNode(p *RBNode, n *RBNode) {
	if less(p.value, n.value) {
		if p.right == nil {
			p.right = n
			n.parent = p
		} else {
			t.addNode(p.right, n)
		}
	} else {
		if p.left == nil {
			p.left = n
			n.parent = p
		} else {
			t.addNode(p.left, n)
		}
	}
}

// 添加节点后, 调整平衡
func (t *RBTree) fixup(n *RBNode) {
	p := n.parent
	if p == nil {
		n.isblack = true
		return
	}
	if p.isblack {
		return
	}
	if p.parent == nil {
		p.isblack = true
		return
	}
	if p.parent.left == p {
		uncle := p.parent.right
		if uncle != nil && !uncle.isblack {
			// case1:
			//  Parent 和 uncle 都为红色节点, 重新着色即可
			//
			//     GP[B]                   GP[R]
			//    /    \                   /   \
			//   P[R]  uncle[R]  =>      P[B]  uncle[B]
			//  /                        /
			// n[R]                     n[R]
			n.parent.isblack = true
			uncle.isblack = true
			p.parent.isblack = false
			t.fixup(p.parent)
			return
		} else {
			if n == p.right {
				// case2:
				//  parent 为红色, uncle 为黑色, 对 parent 左旋转
				//     GP[B]                   GP[B]
				//    /    \                   /   \
				//   P[R]  uncle[B]  =>     n[R]  uncle[B]
				//     \                     /
				//      n[R]               P[R]
				t.leftRotate(p)
				// 转到 case3
				t.fixup(p)
				return
			} else {
				// case3:
				//  parent 为红色, uncle 为黑色, 对 grandparent 右旋转
				//  parent 改成黑, grandparent 改为红色
				//      GP[B]                     P(B)
				//      /   \                     /   \
				//   P[R]  uncle[B]  =>         n[R]  GP[R]
				//    / \                        /     /   \
				//  n[R] X[B]                 Y[B]   X[B] uncle[B]
				//  /
				// Y[B]
				p.isblack = true
				p.parent.isblack = false
				t.rightRotate(p.parent)
				return
			}
		}
	} else {
		uncle := p.parent.left
		// 与上面方向相反
		if uncle != nil && !uncle.isblack {
			n.parent.isblack = true
			uncle.isblack = true
			p.parent.isblack = false
			t.fixup(p.parent)
			return
		} else {
			if n == p.left {
				t.rightRotate(p)
				t.fixup(p)
				return
			} else {
				p.isblack = true
				p.parent.isblack = false
				t.leftRotate(p.parent)
				return
			}
		}
	}
}

func (t *RBTree) delNode(n *RBNode) {
	if n.left != nil && n.right != nil {
		// 删除节点 n 有左右两个子节点, 找到后继节点 next 代替 n, 再删除 next 节点
		next := t.successor(n)
		n.value = next.value
		t.delNode(next)
		return
	}
	if n.left == nil && n.right == nil {
		// n 没有孩子
		if !n.isblack {
			//       p[]              p[]
			//      /   \      =>     /
			//  bro[R]  n[R]        bro[R]
			if n.parent.left == n {
				n.parent.left = nil
			} else {
				n.parent.right = nil
			}
			n.parent = nil
			return
		} else {
			// n 为黑色节点, 删除 n 会导致 n 所在路径黑色数减一, 产生不平衡
			// 所以, 先促使 n 所在路径的黑色点数加一, 然后再删除 n
			t.addNodeWeight(n)
			if n.parent.left == n {
				n.parent.left = nil
			} else {
				n.parent.right = nil
			}
			n.parent = nil
			return
		}
	} else if n.left != nil {
		// n 有一个左孩子
		if n.isblack {
			//         p[]                   p[]
			//         /  \                 /  \
			//       n[B]  bro[B]   =>   x[B]  bro[B]
			//       /
			//      x[R]
			if n.left.isblack {
				panic("black node's only son must be red")
			}
			n.value = n.left.value
			n.left.parent = nil
			n.left = nil
			return
		} else {
			// n 为红节点, n.left 无论黑红, 都破坏红黑树规则, 这种情况不存在
			panic("red node should not has only one child")
		}
	} else {
		// n 有一个右孩子
		if n.isblack {
			if n.right.isblack {
				panic("black node's only son must be red")
			}
			n.value = n.right.value
			n.right.parent = nil
			n.right = nil
			return
		} else {
			// n 为红节点, n.right 无论黑红, 都破坏红黑树规则, 这种情况不存在
			panic("red node should not has only one child")
		}
	}
	fmt.Println("done")
}

// 5 种 case, 最后转到 case2 和 case5.
// 增加 n 所在路径的权重
func (t *RBTree) addNodeWeight(n *RBNode) {
	var bro *RBNode
	broIsRight := false
	if n == n.parent.left {
		bro = n.parent.right
		broIsRight = true
	} else {
		bro = n.parent.left
	}

	if !bro.isblack {
		// case1:
		//    p[B]                Bro[B]
		//    /  \       =>        / \
		//  n[B] bro[R]         p[R]  Y[B]
		//	     /  \            /  \
		//     X[B] Y[B]       n[B] X[B]
		if !n.parent.isblack {
			panic("parent must be black")
		}
		if broIsRight {
			t.leftRotate(n.parent)
		} else {
			t.rightRotate(n.parent)
		}
		n.parent.isblack = false
		bro.isblack = true
		// 转到 case5
		t.addNodeWeight(n)
		return
	} else {
		if (bro.right != nil && !bro.right.isblack && broIsRight) ||
			(bro.left != nil && !bro.left.isblack && !broIsRight) {
			// case2: bro 有一个与其方向一致的红色子节点
			//         P[]                  Bro[]
			//        /  \                   / \
			//      N[B] Bro[B]   =>      P[B] X[B]
			//             \               /
			//             X[R]          N[B]
			if broIsRight {
				t.leftRotate(n.parent)
				bro.right.isblack = true
			} else {
				t.rightRotate(n.parent)
				bro.left.isblack = true
			}
			bro.isblack = n.parent.isblack
			n.parent.isblack = true
			return
		} else if (bro.right != nil && !bro.right.isblack) ||
			(bro.left != nil && !bro.left.isblack) {
			// case3: bro 有红色子节点, 但是不与 bro 方向一致, 则先调整成方向一致即 case1 的情况
			//         P[]                   P[]
			//        /  \                   / \
			//      N[B] Bro[B]   =>      N[B]  X[B]
			//            /                      \
			//          X[R]                     Bro[R]
			if broIsRight {
				bro.left.isblack = true
				bro.isblack = false
				t.rightRotate(bro)
			} else {
				bro.right.isblack = true
				bro.isblack = false
				t.leftRotate(bro)
			}
			// 转到 case2
			t.addNodeWeight(n)
			return
		} else {
			// bro 没有红色子节点
			if bro.parent.isblack {
				// case4: n,bro,p 都为黑节点
				//    p[B]                p[B]
				//    /  \       =>      /  \
				//  n[B] bro[B]        n[B]  bro[R]
				//
				//  或者
				//
				//    p[B]                p[B]
				//    /  \       =>      /  \
				//  n[B] bro[B]        n[B]  bro[R]
				//        / \                 / \
				//      X[B] Y[B]           X[B] Y[B]
				bro.isblack = false
				if n.parent.parent != nil {
					// 如果 p 不是 root, 继续上溯
					t.addNodeWeight(n.parent)
				}
				return
			} else {
				// case5
				//    p[R]                p[B]
				//    /  \       =>      /  \
				//  n[B] bro[B]        n[B]  bro[R]
				n.parent.isblack = true
				bro.isblack = false
				return
			}
		}
	}
	fmt.Print(1)
}

// 左旋转
func (t *RBTree) leftRotate(x *RBNode) {
	y := x.right
	p := x.parent

	x.right = y.left
	if x.right != nil {
		x.right.parent = x
	}
	y.left = x
	x.parent = y

	y.parent = p
	if p == nil {
		t.root = y
	} else {
		if p.left == x {
			p.left = y
		} else {
			p.right = y
		}
	}
}

// 右旋转
func (t *RBTree) rightRotate(y *RBNode) {
	x := y.left
	p := y.parent

	y.left = x.right
	if y.left != nil {
		y.left.parent = y
	}

	x.right = y
	y.parent = x
	x.parent = p
	if p == nil {
		t.root = x
	} else {
		if p.left == y {
			p.left = x
		} else {
			p.right = x
		}
	}
}

// 返回 n 的后继节点
func (t *RBTree) successor(n *RBNode) *RBNode {
	current := n.right
	if current == nil {
		return nil
	}
	for {
		next := current.left
		if next == nil {
			break
		}
		current = next
	}
	return current
}

// 递归查找节点
func (t *RBTree) findNode(n *RBNode, target Index) *RBNode {
	if n == nil {
		return nil
	}
	k := n.value.Compare(target)
	if k == 0 {
		return n
	}
	if k < 0 {
		return t.findNode(n.right, target)
	}

	return t.findNode(n.left, target)
}

// 检验树是否满足红黑树约束
func (t *RBTree) valid() bool {
	if t.root == nil {
		return true
	}

	if !t.root.isblack {
		return false
	}

	maxDepth := -1
	return !abortInvalidNode(t.root, 0, &maxDepth)
}

// 如果深度错误 抛出错误
func abortInvalidDepth(maxDepth *int, depth int) bool {
	if *maxDepth == -1 {
		*maxDepth = depth
		return false
	}
	return *maxDepth != depth
}

// 递归检测节点是否合法
func abortInvalidNode(n *RBNode, depth int, maxDepth *int) bool {
	if n.parent == nil && depth > 0 {
		return true
	}
	if n.isblack {
		depth += 1
	} else {
		if n.left != nil && !n.left.isblack || n.right != nil && !n.right.isblack {
			// 相邻红色报错
			return true
		}
	}

	// 遇到 nil 节点, 有路径走到底了
	if (n.left == nil || n.right == nil) && abortInvalidDepth(maxDepth, depth) {
		return true
	}

	if n.left != nil && abortInvalidNode(n.left, depth, maxDepth) {
		return true
	}

	if n.right != nil && abortInvalidNode(n.right, depth, maxDepth) {
		return true
	}

	return false
}

// 获取树的高度
func (t *RBTree) getHeight(n *RBNode) int {
	if n == nil {
		return 0
	}

	leftHeight := t.getHeight(n.left)
	rightHeight := t.getHeight(n.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// 打印树的信息
func (t *RBTree) print() {
	blocks := map[int][]*nodeDisplayBlock{}
	t.collectDisplayBlocks(t.root, nil, 1, blocks)
	height := t.getHeight(t.root)
	blockOffsets := map[*nodeDisplayBlock]int{}
	for i := 1; i <= height+1; i++ {
		offset := 0
		ls := blocks[i]
		for _, v := range ls {
			parentOffset := blockOffsets[v.parent]
			// 和父层 nodeDisplayBlock 对齐
			for j := offset; j < parentOffset; j++ {
				fmt.Print(" ")
				offset += 1
			}
			fmt.Print(v.display)
			blockOffsets[v] = offset
			offset += len(v.display)
		}
		fmt.Println("")
	}
	fmt.Println("")
}

// 收集每层节点的 nodeDisplayBlock
func (t *RBTree) collectDisplayBlocks(n *RBNode, parent *nodeDisplayBlock, level int, displayBlocks map[int][]*nodeDisplayBlock) int {
	sb := strings.Builder{}
	s := ""
	w := 0
	offset := 0
	l := &nodeDisplayBlock{parent: parent}
	if n == nil {
		s = "<nil>"
	} else {
		s = n.String()
		leftWidth := t.collectDisplayBlocks(n.left, l, level+1, displayBlocks)
		rightWidth := t.collectDisplayBlocks(n.right, l, level+1, displayBlocks)
		w = w + leftWidth + rightWidth
		// center of left and right child
		offset = (leftWidth+rightWidth/2+leftWidth/2)/2 - len(s)/2
	}

	if w < len(s) {
		// margin 2
		w = len(s) + 2
		offset = (w - len(s)) >> 1
	}

	for i := 0; i < offset; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString(s)
	for i := 0; i < w-len(s)-offset; i++ {
		sb.WriteString(" ")
	}

	l.display = sb.String()
	arr, ok := displayBlocks[level]
	if ok {
		displayBlocks[level] = append(arr, l)
	} else {
		displayBlocks[level] = []*nodeDisplayBlock{l}
	}
	return w
}

type RBNode struct {
	value   Index
	parent  *RBNode
	left    *RBNode
	right   *RBNode
	isblack bool
}

func (n *RBNode) String() string {
	s := "R"
	if n.isblack {
		s = "B"
	}
	return fmt.Sprintf("%v[%s]", n.value, s)
}
