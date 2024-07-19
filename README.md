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

### Users

* `GET /student`: retrieve a list of all student
  + **Postman Example**:
    - Method: GET
    - URL: `http://localhost:8080/student`
    - Headers: `Content-Type: application/json`
    - Response: `200 OK` with a list of students in JSON format
* `GET /student/:id`: retrieve a single student by ID
  + **Postman Example**:
    - Method: GET
    - URL: `http://localhost:8080/student/:id`
    - Headers: `Content-Type: application/json`
    - Response: `200 OK` with a list of student in JSON format
* `POST /student`: create a new student
  + **Postman Example**:
    - Method: POST
    - URL: `http://localhost:8080/student`
    - Headers: `Content-Type: application/json`
    - Response: `200 OK` with a list of student in JSON format
* `PUT /student/:id`: update a single student
  + **Postman Example**:
    - Method: PUT
    - URL: `http://localhost:8080/student/:id`
    - Headers: `Content-Type: application/json`
    - Response: `200 OK` with a list of student in JSON format
* `DELETE /student/:id`: delete a single student
  + **Postman Example**:
    - Method: DELETE
    - URL: `http://localhost:8080/student/:id`
    - Headers: `Content-Type: application/json`
    - Response: `200 OK` with a list of student in JSON format

### Subjects

* `GET /subject`: retrieve a list of all subjects
  + **Postman Example**:
    - Method: GET
    - URL: `http://localhost:8080/subject`
    - Headers: `Content-Type: application/json`
    - Response: `200 OK` with a list of subjects in JSON format
* `POST /subject`: create a new subject
  + **Postman Example**:
    - Method: POST
    - URL: `http://localhost:8080/subject`
    - Headers: `Content-Type: application/json`
    - Response: `200 OK` with a list of subject in JSON format
* `PUT /subject/:id`: update a single subject
  + **Postman Example**:
    - Method: PUT
    - URL: `http://localhost:8080/subject/:id`
    - Headers: `Content-Type: application/json`
    - Response: `200 OK` with a list of subject in JSON format
* `DELETE /subject/:id`: delete a single subject
  + **Postman Example**:
    - Method: DELETE
    - URL: `http://localhost:8080/subject/:id`
    - Headers: `Content-Type: application/json`
    - Response: `200 OK` with a list of subject in JSON format

Authors
-------

* [biyoba1] - [ladinnvalera@gmail.com]
