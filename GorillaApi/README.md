#### This is a simple Gorilla mux api I have developed with the help of MongoDB, a new database to me. The aim of this small project was to apply some concepts
#### of Apis and explore the world of non relational databases, as well as setting up something simple, yet useful for praticing my mock testing abilities.

##### Endpoints:

GET:  /api/languages
This method lists all of the compact information about the proggraming languages stored in the database.

Example output:
curl -X GET http://localhost:8000/api/languages

{"Go":"61c0d3959af98503e9fcecdb","Python":"61c0bd26822a5bb4270cac73"}


GET:  /api/languages/ID
This method displays the extended information about a programming language, given it's ID.

Example output:
curl -X GET http://localhost:8000/api/languages/61c0bd26822a5bb4270cac73

{"id":"61c0bd26822a5bb4270cac73","name":"Python","year":"1991","creators":["Guido van Rossum"],"frameworks":["Django","Flask"]}


POST:  /api/languages
This method will insert a new item to the collection, if there is no error of course.

Example output:
curl -X POST http://localhost:8000/api/languages -H 'content-type: application/json' -d '{"name":"Javascript", "year":"1995", "creators":["Brendan Eich"], "frameworks":["React", "Node"]}'

{"id":"61c0dba09af9852f90246f57","name":"Javascript","year":"1995","creators":["Brendan Eich"],"frameworks":["React","Node"]}
