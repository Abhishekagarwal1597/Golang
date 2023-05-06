#This code will find the first bad version of the given release(n)
#This uses the binary serach algorithm to find the mid value and searching on that.

func firstBadVersion(n int) int {
    left := 1
    right := n
    
    for left < right {
        mid := left + (right-left)/2

        if isBadVersion(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left
}
