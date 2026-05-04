package main
func findPeakElement(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    if nums[0] > nums[1] {
        return 0
    }
    if nums[len(nums)-1] > nums[len(nums)-2] {
        return len(nums)-1
    }
    l := 1
    r := len(nums)-2
    for l<=r {
        mid := (l+r)/2
        if nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1]{
            return mid
        } else if nums[mid] > nums[mid - 1] {
            l = mid + 1
        }else {
            r = mid - 1
        }
    }
    return -1
}