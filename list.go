package main

type List struct {
	rootNode *Node
	endNode *Node
}

func (l *List) add(value int) {
	tmp := &Node{
		Data: value,
		Next: nil,
	}

	if (l.rootNode == nil) {
		l.rootNode = tmp;
		l.endNode = tmp;
	} else {
		l.endNode.Next = tmp;
		l.endNode = tmp;
	}
}

func (l *List) find(value int) *Node {
	var tmp = l.rootNode

	for tmp != nil {
		if (tmp.Data == value) {
			return tmp;
		}

		tmp = tmp.Next;
	}

	return nil;
}

func (l *List) insert(key int, value int) *Node {
	var keyNode = l.find(key)

	if (keyNode != nil) {
		tmpNode := &Node{
			Data: value,
			Next: nil,
		}

		tmpNode.Next = keyNode.Next;
		keyNode.Next = tmpNode;

		return tmpNode;
	}

	return nil;
}

func (l *List) get(index int) int {
	tmp := l.rootNode;
	i := 0

	for tmp != nil {
		if (i == index) {
			break;
		}

		tmp = tmp.Next
		i++
	}

	return tmp.Data;
}

func (l *List) count() int {
	tmp := l.rootNode;
	i := 0

	for tmp != nil {
		i++
		tmp = tmp.Next
	}

	return i;
}

func (l *List) remove(value int) bool {
	node := l.find(value);

	if (node != nil) {
		if (node == l.rootNode) {
			l.rootNode = node.Next
		} else {
			tmp := l.rootNode
			for tmp != nil {
				if (tmp.Next == node) {
					break
				}

				tmp = tmp.Next
			}

			tmp.Next = node.Next
			if (tmp == l.endNode) {
				tmp = l.endNode
			}

			return true
		}
	}

	return false;
}