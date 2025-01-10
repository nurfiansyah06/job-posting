## Job Posting
This repository create APIs to manage job postings for companies.

### Features
- Save job details into the repository.
- Validate input data to ensure required fields are provided.
- Redis caching integration.

### Build with

- [Go](https://golang.org/dl/) (version 1.23.4)
- [Redis](https://redis.io/download)
- Dependencies managed using `go mod`.

### Instalation
1. Clone the repository
```
   git clone https://github.com/nurfiansyah06/job-posting.git
   cd job-posting
```
2. Install dependencies
```
go mod tidy
```
3. Run the project
```
go run main.go
```

### Documentation API
- [Postman](https://documenter.getpostman.com/view/11932880/2sAYQUquck)

### What you need to consider
- What if there are thousands of jobs in database?
To handle a large number of jobs in database, I would first implement using pagination to retreive jobs in smaller.
Next, I will optimize query sql with indexing. Indexes can speed up queries.
Additionally, I’d consider caching frequently accessed queries using a system like Redis to reduce database load.

- What if many users are accessing your API at same time?
To handle multiple users accesing an API, I would implement rate limiting in the system and if the system have high traffic
I would implement using load balancer, load balancer can help to distribute load across multiple server
