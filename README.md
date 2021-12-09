#### After studying some of the concepts of the [second part](https://github.com/josnelihurt/golang-bootcamp/wiki/Part-2:-Language-basics#project-2-db-in-memory-cli-application) of the Golang Academy Bootcamp, this is a repository where I will be answering some question as well as making some exercises in order to test my newly acquired abilities in GO.

##### The following text is just a guide prepared by the bootcamp creator:

    How would you write a while statement in Go?
    What does the keyword defer do?
    Does Go support pointers? How do arguments get passed around(by value or by reference)?
    Are arrays in Go fixed length? How about slices?
    Say you have a map: map[string]int, how would do a lookup and check to see if the map holds the value of the key you were looking for?
    How does Go structure programs? What is the difference between a library and a program that executes?
    How do make a function or a type public? And how do you make it private?
    You are going to be building a simple calculator with 4 basic operations(add, subtract, multiply and divide). First build a library that provides those 4 methods. After that implement a program that reads from the command line the operation to be done and prints the result(by calling the library you implemented previously). The operation should be read however you'd like, but for simplicity sake limit yourself to 2 operands and 1 operation character. Something like ./program 1 + 2.
    Suppose you are building a web server that needs a DB that can do a set of simple operations. You know that the requirements of what DB to use will change. You also now that it will be easier for testing purposes to not have to setup something like MySQL. How would you solve this problem using the feature that Go provide?
    How would you build a simple function that can receive any type of argument and prints the if that argument is of a primitive type. Limit to just int, string, float and bool.
    How are errors defined in Go?
    Ok, you know how errors are defined in Go now. Time to build a simple errors package that allows you to build errors that specify what kind of error is it, limit yourself to 3 kinds: Internal, ThirdParty and Other. Then provide a function in that package for users to check if the error they have is of the kind they care about. NOTE: remember not to break with the way errors are defined in Go, take advantage of that.
    What do you use to make two functions concurrent?
    How would you synchronize two concurrent functions?
    Write a program with three functions. One will send stuff(whatever you'd like) over a channel every one second and one will receive it and print it. The third function will tell the other two functions to stop and return(it could be the main func) after 5 seconds. NOTE: the program can not end until the sender and receiver have returned.

## Made by Francisco Calixto
