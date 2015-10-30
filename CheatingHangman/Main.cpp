#include "Game.h"


int main()
{   const int maxWrong = 8;
	//Game g(maxWrong,"dictionary.txt");
	bool playAgain = true;
	char response;
	while (playAgain)
	{ 
		Game g(maxWrong,"dictionary.txt");
		g.play();
	   cout << "Play Again?  Answer Y for yes\n";
	   cin >> response;
	   playAgain = toupper(response)=='Y';
	}

}
