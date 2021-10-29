package linkedlist

import (
  "testing"
)

func TestLinkedList_Create(t *testing.T) {
  ll := Create()
  if ll.root != nil {
    t.Errorf("Root must not be set")
  }
}

func TestLinkedList_Add(t *testing.T) {
  ll := Create()
  ll = Add(ll, 4)
  if ll.root == nil {
    t.Errorf("Root must not be nil")
  }
  if ll.root.value != 4 {
    t.Errorf("Root value must be 4")
  }
}

func TestLinkedList_Add2(t *testing.T) {
  type value struct {
    array []rune
  }
  tests := []struct {
    name string
    value value
    want value
  }{ 
     {
        name: "abc", 
        value: value{[]rune{'a', 'b', 'c'}}, 
        want: value{[]rune{'a', 'b', 'c'}},
     },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      ll := Create()
      for _, v := range tt.value.array {
        ll = Add(ll, v)
      }
      var x *Node
      x = ll.root
      for _, r := range tt.want.array {
        if x == nil {
          t.Errorf("List was truncated")
        }
        if x.value != r {
          t.Errorf("Value does not agree %v != %v", x.value, r)
        }
        x = x.next
      }
    })
  }
}

func TestLinkedList_Push(t *testing.T) {
  type value struct {
    array []rune
  }
  tests := []struct {
    name string
    value value
    want value
  }{ 
     {
        name: "abc", 
        value: value{[]rune{'a', 'b', 'c'}}, 
        want: value{[]rune{'c', 'b', 'a'}},
     },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      ll := Create()
      for _, v := range tt.value.array {
        ll = Push(ll, v)
      }
      var x *Node
      x = ll.root
      for _, r := range tt.want.array {
        if x == nil {
          t.Errorf("List was truncated")
        }
        if x.value != r {
          t.Errorf("Value does not agree %v != %v", x.value, r)
        }
        x = x.next
      }
    })
  }
}

func TestLinkedList_Pop(t *testing.T) {
  type value struct {
    array []rune
  }
  tests := []struct {
    name string
    value value
    want value
  }{ 
     {
        name: "abc", 
        value: value{[]rune{'a', 'b', 'c'}}, 
        want: value{[]rune{'a', 'b', 'c'}},
     },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      ll := Create()
      for _, v := range tt.value.array {
        ll = Add(ll, v)
      }
      for _, r := range tt.want.array {
        v := Pop(ll)
        if v != r {
          t.Errorf("Value does not agree %v != %v", v, r)
        }
      }
    })
  }
}
