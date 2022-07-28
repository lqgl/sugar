package tasty

import "golang.org/x/exp/constraints"

// SliceFiltrate 返回由 filtrate func 过滤后的 Slice
func SliceFiltrate[V any](collection []V, filtrate func(V, int) bool) []V {

	var result []V

	for i, v := range collection {
		if filtrate(v, i) {
			result = append(result, v)
		}
	}
	return result
}

// SliceUpdateElement 返回由 iteratee func 更新后的 Slice
func SliceUpdateElement[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))

	for i, t := range collection {
		result[i] = iteratee(t, i)
	}

	return result
}

// SliceUniq 返回由 iteratee func 过滤后的 Slice
// todo: fix empty return
func SliceUniq[T any, U comparable](collection []T, iteratee func(T) U) []T {
	result := make([]T, len(collection))

	seen := make(map[U]struct{}, len(collection))
	for _, item := range collection {
		key := iteratee(item)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
	}

	return result
}

// SliceGroupBy 返回由 iteratee func 处理后的 map
func SliceGroupBy[T any, U comparable](collection []T, iteratee func(T) U) map[U][]T {
	result := map[U][]T{}

	for _, item := range collection {
		key := iteratee(item)

		result[key] = append(result[key], item)
	}
	return result
}

// CheckInSlice  check value in slice
// T 可比较的类型 that supports the operators < <= >= >.
func CheckInSlice[T constraints.Ordered](a T, s []T) bool {
	for _, val := range s {
		if a == val {
			return true
		}
	}
	return false
}

// DelOneInSlice  delete one element of slice  left->right
// T 可比较的类型 that supports the operators < <= >= >.
func DelOneInSlice[T constraints.Ordered](a T, old []T) (new []T) {
	for i, val := range old {
		if a == val {
			new = append(old[:i], old[i+1:]...)
			return
		}
	}
	return old
}
