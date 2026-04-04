//https://www.geeksforgeeks.org/problems/n-meetings-in-one-room-1587115620/1

class Solution {
  public:
    // Function to find the maximum number of meetings that can
    // be performed in a meeting room.
    struct meeting{
        int start;
        int end;
    };
    
    static bool comp(meeting m1, meeting m2){
        return m1.start < m2.start;
    }
    
    int maxMeetings(vector<int>& start, vector<int>& end) {
        int n = start.size();
        vector<meeting> meetings;
        for(int i = 0; i< n; i++){
            meeting m = {start[i], end[i]};
            meetings.push_back(m);
        }
        sort(meetings.begin(), meetings.end(), comp);
        int cnt = 1;
        int idx = 0;
        for(int i = 1; i < n; i++){
            meeting curr = meetings[i];
            meeting last = meetings[idx];
            if(curr.end < last.end){
                idx = i;
            }
            if(curr.start > last.end){
                cnt++;
                idx = i;
            }
            
        }
        return cnt;
    }
};


//Striver Solution:
// Function to get all meetings that can be scheduled
    vector<int> maxMeetings(vector<int>& start, vector<int>& end) {
        // Store meetings as (end_time, start_time, original_index)
        vector<tuple<int, int, int>> meetings;
        for (int i = 0; i < start.size(); i++) {
            // i+1 for 1-based indexing
            meetings.push_back({end[i], start[i], i + 1}); 
           
        }

        // Sort by end time
        sort(meetings.begin(), meetings.end());

        vector<int> result; // To store meeting indices
        int lastEnd = -1;

        // Traverse sorted meetings
        for (auto& m : meetings) {
            int e = get<0>(m);
            int s = get<1>(m);
            int idx = get<2>(m);

            // If meeting starts after last one ends
            if (s > lastEnd) {
                // Store index
                result.push_back(idx); 
                // Update last end time
                lastEnd = e; 
            }
        }
        return result;
    }
};
