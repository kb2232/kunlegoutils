package utils

func NxArrays[N ~[]float64 | ~[]int](intfloatArray *N, multiplier int) {
	switch arr := any(intfloatArray).(type) {
	case *[]int:
		for i, v := range *arr {
			(*arr)[i] = v * multiplier
		}
	case *[]float64:
		for i, v := range *arr {
			(*arr)[i] = v * float64(multiplier)
		}
	default:
		panic("unsupported type")
	}
}
