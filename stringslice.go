package attic

// IsSameStringSliceNoOrder compare two string slice without order
func IsSameStringSliceNoOrder(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}

// DifferenceBetweenStringSlice returns the elements in `a` that aren't in `b`.
func DifferenceBetweenStringSlice(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// ShouldSameWithoutOrder custom assert for convey
// ref: https://github.com/smartystreets/goconvey/wiki/Custom-Assertions
// compare two string slice without order
func ShouldSameWithoutOrder(actual interface{}, expected ...interface{}) string {
	if IsSameStringSliceNoOrder(actual.([]string), expected[0].([]string)) {
		return ""
	}
	return "not same"
}