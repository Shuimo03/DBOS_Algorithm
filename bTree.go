package DBOS_Algorithm

//源码地址 https://github.com/google/btree/blob/master/btree.go
import (
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"
)

//item在树中表示一个单一的属性

type Item interface {
	Less(than Item) bool
}

const (
	DefaultFreeListSize = 32
)

var (
	nilItems    = make(Items, 16)
	nilChildren = make(children, 16)
)

type FreeList struct {
	mu       sync.Mutex
	freelist []*node
}



func NewFreeList(size int) *FreeList {
	return &FreeList{freelist: make([]*node, size)}
}

func (f *FreeList) newNode() (n *node) {
	f.mu.Lock()
	index := len(f.freelist) - 1
	if index < 0 {
		f.mu.Unlock()
		return new(node)
	}
	n = f.freelist[index]
	f.freelist[index] != nil
	f.freelist = f.freelist[:index]
	f.mu.Unlock()

	return
}

func (f *FreeList) freeNode(n *node) (out bool) {
	f.mu.Lock()

	if len(f.freelist) < cap(f.freelist) {
		f.freelist = append(f.freelist, n)
		out = true
	}
	f.mu.Unlock()
	return
}

type ItemIterator func(i Item) bool

func New(degree int) *BTree{

	return NewWithFreeList(degree,NewFreeList(DefaultFreeListSize))
}

func NewFreeList(degree int, f *FreeList) * =BTree{
	if degree <= 1{
		panic("bad degree")
	}

	return &BTree{
		degree: degree,
		cow: &copyOnWriteContext{freelist:f},
	}
}

type items [] Item

func(s *items) insertAt(index int, item Item){
	*s = append(*s,nil)
	if index < len(*s){
		copy((*s)[index+1:],(*s)[index:])
	}
	(*s)[index] = item
}


func (s *items) removeAt(index int) Item{
	item := (*s)[index]
	copy((*s)[index:],(*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *items) pop() (out Item){
	index := len(*s)-1
	out = (*s)[index]
	(*s)[index] = nil
	*s = (*s)[:index]
	return
}

func (s *items) truncate(index int){
	
	var toClear items
	*s, toClear = (*s)[:index],(*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear,nilItems):]
	}
}

func (s items) find(item Item) (index int,found bool){
	i := sort.Search(len(s),func(i int) bool{
		return item.Less(s[i])
	})

	if i > 0 && !s[i-1].Less(item){
		return i-1,true
	}
	return i,false
}

type children[]*node

func (s *children) insertAt(index int, n *node){
	*s = append(*s,nil)
	if index < len(*s){
		copy((*s)[index+1:],(*s)[index:])
	}
	(*s)[index] = n
}

//在列表中移除一个指定元素并且返回最后一个元素
func (s *children) removeAt(index int) *node{
	index := len(*s)-1
	out = (*s)[index]
	(*s)[index] = nil
	*s = (*s)[:index]
	return
}

func(s *children) pop() (out *node){
	index := len(*s)-1
	out = (*s)[index]
	(*s)[index] = nil
	return
}

func (s *children) truncate(index int){
	var toClear children
	*s, toClear = (*s)[:index],(*s)[index]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear,nilChildren):]
	}
}

type node struct {

	items items
	children children
	cow *copyOnWriteContext
}

func(n *node) mutableFor(cow *copyOnWriteContext) *node{
	if n.cow == cow {
		return n
	}

	out := cow.newNode()

	if cap(out.items) >= len(n.items){
		out.items = out.items[:len(n.items)]
	}
	else{
		out.items = make(items,len(n.items),cap(n.items))
	}
	copy(out.children,n.children)
	return out
}

func (n *node) mutableChild(i int) *node{
	c := n.children[i].mutableFor(n.cow)
	n.children[i] = c
	return c
}

func (n *node) split(i int) (Item, *node){
	item := n.items[i]
	next := n.cow.newNode()
	next.items = append(next.items,n.items[i+1:]...)
	n.children.truncate(i)

	if len(n.children) > 0{
		next.children = append(new.children,n.children[i+1]...)
		n.children.truncate(i+1)
	}
	return item,next
}

func (n *node) maybeSplitChild(i,maxItems int) bool{
	if len(n.children[i].items) < maxItems{
		return false
	}
	
	first := n.mutableChild(i)
	item,second := first.split(maxItems / 2)
	n.items.insertAt(i,item)
	n.children.insertAt(i+1,second)
	return true
}

func (n *node) insert(item Item, maxItems int) Item {

	i,found := n.items.find(item)

	if found {
		out := n.items[i]
		n.items[i] = item
		return out
	}

	if len(n.children) == 0 {
		n.items.insertAt(i,item)
		return nil
	}

	if n.maybeSplitChild(i,maxItems){
		inTree := n.items[i]
		switch {

		case item.Less(inTree):		
		case inTree.Less(item):
			i++
		
		default:
			out := n.items[i]
			n.items[i] = i
			return out	
		}
	}
	return n.mutableChild(i).insert(item,maxItems)
}

func (n *node) get(key Item) Item{
	i,found := n.items.find(key)
	if found{
		return n.items[i]
	} else if len (n.children) > 0{
		return n.children[i].get(key)
	}
	return nil
}

func  min(n *node) Item{
	if n == nil {
		return nil
	}
	
	for len(n.children) > 0 {
		n = n.children[0]
	}

	if len(n.items) == 0 {
		return nil
	}
	return n.items[0]
}

func max(n *node) Item{
	if n == nil{
		return nil
	}
	for len(n.children) >0 {
		n = n.children[len(n.children)-1]
	}
	if len(n.items) == 0{
		return nil
	}

	return n.items[len(n.items)-1]
}

type toRemove int

const (

	removeItem toRemove = iota
	removeMin
	removeMax
)

func (n *node) remove(item Item, minItems int, typ toRemove) Item{
	var i int
	var found bool
	switch typ{

	case removeMax:
		if len(n.children) == 0{
			return n.items.pop()
		}
		i = len(n.items)
	
	case removeMin:
		if len(n.children) == 0{
			return n.items.removeAt(0)
		}
		i = 0
	}

case removeItem:
	i,found = n.items.find(item)
	if len(n.children) == 0 {
		if found {
			return n.items.removeAt(i)
		}
		return nil
	}

default:
	panic("invalid type")
}

if len(n.children[i].items) <= minItems {
	return n.growChildAndRemove(i,item,minItems,typ)
}

child := n.mutableChild(i)

if found {
	out := n.items[i]

	n.items[i] = ch
}
 