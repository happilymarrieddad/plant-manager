package utils

func ExistsInt64(a int64, nums []int64) bool {
	for _, n := range nums {
		if n == a {
			return true
		}
	}
	return false
}

func Int64SliceToInterfaceSlice(nums []int64) (ret []interface{}) {
	for _, n := range nums {
		ret = append(ret, n)
	}
	return
}
