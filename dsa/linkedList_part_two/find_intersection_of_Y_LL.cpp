//https://leetcode.com/problems/intersection-of-two-linked-lists/

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
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        //Bruteforce
        // map<ListNode*, int> mp;
        // ListNode* t1 = headA;
        // ListNode* t2 = headB;
        // while(t1 != NULL){
        //     mp[t1]++;
        //     t1 = t1 -> next;
        // }
        // while(t2 != NULL){
        //     if(mp.find(t2)!=mp.end()){
        //         return t2;
        //     }
        //     mp[t2]++;
        //     t2 = t2 -> next;
        // }
        // return NULL;

        //Better: Using length method

        // ListNode* t1 = headA;
        // ListNode* t2 = headB;
        // int l1 = 0;
        // int l2 = 0;
        // while(t1 != NULL){
        //     t1 = t1 -> next;
        //     l1++;
        // }
        // while(t2 != NULL){
        //     t2 = t2 -> next;
        //     l2++;
        // }
        // t1 = headA;
        // t2 = headB;
        // if(l1 > l2){
        //     int diff = l1 - l2;
        //     while(diff!=0){
        //         t1 = t1 -> next;
        //         diff--;
        //     }
        // }else if(l2 > l1){
        //     int diff = l2 - l1;
        //     while(diff!=0){
        //         t2 = t2 -> next;
        //         diff--;
        //     }
        // }
        // while(t1!=t2 && t1!=NULL && t2!=NULL){
        //     t1 = t1 -> next;
        //     t2 = t2 -> next;
        // }
        // if(t1 == t2 && t1!= NULL){
        //     return t1;
        // }
        // return NULL;

        //Optimal: Switch pointers
        ListNode * t1 = headA;
        ListNode * t2 = headB;
        while(t1!=t2){
            t1 = t1 -> next;
            t2 = t2 -> next;
            if(t1 == t2){
                return t1;
            }
            if(t1 == NULL){
                t1 = headB;
            }
            if(t2  == NULL){
                t2 = headA;
            }
        }
        return t1;
        //ChatGPT Solution:
        // ListNode *t1 = headA;
        // ListNode *t2 = headB;

        // while(t1 != t2){
        //     t1 = (t1 == NULL) ? headB : t1->next;
        //     t2 = (t2 == NULL) ? headA : t2->next;
        // }

        // return t1;
    }
};