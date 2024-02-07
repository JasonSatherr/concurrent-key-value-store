# Concurrent key-value datastore

---
## This is a lo-fi implementation of a key-value datastore that is able to concurrently handle requests to access the database over the network 📖

---
## Why Concurrency? 🐝🐝🐝

The pattern of concurrency is a useful paradigm for designing systems because it breaks down application logic into independently executing sections.  As [Rob Pike explains in his 2012 presentation](http://www.youtube.com/watch?v=f6kdp27TYZs), this enables easy implementation of parallelism within the application to make it run faster.  The tldr behind the talk is that because you have already broken your code into pieces that compute independently, if you find that a section of your application is creating a bottle neck, you can just add another concurrent process to do that task, theoretically doubling your throughput.

--- 
## How does this datastore work?

In short, the code first recieves a request to access the datastore in some way over the network.  It will then wait for an appropriate time to service the request to either read or write to the datastore.  Finally, a response will be sent to the client.

In this project, we use independently executing go routines and channels in order to break out the above logic into a concurrent design.
We have go routine(s) that take exclusive responsibility over the following:
1. Recieving the incoming network requests
2. Writing to the database
3. Reading from the database
4. Determining if it is safe to read from the database (are any writes currently happening?  If so, postpone reading)

In a frequently read from database, we would be able to increase the number of go routines or workers that work on task 3 which would increase the performance of the application.

---

## Let's try it! 🧪
fork the repo and run with ``` go run ./src ```

then use this curl command for updating the datastore
```
curl --location --request PUT 'http://localhost:8080/person/update' \
--header 'Content-Type: application/json' \
--data '{
    "id": 100,
    "name": "mob"
}'
```

and this one to get the data
```
curl --location --request GET 'http://localhost:8080/person/get' \
--header 'Content-Type: application/json' \
--data '{
    "id": 100
}'
```
