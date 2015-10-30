#include <string>
#include <iostream>
#include <fstream>
#include <assert.h>
using namespace std;

// Node in a doubly linked list of words
class Node
{
public:
	string word;  // data
	Node * next;  // next word
	Node * prev;  // previous word
	Node(string aword="",Node * anext=NULL);
};

