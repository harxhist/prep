//https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

package main

func searchRange(nums []int, target int) []int {
    first := firstOccurence(nums, target)
    last := lastOccurence(nums, target)
    return []int{first, last}
}

func firstOccurence(nums []int, target int) int {
    l := -1 
    r := len(nums)
    for r > l + 1 {
        mid := l + (r - l )/2
        if nums[mid] < target {
            l = mid
        }else {
            r = mid
        }
    }
    if r < len(nums) && nums[r] == target{
        return r
    }
    return -1
}

func lastOccurence(nums []int, target int) int {
    l := -1 
    r := len(nums)
    for r > l + 1 {
        mid := l + (r - l )/2
        if nums[mid] <= target {
            l = mid
        }else {
            r = mid
        }
    }
    if l > -1 && nums[l] == target{
        return l
    }
    return -1
}