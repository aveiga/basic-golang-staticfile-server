# basic-golang-staticfile-server

## Topics covered

- ✅ REST
- [Messaging using AMQP]()
- [Input Validation]()
- [DB with PostgreSQL](https://bun.uptrace.dev)
- [DEV DB with SQLite](https://bun.uptrace.dev)
- [DB Versioning](https://bun.uptrace.dev/guide/migrations.html)
- Authentication and Authorization using OAuth v2
- [Service Discovery](https://github.com/ArthurHlt/go-eureka-client)
- [Rate Limiting](https://github.com/ulule/limiter)
- Logging
- Error Handling
- Testing
- [API Documentation](https://medium.com/@pedram.esmaeeli/generate-swagger-specification-from-go-source-code-648615f7b9d9)
- [Monitoring](https://prometheus.io/docs/guides/go-application/)
- [Websockets](https://github.com/gorilla/websocket)
- Developing and serving UI fragments
- Serving UI assets

## To cleanup

## FAQ

### How to get Keycloak to run on Docker Compose on M1 MacBooks 💻?

Quick answer:

- build the image locally (more info here: https://github.com/docker/for-mac/issues/5310)
- mount the pgdata volume to a directory below your home folder (or, preferably, in the repo folder)

### How to access the RabbitMQ Management UI?

- Go to http://localhost:15672/ using username and password: guest
