package level_2

import "testing"

/*

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]

输入:
["MaxQueue","max_value","pop_front","pop_front","push_back","push_back","push_back","pop_front","push_back","pop_front"]
[[],[],[],[],[94],[16],[89],[],[22],[]]
输出
[null,-1,-1,-1,null,null,null,94,null,16]

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]

*/

func TestConstructor(t *testing.T) {
	maxQueue := Constructor()

	max := maxQueue.Max_value()
	t.Log(max)

	maxQueue.Push_back(94)
	maxQueue.Push_back(16)
	maxQueue.Push_back(89)

	front := maxQueue.Pop_front()
	t.Log(front)

	maxQueue.Push_back(22)

	front = maxQueue.Pop_front()
	t.Log(front)
}
