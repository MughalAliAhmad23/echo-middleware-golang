# 1- you should install Go on your system 
my version of Go is 1.22.4

#2- follow the instruction of using echo framework from offical site

#3- follow the instrution of jwt Authentication through middleware from any medium article 

#4- I use my postgreSql data base
containing tables named calculator with columns id, no1, no2, operation, result & userrecord with columns id, username, pasword, email & filestats with columns id, totalline, totalwords, totalspaces, totalvowels, totalpunctuations and timestamp

#5- initialzes the framework and make each end point with the help of middle ware to authorize and authenticate the generated token

#6- first user signup with their crediantials username, password and email 

#7- login with the username and password and a token is generated 

#8- Authenticate and Authorized the token after that operation will be perfomed 

#9- a part of a program also read the file from the user and gives the result total spaces, total vowels, total words, total sentences and word frequecy each 
function is run in threads 

#10- also ready the frontend with the help of HTML and JAVASCRIPT for show the result of filestats on browser http://127.0.0.1:3000/frontend/index.html 

#11- user can get a response of file Uploaded Successfully in json foam an then the stats of the uploaded file can be sent to user through Gorilla Web Socket

#12- SUMMARY:This program is all about getthing in touch with the concept of echo framework and use of middleware understanding of CURD operation getting in touch
with data base making their tables and their queries i also use concurrency in a specific part of a code to get the result of a large file in minimun time.

#13- i also importe the collection of urls you can test them with post man by giving the following parameters to them.

#14- to check the parameters online on browser there is also an implementation of Swagger that auto generates the api documentations

#15- to run that application as it is as it run on my system locally i also use Docker to Dockerize my application through Docker init command that auto generates the docker files for the user.
