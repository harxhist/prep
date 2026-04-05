//https://leetcode.com/problems/sort-colors/
//k<=j
class Solution {
public:
    void swwap(vector<int>& nums, int i, int j){
        int n = nums[i];
        nums[i] = nums[j];
        nums[j] = n;
    }
    void sortColors(vector<int>& nums) {
        int i = 0;
        int j = nums.size() - 1;
        int k = 0;
        while(k <= j){
            if(nums[k]==0){
                swwap(nums, i, k);
                i++;
                k++;
            }else if(nums[k]==1){
                k++;
            }else{
                swwap(nums, j, k);
                j--;
            }
        }
    }
};
