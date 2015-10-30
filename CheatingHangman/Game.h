#include <string>
#include <ctime>
#include <iostream>
#include <fstream>
#include <assert.h>
#include "DoublyLinkedList.h"
using namespace std;

class Game
{
public:   
	int maxWrong;  // Maximun number of guesses allowed
	ifstream ifs;  // input stream containing dictionary
	DoublyLinkedList words; //Collection of words in dictionary
	DoublyLinkedList cheatlist; //Collection of words used inorder to cheat
 	void play();

	// Setup game and read dictionary from "filename"
	Game (int maxWrong = 8, char * filename ="dictionary.txt"){
		this->maxWrong = maxWrong;
		ifs.open(filename);
	    assert(ifs!= NULL);  // Abort if file is not found
		readWords();
	}
	void readWords();
	string getWord();
	char getGuess(string);
	int getWordLength();
};


