struct ListNode
{
	int val;
	ListNode *next;
	ListNode() : val(0), next(nullptr) {}
	ListNode(int x) : val(x), next(nullptr) {}
	ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution
{
public:
	// Merge sort use bottom-up policy,
	// so Space Complexity is O(1), Time Complexity is O(NlgN)
	ListNode *sortList(ListNode *head)
	{
		if (!head || !(head->next))
			return head;

		ListNode *cur = head;
		int len = 0;
		while (cur)
		{
			len++;
			cur = cur->next;
		}

		ListNode dummy(0);
		dummy.next = head;
		ListNode *left, *right, *tail;
		for (int step = 1; step < len; step <<= 1)
		{
			cur = dummy.next;
			tail = &dummy;
			while (cur)
			{
				left = cur;
				right = split(left, step);
				cur = split(right, step);
				tail = merge(left, right, tail);
			}
		}
		return dummy.next;
	}

private:
	// Divide the linked list into two lists, while the first list contains first n ndoes, return the second list's head.
	ListNode *split(ListNode *head, int n)
	{
		for (int i = 1; head && i < n; i++)
			head = head->next;

		if (!head)
			return nullptr;

		ListNode *second = head->next;
		head->next = nullptr;

		return second;
	}

	//  merge the two sorted linked list l1 and l2, then append the merged sorted linked list to the node head
	// return the tail of the merged sorted linked list
	ListNode *merge(ListNode *l1, ListNode *l2, ListNode *head)
	{
		ListNode *cur = head;
		while (l1 && l2)
			if (l1->val > l2->val)
			{
				cur->next = l2;
				cur = l2;
				l2 = l2->next;
			}
			else
			{
				cur->next = l1;
				cur = l1;
				l1 = l1->next;
			}

		cur->next = (l1 ? l1 : l2);

		while (cur->next)
			cur = cur->next;

		return cur;
	}
};