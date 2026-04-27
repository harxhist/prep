package main
func search(nums []int, target int) int {
    l := -1 
    r := len(nums)
    for r > l + 1 {
        mid := l + (r-l)/2
        if nums[mid] < target {
            l = mid;
        }else{
            r = mid;
        }
    }
    if r < len(nums) && nums[r] == target {
        return r;
    }
    return -1;
}