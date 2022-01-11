package collection

type SliceList[E comparable] []E

// iterator

type sliceListIterator[E comparable] struct {
	index int
	sl    SliceList[E]
}

func (sli *sliceListIterator[E]) HasNext() bool {
	return sli.index < len(sli.sl)
}

func (sli *sliceListIterator[E]) Next() (res E) {
	res = sli.sl[sli.index]
	sli.index++
	return
}

func (s *SliceList[E]) Iterator() Iterator[E] {
	if s == nil {
		return &sliceListIterator[E]{0, nil}
	}

	return &sliceListIterator[E]{0, *s}
}

// collection

func (s *SliceList[E]) Add(e E) bool {
	if s == nil {
		*s = []E{e}
	} else {
		*s = append(*s, e)
	}
	return true
}

func (s *SliceList[E]) Clear() {
	if s == nil {
		return
	}

	*s = nil
}

func (s *SliceList[E]) Contains(e E) bool {
	if s == nil {
		return false
	}

	for _, f := range *s {
		if e == f {
			return true
		}
	}

	return false
}

func (s *SliceList[E]) IsEmpty() bool {
	if s == nil {
		return true
	}

	return len(*s) == 0
}

func (s *SliceList[E]) Remove(e E) bool {
	if s == nil {
		return false
	}

	es := []E(*s)

	for i, f := range es {
		if e == f {
			*s = append(es[:i], es[i+1:]...)
			return true
		}
	}

	return false
}

func (s *SliceList[E]) Size() int {
	if s == nil {
		return 0
	}

	return len(*s)
}

func (s *SliceList[E]) ToSlice() []E {
	if s == nil {
		return nil
	}

	res := make([]E, len(*s))
	copy(res, *s)
	return res
}

// list

func (s *SliceList[E]) AddAt(i int, e E) {
	es := []E(*s)
	*s = append(es[:i], append([]E{e}, es[i:]...)...)
}

func (s *SliceList[E]) GetAt(i int) E {
	return []E(*s)[i]
}

func (s *SliceList[E]) IndexOf(e E) int {
	if s == nil {
		return -1
	}

	for i, f := range *s {
		if e == f {
			return i
		}
	}

	return -1
}

func (s *SliceList[E]) RemoveAt(i int) E {
	es := []E(*s)
	res := es[i]
	*s = append(es[:i], es[i+1:]...)
	return res
}

func (s *SliceList[E]) SetAt(i int, e E) (res E) {
	if i == len(*s) {
		s.Add(e)
		return
	}

	es := []E(*s)
	res = es[i]
	es[i] = e
	return
}
