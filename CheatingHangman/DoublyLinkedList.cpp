#include "DoublyLinkedList.h"


// insert word "w" into doubly linked list before the node pointed to by "next"
void DoublyLinkedList::insertBefore(string w, Node *next)
{  Node * n = new Node(w,next);
   size++;
   if (head==next)
   {head = n;
   }
}

//List the first "max" words in the linked list.
void DoublyLinkedList::listIt(int max)
{   Node * curr = head;
    int ct = 0;
	cout << "LIST IT " ;
    while (curr != NULL && ct < max)
	{  ct++;
	   cout << curr->word <<" ";
	   if (ct%10==0) cout << endl;
	   curr = curr->next;
	 }
	cout << "\n\n";
}	

//Add the word "w" to the first of the doubly linked list
void DoublyLinkedList::add(string w)
{   insertBefore(w,head);
}

// Return the "nth" word in the linked list
string DoublyLinkedList::get(int n)
{   int ct =1; 
    Node * curr = head;
	while (curr != NULL && ct < n)
	{   curr = curr->next;
	    ct++;
	}
	if (curr ==NULL) return "";
	return curr->word;
}

//Return the length of the longest word in the list
int DoublyLinkedList::biggestString()
{
	int length = 0;
	string temp;
	Node * curr = head;
	while (curr !=NULL)
	{
		temp=curr->word;
		if(temp.size()>length)
		{length = temp.size();}
		curr = curr->next;
	}
	return length;
}
//Updates the list and removes all the nodes with "char c" in it
bool DoublyLinkedList::update(char c)
{
	bool updated= false;
	Node * curr = head, *temp = NULL, *temp3 = NULL;
	string temp2;
	while(curr != NULL)
	{
		temp2=curr->word;
		if(temp2.find(c) != string::npos)// True if "c" is in "temp2"
		{
			if(curr==head)
			{head=head->next;
			if(head!=NULL)
			{head->prev=NULL;}
			delete curr;
			this->size--; //keeps track of the size of the linked list
			curr=head;}
			else
			{temp = curr->prev;
			temp3 = curr->next;
			temp->next = curr->next;
			if(temp3!=NULL)
			{temp3->prev = temp;}
			delete curr;				//delete the node with "char c" in it
			this->size--; //keeps track of the size of the linked list
			curr=temp;
			updated=true;		
			curr=curr->next;}
		}
		else
		{curr=curr->next;}
	}
	return updated;
}
//Updates the list and removes all the nodes that are not of length "l"
bool DoublyLinkedList::update(int l)
{
	bool updated= false;
	Node * curr = head, *temp = NULL, *temp3 = NULL;
	int Size;
	string temp2;
	while(curr != NULL)
	{
		temp2=curr->word;
		Size= temp2.size();
		if(Size!= l)// True if the size of "temp2" does not equal "l"
		{
			if(curr==head)
			{head=head->next;
			if(head!=NULL)
			{head->prev=NULL;}
			delete curr;
			this->size--; //keeps track of the size of the linked list
			curr=head;}
			else
			{temp = curr->prev;
			temp3 = curr->next;
			temp->next = curr->next;
			if(temp3!=NULL)
			{temp3->prev = temp;}
			delete curr;				//delete the node with "char c" in it
			this->size--; //keeps track of the size of the linked list
			curr=temp;
			updated=true;
			curr=curr->next;}
		}
		else
		{curr=curr->next;}
	}
	return updated;
}
//returns true if "c" is in any of the words in the list
bool DoublyLinkedList::isinList(char c)
{
	bool inlist=false;
	Node *curr=head;
	string temp;
	while(curr != NULL)
	{
		temp=curr->word;
		if(temp.find(c) != string::npos)// True if "c" is in "temp"
		{
			inlist=true;
		}
		curr=curr->next;
	}
	return inlist;
}
//returns true if char "c" is in every word in the list
bool DoublyLinkedList::inallWords(char c)
{
	bool inall=true;
	Node *curr=head;
	string temp;
	while(curr != NULL)
	{
		temp=curr->word;
		if(temp.find(c) == string::npos)// True if "c" is not in "temp"
		{
			inall=false;
		}
		curr=curr->next;
	}
	return inall;
}