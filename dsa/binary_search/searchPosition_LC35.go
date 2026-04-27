package main

//https://leetcode.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
    last := firsOccurence(nums, target)
    return last
}

func firsOccurence(nums []int, target int) int {
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
    return r
}