#include "Node.h"
class DoublyLinkedList
{	int size;  // Number of words in linked list
	Node * head;  // First node in list

	void insertBefore(string,Node *);
public:
	DoublyLinkedList(void) {head=NULL; size=0;}
	void listIt(int max=1000);
	void add(string);
	int getSize() {return size;}
	string get(int n);
	int biggestString();
	bool update(char c);
	bool update(int l);
	bool isinList(char c);
	bool inallWords(char c);
	//I used this function to help debug
	/*void printList()
	{
		Node * curr = head;
		while (curr != NULL)
		{  
			cout<<curr->word<<endl;
			curr = curr->next;
		}
	}*/
};

