# Load test
I used Domain Driven Design (DDD) for this project. it's a fantastic architecture that focuses on business logic.

The business domain doesn't care what framework you use for controllers or the databases to which the repository is hooked.


The separation of concern is quite resourceful because it's made possible by the use of interfaces, which means you can switch frameworks or databases very easily.

## how to run the app

the easiest way to check the functionality of the load test visit 

https://github.com/myrachanto/loader/actions - github actions

and check the test 

## or

depending on how you have set up your machine 

you can use either to pull the code to your local machine

@ ssh git clone git@github.com:myrachanto/loader.git

or

https git clone https://github.com/myrachanto/loader.git

then you can navigate to the project folder run "go run main.go"

"I am assuming you have golang already set in you machine!"

## apis
I have build two apis fo this load test

http://localhost:3500/api/load?url="https://jsonplaceholder.typicode.com/todos/1"

which can also work as

http://localhost:3500/api/load -- makes the query

this works because if the url passed as query parameter is empty then the url will be assinged to https://jsonplaceholder.typicode.com/todos/1 in the code

the other is 

http://localhost:3500/api/loads --gets the results 

this api get the store data in array

## Testing Thread safe
so I did write a test to test how thread safe the apis are or better yet how thread safe the logic is

and as such I

span 3 goroutines each perfomming 10 calls https://jsonplaceholder.typicode.com/todos/1 and for better results
i used the modulas of 4 of each goroutines calls to make an api call with an invalid url so as to obtain 404 for better array results

again those results are displayed at 

https://github.com/myrachanto/loader/actions on the test part of the github actions
