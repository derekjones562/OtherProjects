#include "Game.h"
//Read the words from dictionary into repository
void	Game::readWords()
{
	string temp;
	while(ifs>>temp)
	{
		words.add(temp);
		cheatlist.add(temp);
	}
	cout <<  "WORD COUNT: " << words.getSize();
}

//Randomly select a word from the repository 
string Game::getWord()
{
	srand(time(0));
	int n = rand()% cheatlist.getSize()+1;
	//cout << "you asked for the " << n << "th word\n";
    return cheatlist.get(n);
}

//Ask the user for a guess, making sure he/she hasn't already guessed the letter
//"used" contains all the guessed letters
char Game::getGuess(string used)
{   char guess;
	cout << "\n\nEnter your guess: ";
	cin >> guess;
	guess = tolower(guess); //make lowercase since secret word in lowercase
	while (used.find(guess) != string::npos)  // indicates success
	{
		cout << "\nYou've already guessed " << guess << endl;
		cout << "Enter your guess: ";
		cin >> guess;
		guess = tolower(guess);
	}
	return guess;
}

/*Play a single game of hangman.  
1. The computer chooses a secret word, and then writes out a number of dashes equal to the word length.
2. The player begins guessing letters. 
   Whenever he guesses a letter contained in the hidden word, the computer reveals each instance of that 
   letter in the word. Otherwise, the guess is wrong.
3. The game ends either when all the letters in the word have been revealed or when the guesser has run out of guesses.
*/
/*cheating version 
1. The computer chooses a length and eliminates all other words in the list equal to that length
	, and then writes out a number of dashes equal to that length
2. The player begins guessing letters.
	the computer eliminates words from it list that contain each letter until 
	it can no longer delete words without deleting the list entirely.
	At this point forward, Whenever The player guesses a letter contained in the hidden word, the computer reveals each instance of that 
   letter in the word. Otherwise, the guess is wrong.
3. The game ends either when all the letters in the word have been revealed or when the guesser has run out of guesses.
*/

void Game::play()
{ 
	string theWord;// The word to guess
    int wrong = 0;// number of incorrect guesses
	int wordlength=getWordLength();
    string soFar(wordlength, '-'); // word guessed so far
	cheatlist.update(wordlength);
    string used = "";// letters already guessed
	bool wordchosen=false;// used to generatea word only once
    cout << "\t\t\tWelcome to Hangman. Good luck!\n\n\n" ;
	//cout  << "word is " << theWord << "\n";
    while ((wrong < maxWrong) && (soFar != theWord))
    {
        cout << "\n\nYou have " << (maxWrong - wrong) << " guesses left...\n";
        cout << "\nYou've used the following letters:\n" << used << endl;
        cout << "\nSo far, the word is:\n\n" << soFar << endl;
        char guess = getGuess(used);
        used += guess;
		if(wordchosen==false)//if the word has not been chosen yet
		{
			if(cheatlist.isinList(guess)&& !cheatlist.inallWords(guess))//check to see if the guess is in any of the words in the list and if all the words contain that letter
			{
				cheatlist.update(guess);	
			}
			else 
			{
				theWord= getWord();
				wordchosen=true;
			}
		}
			
        if (theWord.find(guess) != string::npos)  // True if "guess" is in "theWord"
        {
            cout << "That's right! " << guess << " is in the word.\n";
            // update soFar to include newly guessed letter
            for (int i = 0; i < theWord.length(); ++i)
               if (theWord[i] == guess)
                   soFar[i] = guess;
        }
        else
        {
            cout << "Sorry, " << guess << " isn't in the word.\n";
            ++wrong;
        }
    }
    // game over
    if (wrong == maxWrong)
        cout << "\nYou've been hanged!";
    else
       cout << "\nYou guessed it!";
	if(wordchosen=false)
		theWord=getWord();
    cout << "\nThe word was " << theWord << endl;
}

//randomly determine the length of the word 
int Game::getWordLength()
{
	srand(time(0));
	int size= rand()% words.biggestString()+1;
	if(size<3)
	{size= size+2;}//never have a word less than 3 char long
    return size;
}