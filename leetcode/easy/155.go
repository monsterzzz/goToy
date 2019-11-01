package easy

/*

设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.



*/
type MinStack struct {
	Val    []int
	helper []int // 辅助栈
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(x int) {
	m.Val = append(m.Val, x)
	if len(m.helper) == 0 { // 如果辅助栈为空或者push的元素小于最后一个元素才加入
		m.helper = append(m.helper, x)
	} else {
		if m.helper[len(m.helper)-1] > x {
			m.helper = []int{x}
		} else if m.helper[len(m.helper)-1] == x {
			m.helper = append(m.helper, x)
		}
	}

}

func (m *MinStack) Pop() {
	m.Val = m.Val[:len(m.Val)-1]
	if m.Val[len(m.Val)-1] == m.helper[len(m.helper)-1] { // pop出的元素 与 辅助栈最后一个元素 相同 辅助栈才pop
		m.helper = m.helper[:len(m.helper)-1]
	}
}

func (m *MinStack) Top() int {
	return m.Val[len(m.Val)-1]
}

func (m *MinStack) GetMin() int {
	return m.helper[len(m.helper)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
