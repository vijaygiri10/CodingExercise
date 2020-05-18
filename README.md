## Project Rest API Code Sample

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Docker,
Docker-Compose,
Lots of patience. ;)


### Installing


* [Docker installation on Mac](https://docs.docker.com/docker-for-mac/install/)
* [Docker installation on Linux(Debian based)](https://www.tutorialspoint.com/docker/installing_docker_on_linux.htm)


## Deployment
Go to project root directory and run following. "Base" Image is important to build first.

```
docker-compose build base
```
After successful build of Base Image run following:

```
docker-compose build http
```

Run all services altogether.

```
docker-compose up 
```

To run/build individual services use following:

```
docker-compose build/up <service name>

```

## Built With

* [Golang](https://golang.org/) - The Programmming language used.

* [Postgres](https://www.postgresql.org/) - PostgreSQL is a powerful, open source object-relational database system with over 30 years of active development that has earned it a strong reputation for reliability, feature robustness, and performance.

* [Docker](https://www.docker.com/)- Enterprise Application Container





## Authors
Vijay Kumar Giri


## License

@Open Source

## EndPoints If Running in Local Domain
[Create Student ENDPOINT]
* curl -X POST 'http://localhost:5050/create/student' \
--header 'Content-Type: application/json' \
--data-raw '{
	"name": "giri"
}'

[Create Assignment ENDPOINT]
* curl -X POST 'http://localhost:5050/create/student' \
--header 'Content-Type: application/json' \
--data-raw '{
	"student_id":1,
	"assignment_id":1,
	"score":30,
	"maximum_score":120
}'

[Update Student Assignment Score]
*  curl -X POST 'http://localhost:5050/update/student/score' \
--header 'Content-Type: application/json' \
--data-raw '{
	"name":"science_test",
	"maximum_score": 100
}'

[Delete Student Record]
* curl -X DELETE 'http://localhost:5050/delete/student/{student_ID}'

[Get All Students Record]
* curl -X GET 'http://localhost:5050/get/students'

[Assigin Student Score ]
* curl -X POST 'http://localhost:5050/create/student/score' \
--header 'Content-Type: application/json' \
--data-raw '{
	"student_id":3,
	"assignment_id":3,
	"score": 90
}'

[GET Specific Student Record]
* curl -X GET http://localhost:5050/get/students/1

[GET Assignment List]
* curl -X GET  http://localhost:5050/get/assignments