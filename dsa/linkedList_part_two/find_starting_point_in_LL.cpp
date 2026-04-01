// https://leetcode.com/problems/linked-list-cycle-ii/
//ChatGPT Diagram(mine was worse):
//             ┌───────────────┐
//             │               │
//             │       t       │
//             │               ↓
// Head ──→ A ──→ B ──→ C ──→ D ──→ E
//           ↑                   │
//           └────────── z ──────┘
//
// z = distance from Head to start of cycle (A)
// t = distance from cycle start to meeting point
// Cycle length = t + remaining part
//
// Slow & fast coincides initially after slow moves (z+t) distance & fast moves 2*(z+t) distance.
// Then we move slow to head.
// Then Mmve slow & fast at same pace of one. Both travel z distance & coincides again at starting point.
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        //Brute-force
        // map<ListNode*, int> mp;
        // ListNode* tmp = head;
        // while(tmp != NULL){
        //     if(mp.find(tmp) != mp.end()){
        //         return tmp;
        //     }
        //     mp[tmp]++;
        //     tmp = tmp -> next;
        // }
        // return NULL;
        // Optimal
        ListNode* slow = head;
        ListNode* fast = head;
        if(fast== NULL || fast -> next == NULL){
            return NULL;
        }
        do{
            slow = slow -> next;
            fast = fast -> next -> next;
        }while(slow!=fast && fast!= NULL && fast->next!=NULL);
        if(fast== NULL || fast -> next == NULL){
            return NULL;
        }
        slow = head;
        while(slow!=fast){
            slow = slow -> next;
            fast = fast -> next;
        }
        return slow;
    }

    //Official Solution:
        // ListNode* slow = head;
        // ListNode* fast = head;

        // // Step 1: Detect cycle
        // while (fast != NULL && fast->next != NULL) {
        //     slow = slow->next;
        //     fast = fast->next->next;

        //     if (slow == fast) {
        //         // Step 2: Find cycle start
        //         slow = head;
        //         while (slow != fast) {
        //             slow = slow->next;
        //             fast = fast->next;
        //         }
        //         return slow;
        //     }
        // }

        // return NULL;
};