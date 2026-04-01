// https://leetcode.com/problems/max-consecutive-ones/
class Solution {
public:
    
    //Complicated code
    // int findMaxConsecutiveOnes(vector<int>& nums) {
    //     int counter = nums[0] == 1 ? 1 : 0; 
    //     int maxi = counter;
    //     for(int i = 1;i<nums.size();i++){
    //         if(nums[i-1] == 1 && nums[i]==1){
    //             counter++;
    //         }else if(nums[i-1] == 0 && nums[i] ==1){
    //             counter = 1;
    //         }else if(nums[i] == 0){
    //             counter == 0;
    //         }
    //         maxi = max(counter, maxi);
    //     }
    //     return maxi;
    // }


    int findMaxConsecutiveOnes(vector<int>& nums) {
        int counter = 0;
        int maxi = 0;
        for(int i = 0;i<nums.size();i++){
            if(nums[i]==1){
                counter++;
            }else{
                counter = 0;
            }
            maxi = max(counter, maxi);
        }
        return maxi;
    }
};