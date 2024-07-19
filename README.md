Electronic Diary Server
=====================

A RESTful API for managing electronic diaries.

### <font color="#008000">Getting Started</font>
---------------

### Prerequisites

* Go 1.17 or later
* PostgreSQL 13 or later
* Docker (optional)

### Building and Running

1. Clone the repository: `git clone https://github.com/biyoba1/electronic-diary-server-1.git`
2. Navigate to the project directory: `cd electronic-diary-server-1/electronic-diary`
3. Build the application: `go build -o main main.go`
4. Run the application: `./main`

Alternatively, you can use Docker to build and run the application:

1. Build the Docker image: `docker build -t electronic-diary-server .`
2. Run the Docker container: `docker run -p 8080:8080 electronic-diary-server`

API Endpoints
-------------

### <font color="#ff9900">Users</font>

* `GET /student`: retrieve a list of all student
* `GET /student/:id`: retrieve a single student by ID
* `POST /student`: create a new student
* `PUT /student/:id`: update a single student
* `DELETE /student/:id`: delete a single student

### Subjects

* `GET /subject`: retrieve a list of all subjects
* `POST /subject`: create a new subject
* `PUT /subject/:id`: update a single subject
* `DELETE /subject/:id`: delete a single subject

Authors
-------

* [biyoba1] - [ladinnvalera@gmail.com]
