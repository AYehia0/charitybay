# CharityBay
Application to gather charities by having a network of people contributing to a parant node, following the Onion (in a french accent) Architecture


## The Architecture
Well, huh, it's not that Onioun but it's a mix of other architectures with good practices: 
- Separation of concerns (aka abstracting away implementation details): well to keep related things within a boundry and as they always say if you want to change the response format you don't have to change the db schema!
- Using some known patterns like repository pattern
- Ports and Adapters (confusion with the Hexagonal Architecture!): A port is an input to your application, visa versa an Adapter is a way to talk to the external world (HTTP, gRPC or whatever)
- Dependancy Inversion: outer layers (implementation details) can refer to inner layers (abstractions), but not the other way around. (Adapters -> Ports -> Application -> Domain)
- 

Usually it's a good practice to hide the implementation so when you read the code (usecases aka application logic) you wouldn't know what db is used or what URL/External call is being called. and could be also extended by introducing a domain layer (from DDD) to hold the business logic.

Interfaces (in go) should be declared where they are used, that's why the repo are in the usecases (sorry .NET guys)
