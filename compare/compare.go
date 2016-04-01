package compare

import "strings"

// Compare whether v1 is greater than v2.
// Return a positive integer if greater, 0 if equal, a negative if less.
//
// v1 and v2 may be a byte, rune, int, uint, int8, int16, int32, int64,
// uint8, uint16, uint32, uint64, float32, float64, string, or their slice,
// or a struct implementing the interface of Comparer.
func Compare(v1, v2 interface{}) int {
	if _v1, ok := v1.(Comparer); ok {
		return _v1.Compare(v2)
	}

	var first, second float64
	switch _v1 := v1.(type) {
	case int:
		if _v2, ok := v2.(int); !ok {
			panic("Type is not the same")
		} else {
			first, second = float64(_v1), float64(_v2)
		}
	case uint:
		if _v2, ok := v2.(uint); !ok {
			panic("Type is not the same")
		} else {
			first, second = float64(_v1), float64(_v2)
		}
	case int32:
		if _v2, ok := v2.(int32); !ok {
			panic("Type is not the same")
		} else {
			first, second = float64(_v1), float64(_v2)
		}
	case uint32:
		if _v2, ok := v2.(uint32); !ok {
			panic("Type is not the same")
		} else {
			first, second = float64(_v1), float64(_v2)
		}
	case int16:
		if _v2, ok := v2.(int16); !ok {
			panic("Type is not the same")
		} else {
			first, second = float64(_v1), float64(_v2)
		}
	case uint16:
		if _v2, ok := v2.(uint16); !ok {
			panic("Type is not the same")
		} else {
			first, second = float64(_v1), float64(_v2)
		}
	case int64:
		if _v2, ok := v2.(int64); !ok {
			panic("Type is not the same")
		} else {
			first, second = float64(_v1), float64(_v2)
		}
	case uint64:
		if _v2, ok := v2.(uint64); !ok {
			panic("Type is not the same")
		} else {
			first, second = float64(_v1), float64(_v2)
		}
	case int8:
		if _v2, ok := v2.(int8); !ok {
			panic("Type is not the same")
		} else {
			first, second = float64(_v1), float64(_v2)
		}
	case uint8:
		if _v2, ok := v2.(uint8); !ok {
			panic("Type is not the same")
		} else {
			first, second = float64(_v1), float64(_v2)
		}
	case float32:
		if _v2, ok := v2.(float32); !ok {
			panic("Type is not the same")
		} else {
			first, second = float64(_v1), float64(_v2)
		}
	case float64:
		if _v2, ok := v2.(float64); !ok {
			panic("Type is not the same")
		} else {
			first, second = _v1, _v2
		}
	case string:
		if _v2, ok := v2.(string); !ok {
			panic("Type is not the same")
		} else {
			return strings.Compare(_v1, _v2)
		}
	default:
		return CompareSlice(v1, v2)
	}

	if first > second {
		return 1
	} else if first < second {
		return -1
	} else {
		return 0
	}
}

// Same as Compare, but return true if v1 is less than v2, or return false.
func LT(v1, v2 interface{}) bool {
	return Compare(v1, v2) < 0
}

// Same as LT, but greater than.
func GT(v1, v2 interface{}) bool {
	return Compare(v1, v2) > 0
}

// Same as LT, but equal to.
func EQ(v1, v2 interface{}) bool {
	return Compare(v1, v2) == 0
}

// Same as LT, but not equal to.
func NE(v1, v2 interface{}) bool {
	return !EQ(v1, v2)
}

// Same as LT, but greater than or equal to.
func GE(v1, v2 interface{}) bool {
	return GT(v1, v2) || EQ(v1, v2)
}

// Same as LT, but less than or equal to.
func LE(v1, v2 interface{}) bool {
	return LT(v1, v2) || EQ(v1, v2)
}