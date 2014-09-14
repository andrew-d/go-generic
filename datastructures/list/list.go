// Parts of this file are inspired by the container/list
// package in the Go source code.  The original license
// follows:
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import (
	"github.com/joeshaw/gengen/generic"
)

// A single node in the list.
type ListNode struct {
	// Doubly-linked list pointers
	next, prev *ListNode

	// The list to which this node belongs.
	list *List

	// The underlying value.
	Value generic.T
}

func (n *ListNode) Next() *ListNode {
	return n.next
}

func (n *ListNode) Prev() *ListNode {
	return n.prev
}

// A doubly-linked list.
type List struct {
	root  ListNode
	count int
}

// Create a new list and return it.
func New() *List {
	return new(List).Init()
}

// Return the last element of the list or nil.
func (l *List) Back() *ListNode {
	if l.count == 0 {
		return nil
	}
	return l.root.prev
}

// Return the first element of the list or nil.
func (l *List) Front() *ListNode {
	if l.count == 0 {
		return nil
	}
	return l.root.next
}

// Initialize / clear the list.
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.root.list = l

	l.count = 0
	return l
}

// Insert `node` after `at`, increment the list's count and return `node`.
func (l *List) insert(node, at *ListNode) *ListNode {
	next := at.next
	at.next = node
	node.prev = at
	node.next = next
	next.prev = node

	node.list = l
	l.count++

	return node
}

// Helpful wrapper around `insert` with a value.
func (l *List) insertValue(v generic.T, at *ListNode) *ListNode {
	return l.insert(&ListNode{Value: v}, at)
}

// Remove `n` from the list, decrement the count, and return `n`.
func (l *List) remove(n *ListNode) *ListNode {
	n.prev.next = n.next
	n.next.prev = n.prev

	// Avoid memory leaks
	n.prev = nil
	n.next = nil

	n.list = nil
	l.count--

	return n
}

// Insert the given value immediately after mark and return the new
// list node.  If mark is not an element of l, the list is not modified.
func (l *List) InsertAfter(v generic.T, mark *ListNode) *ListNode {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

// Insert the given value immediately before mark and return the new
// list node.  If mark is not an element of l, the list is not modified.
func (l *List) InsertBefore(v generic.T, mark *ListNode) *ListNode {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

// Return the number of items in the list.  The complexity of this
// operation is O(1).
func (l *List) Len() int {
	return l.count
}

// TODO: document
func (l *List) MoveAfter(n, mark *ListNode) {
	if n.list != l || n == mark || mark.list != l {
		return
	}
	l.insert(l.remove(n), mark)
}

// TODO: document
func (l *List) MoveBefore(n, mark *ListNode) {
	if n.list != l || n == mark || mark.list != l {
		return
	}
	l.insert(l.remove(n), mark.prev)
}

// TODO: document
func (l *List) MoveToBack(n *ListNode) {
	if n.list != l || l.root.prev == n {
		return
	}
	l.insert(l.remove(n), l.root.prev)
}

// TODO: document
func (l *List) MoveToFront(n *ListNode) {
	if n.list != l || l.root.prev == n {
		return
	}
	l.insert(l.remove(n), &l.root)
}

// TODO: document
func (l *List) PushBack(v generic.T) *ListNode {
	return l.insertValue(v, l.root.prev)
}

// TODO: document
func (l *List) PushBackList(other *List) {
	for i, n := other.Len(), other.Front(); i > 0; i, n = i-1, n.Next() {
		l.insertValue(n.Value, l.root.prev)
	}
}

// TODO: document
func (l *List) PushFront(v generic.T) *ListNode {
	return l.insertValue(v, &l.root)
}

// TODO: document
func (l *List) PushFrontList(other *List) {
	for i, n := other.Len(), other.Back(); i > 0; i, n = i-1, n.Prev() {
		l.insertValue(n.Value, &l.root)
	}
}

// TODO: document
func (l *List) Remove(n *ListNode) generic.T {
	if n.list == l {
		l.remove(n)
	}
	return n.Value
}
