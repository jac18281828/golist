package linkedlist

type LinkedList struct {
  root *Node
}

type Node struct {
  next *Node
  value interface{}
}

func Create() *LinkedList {
  root := &LinkedList{ 
    root: nil,
  }
  return root
}

func Add(list *LinkedList, value interface{}) *LinkedList {
  if list.root == nil {
    list.root = &Node{nil, value}
  } else {
    next := list.root
    for ; next.next != nil; next = next.next {}
    next.next = &Node{nil, value}
  }
  return list
}

func Push(list *LinkedList, value interface{}) *LinkedList {
  list.root = &Node{list.root, value}
  return list
}

func Pop(list *LinkedList) interface{} {
  value := list.root
  if list.root != nil {
    list.root = list.root.next
  }
  return value.value
}

func Get(list *LinkedList, dx int) interface{} {
  next := list.root
  i := 0
  for ; next != nil && i < dx ; i++ {
    next = next.next
  }
  if i == dx {
    return next.value
  }
  return nil
}