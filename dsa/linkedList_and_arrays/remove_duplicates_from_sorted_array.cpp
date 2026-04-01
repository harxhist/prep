// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int place = 0;
        int currIndex = 0;
        while(currIndex < nums.size()){
            if(currIndex == 0){
                place++;
                currIndex++;
                continue;
            }
            if(nums[currIndex] != nums[place - 1]){
                nums[place] = nums[currIndex];
                currIndex++;
                place++;
            }else{
                currIndex++;
            }
        }
        return place;
    }
};