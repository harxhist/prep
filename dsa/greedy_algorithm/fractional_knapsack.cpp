//https://www.geeksforgeeks.org/problems/fractional-knapsack-1587115620/1
//write comparator carefully
class Solution {
  public:
    
    static bool comp(tuple<int, int, double> v1, tuple<int, int, double> v2){
        return get<2>(v1) > get<2>(v2);
    }
        
    double fractionalKnapsack(vector<int>& val, vector<int>& wt, int capacity) {
        vector<tuple<int, int, double>> knap;
        int n = wt.size();
        for(int i = 0; i<n; i++){
            double vforw = double(val[i])/double(wt[i]);
            knap.push_back({val[i], wt[i], vforw});
        }
        sort(knap.begin(), knap.end(), comp);
        double maxvalue = 0.0;
        for(int i = 0; i<n;i++){
            tuple t = knap[i];
            int value = get<0>(t);
            int weight = get<1>(t);
            double unit = get<2>(t);
            if(weight > capacity){
                maxvalue += double(capacity) * unit;
                capacity = 0.0;
                break;
            }else{
                maxvalue += double(value);
                capacity -= weight;
            }
        }
        return maxvalue;
    }
};
