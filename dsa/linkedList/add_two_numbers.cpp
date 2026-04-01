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

 //Do it noob way
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int carry = 0;
        ListNode* ans = new ListNode(-1);
        ListNode* a = ans;
        while(l1 != NULL && l2 != NULL){
            int v1 = l1 -> val;
            int v2 = l2 -> val;
            int num = v1 + v2 + carry;
            int n = num%10;
            carry = num/10;
            ListNode* node = new ListNode(n);
            a -> next = node;
            a = a -> next;
            l1 = l1-> next;
            l2 = l2-> next;
        }
        while(l1 != NULL){
            int v1 = l1 -> val;
            int num = v1 + carry;
            int n = num%10;
            carry = num/10;
            ListNode* node = new ListNode(n);
            a -> next = node;
            a = a -> next;
            l1 = l1 -> next;
        }
        while(l2 != NULL){
            int v2 = l2 -> val;
            int num = v2 + carry;
            int n = num%10;
            carry = num/10;
            ListNode* node = new ListNode(n);
            a -> next = node;
            a = a -> next;
            l2 = l2 -> next;
        }
        if(carry!=0){
            ListNode* node = new ListNode(carry);
            a -> next = node;
            a = a -> next;
        }
        return ans -> next;
    }
};