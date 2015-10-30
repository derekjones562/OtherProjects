#include "Node.h"


// Create a new node before "anext" in the linked list.
Node::Node(string aword,Node * anext)
{ word=aword;
  next =anext;
  if (next==NULL){
	  prev = NULL;
  }
  else
  {
	  prev = next->prev;
	  next->prev = this;
	  if (prev != NULL) prev->next = this;
  }
  //cout  << "NODE " << word  << " "<< next << " " << prev << "END\n";
}

