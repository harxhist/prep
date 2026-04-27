
//https://leetcode.com/problems/search-in-rotated-sorted-array/submissions/1989258414/
package main
func serch(nums []int, target int) int {
    l := 0
    r := len(nums) - 1
    for l<=r {
        mid := (l + r)/2
        if(nums[mid]==target){
            return mid;
        }
        if nums[mid] >= nums[l] {
            //left part is sorted
            if nums[mid] >= target && nums[l] <= target {
                r = mid - 1
            }else{
                l = mid + 1
            }
        }else{
            //right part is sorted
            if nums[mid] <= target && nums[r] >= target{
                l = mid + 1
            }else{
                r = mid - 1
            }
        }
    }
    return -1;
}