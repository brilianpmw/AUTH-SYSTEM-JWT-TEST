# LoginApps
## Created for Test purposes
for production use need to store signature data to secret manager like GSM/anything



## Features

- Login User
- Access page that only user logged can access


## what can be improve : 
- create config based on secret manager
- completing the unit test 
- using mock uber for mocking https://pkg.go.dev/go.uber.org/mock/gomock
- implement gracefull shutdown 
- hash password and store in different place



## Development

```sh
make test
```

#### Building for source
```sh
make build
```
#### running the code
```sh
make run-http
```

## Docker
### Setup supporting service
*currently this service doesnt use any DB,Redis and anything but they are ready to use!
if you have add some DB/Cache/anything you might need to update docker-compose first!
Make sure you have docker service running.
On the root directory, run this command:
```
docker-compose up -d
```


## answer for several question
#### what is struct ?
A struct is a composite data type that groups together variables under a single name. Each variable in a struct is called a field.

you can find at presentation files

#### what is interface ?
An interface is a type that specifies a set of method signatures. A type satisfies an interface if it implements all the methods declared by the interface.

you can find at presentation files



#### golang package management
for using external package u just need to get the module to your go mod using go get module-name then using go mod vendor to store it in your local




