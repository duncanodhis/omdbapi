# OMDBAPI
reads an api and filters data


How to run the program
1. You can go ahead and create the docker image using the docker file and then just run the docker image as shown
![Alt text](https://github.com/duncanodhis/omdbapi/blob/dce796922b793acab27040bdef5b5a060ae6a80c/Screenshot%20from%202022-11-06%2016-28-05.png "How to run it as a docker image")

2.Fill the flag options as shown in the diagram below:

![Alt text](Screenshot from 2022-11-06 16-03-46.png "Running using flags")


What the program does:

Task 5.1: Whe the  program exceeds the maxRunTime, it  gracefully exit. All resources  released and any running goroutines is  stopped.Output printed . 
Task 5.2: When your program receives the SIGTERM signal, it gracefully exit. All resources  released and any running goroutines is stopped. 
Task 5.3: The program is   able to handle any rate limiting from omdbapi. It does  not panic or error.
Task 6.1: The program implements the maxRequests flag and fast exits when the limit is reached.
Task 6.2: My unit test coverage  is  88.4%

