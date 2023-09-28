# Todo API Golang
I design with clean architecture with Golang by `Hexagonal Architecture`

## Structure 
- `db` like a layer of Repository, Database, External Service 
- `router` like a  layer of Controller, Framework, Delivery
- `todo` like a layer of Domain, Usecase, Business logic, Entity, Model, interface 


from above look like this one

`API <----> Domain <----> SPI`

in the Domain layer will have any interface you would like to another layer implement you can see from 

- `todo/repo/todo.repository.interface.go` 
 interface for Repository, Database or External Service to implement and

- `todo/todo.go` in this layer have a Business logic or any Usecase and interface for any Framework to implement


