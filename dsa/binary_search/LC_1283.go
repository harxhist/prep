package main

func monotone(x int, nums []int, threshold int) bool {
    t := 0
    for _, v := range nums {
        if v % x == 0 {
            t += v/x
        }else{
            t += 1 + v/x
        }
    }
    return t <= threshold
}

func smallestDivisor(nums []int, threshold int) int {
    l := 0
    r := 1000001
    for r > l + 1 {
        mid := l + (r - l)/2;
        if !monotone(mid, nums, threshold){
            l = mid
        }else{
            r = mid
        }
    }
    return r
}