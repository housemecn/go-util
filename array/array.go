package array

// InStringSlice 元素是否在一个string类型的slice里面
func InStringSlice(s []string, x string) (bool, int) {
    if !(s != nil && len(s) != 0) {
        return false, -1
    }
    for i, v := range s {
        if x == v {
            return true, i
        }
    }
    return false, -1
}

// ContainString determine whether the string element is in the array
func ContainString(array []string, item string) bool {
    if !(array != nil && len(array) != 0) {
        return false
    }
    for _, v := range array {
        if v == item {
            return true
        }
    }
    return false
}
