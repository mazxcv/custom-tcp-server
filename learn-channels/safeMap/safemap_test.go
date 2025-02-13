package safemap

import "testing"

func TestSafeMapInsertGet(t *testing.T) {
	m := New[int, int]()

	for i := 0; i < 1000; i++ {
		go func(i int) {
			m.Insert(i, i*2)
			val, err := m.Get(i)
			if err != nil {
				t.Error(err)
			}
			if val != i*2 {
				t.Errorf("expected %d, got %d", i*2, val)
			}
		}(i)
	}
}
