//https://www.naukri.com/code360/problems/aggressive-cows_1082559?leftPanelTabValue=PROBLEM
#include <algorithm>

bool solvable(int d, vector<int> &stalls, int k){
    //k cows, d distance in each cow & stalls = stalls
    int z = k - 1;
    int location = stalls[0];
    for(int i = 1; i < stalls.size(); i++){
        int loc = stalls[i];
        if(loc - location >=d){
            z--;
            location = loc;
        }else{
            continue;
        }
        if(z==0){
            return true;
        }
    }
    return false;
}

int aggressiveCows(vector<int> &stalls, int k)
{   
    sort(stalls.begin(), stalls.end());
    //d solvable problem
    int l = 0;
    int r = 1e9 + 1;
    while( r > l + 1){
        int mid = l + (r - l)/2;
        if(solvable(mid, stalls, k)){
            l = mid;
        }else{
            r = mid;
        }
    }
    return l;
}