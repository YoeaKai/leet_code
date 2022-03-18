package find_median_from_data_stream

import "container/heap"

/**
 * MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func MedianTest() {
	obj := Constructor()
	obj.AddNum(12)
	obj.AddNum(10)
	obj.AddNum(13)
	obj.AddNum(11)
	obj.AddNum(5)
	obj.AddNum(15)
	obj.AddNum(1)
	obj.AddNum(11)
	obj.AddNum(6)
	obj.AddNum(17)
	obj.AddNum(14)
	obj.AddNum(8)
	obj.AddNum(17)
	obj.AddNum(6)
	obj.AddNum(4)
	obj.AddNum(16)
	obj.AddNum(8)
	obj.AddNum(10)
	obj.AddNum(2)
	obj.AddNum(12)
	obj.AddNum(0)
	param2 := obj.FindMedian()
	println(param2)
}

// Method 1
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	min := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return min
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	min := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return min
}

type MedianFinder struct {
	minHeap MinHeap
	maxHeap MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: MinHeap{},
		maxHeap: MaxHeap{},
	}
}

func (m *MedianFinder) AddNum(num int) {
	if m.maxHeap.Len() == 0 || num <= m.maxHeap[0] {
		heap.Push(&m.maxHeap, num)
	} else {
		heap.Push(&m.minHeap, num)
	}
	if m.maxHeap.Len() > m.minHeap.Len()+1 {
		maxLeft := heap.Pop(&m.maxHeap)
		heap.Push(&m.minHeap, maxLeft)
	}
	if m.minHeap.Len() > m.maxHeap.Len() {
		minRight := heap.Pop(&m.minHeap)
		heap.Push(&m.maxHeap, minRight)
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.minHeap.Len() == m.maxHeap.Len() {
		return float64(m.minHeap[0]+m.maxHeap[0]) / 2
	}
	return float64(m.maxHeap[0])
}

// Method 2
type MedianFinder2 struct {
	storage []int
}

func Constructor2() MedianFinder2 {
	return MedianFinder2{storage: []int{}}
}

func (m *MedianFinder2) AddNum2(num int) {
	head := 0
	high := len(m.storage) - 1

	for {
		mid := (high + head) / 2
		if mid < 0 {
			m.storage = append([]int{num}, m.storage...)
			return
		} else if mid >= len(m.storage) {
			m.storage = append(m.storage, num)
			return
		} else if num == m.storage[mid] {
			m.storage = append(m.storage[:mid+1], m.storage[mid:]...)
			m.storage[mid] = num
			return
		} else if num < m.storage[mid] {
			if mid == 0 || num > m.storage[mid-1] {
				m.storage = append(m.storage[:mid+1], m.storage[mid:]...)
				m.storage[mid] = num
				return
			}
			high = mid - 1
		} else {
			if mid == len(m.storage)-1 {
				m.storage = append(m.storage, num)
				return
			}
			head = mid + 1
		}
	}
}

func (m *MedianFinder2) FindMedian2() float64 {
	if len(m.storage)%2 == 0 {
		return float64(m.storage[len(m.storage)/2-1]+m.storage[len(m.storage)/2]) / 2
	}
	return float64(m.storage[len(m.storage)/2])
}
