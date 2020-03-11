package level_2

/*
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]
 

限制：

1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5

可能有人看不懂题，本小白说一下我自己的理解： 输入的两个列表，其元素分别是执行的函数与给的值。 以输入1为例：

1、  “MaxQueue”  生成队列
        []                    传入的值为空
        函数无输出，故输出为null
2、  "push_back"  向队列传入元素
        [1]                  传入的值为1
        此时队列中只有一个1
        函数无输出，故输出为null
3、  "push_back"  向队列传入元素
        [2]                  传入的值为2
        此时队列为    --进-> [2,1] --出->
        输出为null
4、   "max_value" 求队列中的最大值
        []                    传入的值为空
        输出为2
5、   "pop_front"   删除队列头部元素
        []                    传入的值为空
        函数pop_front还要输出删除元素的值，故输出为1
6、   "max_value" 求队列中的最大值
        []                    传入的值为空
         此时队列为[2]，所以最大值输出为2。
希望能起到一点作用，也感谢大佬们能指教我描述不准确的的地方。（不知为何，觉得自己写的很别扭。。）
（编辑：我打字的时候是分行的，但预览又成一整段了。。没办法只能用代码块了，看着更别扭了- -，但总比一整段清晰一点。。。）
*/

type Node struct {
	Value int
	Next  *Node
}

type MaxQueue struct {
	HeadNode *Node
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	// 判断头节点是否存在
	if nil == this.HeadNode { // 不存在新增
		return -1
	}

	max := this.HeadNode.Value
	next := this.HeadNode

	for nil != next.Next {
		next = next.Next

		if max < next.Value {
			max = next.Value
		}
	}

	return max
}

func (this *MaxQueue) Push_back(value int) {
	node := &(Node{value, nil})
	// 判断头节点是否存在
	if nil == this.HeadNode { // 不存在新增
		this.HeadNode = node
		return
	}

	next := this.HeadNode

	for nil != next.Next {
		next = next.Next
	}

	next.Next = node
}

func (this *MaxQueue) Pop_front() int {
	// 判断头节点是否存在
	if nil == this.HeadNode { // 不存在新增
		return -1
	}

	front := this.HeadNode.Value

	if nil == this.HeadNode.Next {
		this.HeadNode = nil
	} else {
		this.HeadNode = this.HeadNode.Next
	}

	return front
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
