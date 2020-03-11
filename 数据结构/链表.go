package main

import "log"

// https://studygolang.com/articles/13795
type Object interface {
}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	HandNode *Node
}

// 增加头节点
func (this *List) Add(data Object) {
	node := &Node{data, nil}

	if !this.IsExists() {
		this.HandNode = node
	} else {
		pre := this.HandNode
		node.Next = pre
		this.HandNode = node
	}
}

// 追加节点
func (this *List) Append(data Object) {
	node := &Node{data, nil}

	if !this.IsExists() { // 不存在首节点
		this.HandNode = node
	} else { // 存在首节点
		pre := this.HandNode

		for nil != pre.Next { // 判断节点是否有子阶段，一直循环到无子节点（最后一个节点）
			pre = pre.Next
		}

		pre.Next = node
	}
}

// 在指定位置插入元素
func (this *List) Insert(index int, data Object) {
	node := &Node{data, nil}

	if 0 >= index { // index 小于等于0，从头部插入
		this.Add(node)
	} else if index >= this.Length() { // 大于等于链表长度，从尾部插入
		this.Append(node)
	} else {
		var count int

		pre := this.HandNode

		for count < index-1 {
			pre = pre.Next
			count ++
		}

		node.Next = pre.Next
		pre.Next = node
	}
}

// 删除指定位置链表
func (this *List) DelIndex(index int) {
	pre := this.HandNode
	if 0 >= index { // index 小于等于0，删除头结点
		this.HandNode = pre.Next
	} else if index >= this.Length() {
		log.Println("index 超出链表长度")
	} else {
		var count int
		for count != (index-1) && nil != pre.Next {
			count++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

// 判断头链表是否存在
func (this *List) IsExists() bool {
	if nil == this.HandNode {
		return false
	}

	return true
}

// 返回链表的长度
func (this *List) Length() int {
	if !this.IsExists() {
		return 0
	}

	var count int

	curr := this.HandNode

	for nil != curr {
		count ++
		curr = curr.Next
	}

	return count
}

// 判断链表是否包含某个元素
func (this *List) Contain(data Object) bool {
	pre := this.HandNode

	for nil != pre {
		if data == pre.Data {
			return true
		}
		pre = pre.Next
	}

	return false
}

func (this *List) Print() {
	if this.IsExists() {
		pre := this.HandNode
		for {
			log.Println("node:", pre.Data)
			if nil == pre.Next {
				break
			}
			pre = pre.Next
		}
	}
}

func main() {
	list := new(List)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Add(0)
	list.Insert(1, 3)
	list.DelIndex(3)
	list.Contain(2)
	list.Print()
}
