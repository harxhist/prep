https://www.geeksforgeeks.org/problems/minimum-platforms-1587115620/1

class Solution {
  public:
    static bool comp(pair<int, int> p1, pair<int, int> p2){
    if(p1.first == p2.first){
        return p1.second < p2.second;
    }
    return p1.first < p2.first;
}
  
    int minPlatform(vector<int>& arr, vector<int>& dep) {
        vector<pair<int, int>> times;
        int n = arr.size();
        for(int i =0; i< n; i++){
            times.push_back({arr[i], 0});
            times.push_back({dep[i], 1});
        }
        sort(times.begin(), times.end(), comp);
        int count = 0;
        int maxi = 0;
        for(int i = 0; i < 2*n; i++){
            if(times[i].second == 0){
                count++;
            }else{
                count--;
            }
            maxi = max(maxi, count);
        }
        return maxi;
    }
};
