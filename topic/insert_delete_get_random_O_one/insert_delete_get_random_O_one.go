package insert_delete_get_random_O_one

import "math/rand"

// Your RandomizedSet object will be instantiated and called as such:
// obj := Constructor();
// param_1 := obj.Insert(val);
// param_2 := obj.Remove(val);
// param_3 := obj.GetRandom();

type RandomizedSet struct {
	set map[int]struct{}
}

func Constructor() RandomizedSet {
	return RandomizedSet{set: make(map[int]struct{})}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.set[val]; !ok {
		s.set[val] = struct{}{}
		return true
	}
	return false
}

func (s *RandomizedSet) Remove(val int) bool {
	if _, ok := s.set[val]; ok {
		delete(s.set, val)
		return true
	}
	return false
}

func (s *RandomizedSet) GetRandom() int {
	randNumber := rand.Intn(len(s.set))
	count := 0

	for val := range s.set {
		if count == randNumber {
			return val
		}
		count++
	}

	return -1
}
