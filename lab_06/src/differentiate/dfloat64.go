package differentiate

// DFloat64 used to represent value and derivative existance.
type DFloat64 struct {
	value float64
	null  bool
}

func (n *DFloat64) retvalue() interface{} {
	if n.null {
		return nil
	}
	return n.value
}

func newVal(x float64) DFloat64 {
	return DFloat64{
		value: x,
		null:  false,
	}
}

func newNil() DFloat64 {
	return DFloat64{
		value: 0.,
		null:  true,
	}
}
