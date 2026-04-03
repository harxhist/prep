//https://leetcode.com/problems/palindrome-linked-list/
//Global left node & recursion

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

    ListNode* left;

    bool rec(ListNode* right){
        if(right == NULL){
            return true;
        }
        bool r = rec(right -> next);
        bool t = left -> val == right -> val;
        left = left -> next;
        return (r && t);
    }

    bool isPalindrome(ListNode* head) {
        left = head;
        return rec(head);
    }
};
// ChatGPT:
// ✅ Optimal Approach (O(n) time, O(1) space)
// 🔹 Steps:
// Find middle of linked list
// Use slow and fast pointers
// slow moves 1 step, fast moves 2 steps
// Reverse second half
// Reverse the list starting from slow
// Compare both halves
// Compare first half and reversed second half
// (Optional) Restore list (not required in interviews usually)