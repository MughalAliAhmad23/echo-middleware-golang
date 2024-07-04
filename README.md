# 1- You should install Go on your system. My version of Go is 1.22.4.

#2- Follow the instructions for using the Echo framework from the official site.

#3-Follow the instructions for JWT authentication through middleware from any Medium article.

#4- I use my PostgreSQL database containing tables named:
calculator with columns: id, no1, no2, operation, result
userrecord with columns: id, username, password, email
filestats with columns: id, totalline, totalwords, totalspaces, totalvowels, totalpunctuations, timestamp

#5- Initialize the framework and create each endpoint with the help of middleware to authorize and authenticate the generated token.

#6- First, users sign up with their credentials: username, password, and email.

#7- Log in with the username and password, and a token is generated.

#8-Authenticate and authorize the token. After that, operations will be performed.

#9- A part of the program also reads the file from the user and provides the result: total spaces, total vowels, total words, total sentences, and word frequency for each function. This is run in threads.

#10- Also, prepare the frontend with HTML and JavaScript to show the result of filestats on the browser at http://127.0.0.1:3000/frontend/index.html. 

#11- Users will receive a response of "File Uploaded Successfully" in JSON format, and then the stats of the uploaded file will be sent to the user through Gorilla WebSocket.

#12- SUMMARY: This program is about getting familiar with the Echo framework and the use of middleware, understanding CRUD operations, interacting with the database, creating tables and queries, and using concurrency in a specific part of the code to process a large file in minimum time.

#13- I also imported a collection of URLs. You can test them with Postman by providing the following parameters.

#14- To check the parameters online on the browser, there is also an implementation of Swagger that auto-generates the API documentation.

#15- To run the application as it runs on my local system, I also used Docker to Dockerize my application through the docker init command, which auto-generates the Docker files for the user.
