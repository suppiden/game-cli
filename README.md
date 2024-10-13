🧠 Quiz Game
This is a simple CLI-based quiz game built with Go. It fetches random trivia questions from the Open Trivia API and lets you answer them in the command line interface. You can choose how many questions to answer, track your score, and compare it with other users who have played the quiz!

📋 Features
Fetches trivia questions from the Open Trivia Database
Dynamic number of questions (choose how many questions you want to answer)
Score tracking and comparison with other players
User-friendly CLI interface
Easy to replay as the same or different user
🚀 Getting Started
Prerequisites
To run this game, you'll need:

Go installed on your system (version 1.16+ recommended)
Installation
Clone the repository:

bash
Copiar código
git clone https://github.com/your-username/quiz-game.git
cd quiz-game
Install dependencies (if any):

bash
Copiar código
go mod tidy
Start the server:

bash
Copiar código
go run main.go
Run the quiz in a separate terminal:

bash
Copiar código
go run main.go start
🕹 How to Play
Choose how many questions you want to answer when prompted.
Enter your name.
Answer the questions by selecting the number corresponding to the answer you believe is correct.
Once you finish, the game will show your score and compare it to other players.
You can choose to play again as the same user (overwriting your score) or enter a new name for a different user.
To exit, simply type "no" when asked if you want to play again.
📚 Example Gameplay
yaml
Copiar código
How many questions do you want to answer? 3
Enter your name: Alice

Question 1: What is the capital of France?
1: Paris
2: Rome
3: Berlin
4: Madrid
Your answer: 1

Question 2: Who wrote 'The Catcher in the Rye'?
1: J.K. Rowling
2: Ernest Hemingway
3: J.D. Salinger
4: Mark Twain
Your answer: 3

Question 3: What is 2 + 2?
1: 3
2: 4
3: 5
4: 6
Your answer: 2

You scored 3 out of 3.
You were better than 100% of all quizzers.

Do you want to play again? (yes/no):
💻 Tech Stack
Go: The core language used to build both the server and the CLI.
Open Trivia API: Used to fetch random trivia questions.
🤝 Contributing
Feel free to submit issues or pull requests if you find a bug or want to improve the project!

📄 License
This project is licensed under the MIT License.
