About the project
The details of the project can be found at `resources/uber.pdf`.  

How to read the project:
1. The entrypoint of the project is `cmd/main.go`.
2. The functions case-sensitive `init` are special to golang and are executed automatically when its parent package is 
imported. This is executed once only even if the package is imported multiple times in various different files. This is 
ideally used to achieve the singleton pattern. The init `config/config.go` is implicitly called from `cmd/main.go` by 
doing a blank import. The `config/config.go` initializes the viper library which is a famous config management library 
for golang.
3. The `internal/` contains all the business logic and the internal keyword is special to golang as the packages inside 
it cannot be exposed to the outer world (other projects). Inside the internal folder, the packages are created based on 
their separation of concerns.
4. The `pkg/` folder contains the library equivalent codebase which can also be exposed to the outer world. 
5. The `resources/` folder contains the problem statement pdf, and the API collection in the json format.
6. The `internal/api/api_registrar.go` contains the segregated apis for user, vehicle, driver_task and customer_task. 
7. The directories,`internal/user`, `internal/vehicle`, `internal/driverTask` and  `internal/customerTask` follow the 
similar structure of handler->controller->dao MVC pattern. The `factory.go` file initializes the handler and controller 
objects.
8. The `api_registrar` calls `handler.go` which in turn call `ctrl.go` which in turn call `dao.go` in the respective 
modules.

What's not so good about the project:
1. Test coverage is not great ~ 55-60%. Since the project was large.
2. Some validations are missing.
3. TODOs and comments need addressing.

What's good about the project:
1. Automated testing framework model.
2. All the tests have been written in test suite fashion.
3. The service(s) are dockerized to make the project platform independent.
4. The whole project is written in a very organized MVC based pattern incorporating the handler-controller-dao layers.

How to run the project:
1. Change directory to the root of the project, i.e, `~/.../uber`.
2. Run: `make docker-build`.  This command starts the services in docker-compose. 
3. Run `make migrate_up`. This command applies the migrations.
4. Use the API collection present in `resources/uber.postman_collection.json` to verify the APIs. This json file can 
be directly imported in the postman app.

How to run tests:
1. Run `make docker-test` to run the tests.
2. Run: `make docker-test-coverage` to check code coverage.

