/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    // too many edge cases, dry run
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* fast = head;
        ListNode* slow = head;
        while(n!=-1 && fast!= NULL){
            fast = fast -> next;
            n--;
        }
        if(fast == NULL && n!=-1){
            return head -> next;
        }
        while(fast!=NULL){
            slow = slow -> next;
            fast = fast -> next;
        }
        slow -> next = slow -> next -> next;
        return head;
    }
};