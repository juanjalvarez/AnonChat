# AnonChat

An authentication-less messaging webapp.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them


1. [node](https://nodejs.org/en/)
2. [golang](https://golang.org/)
3. [golang/dep](https://github.com/golang/dep)


### Installing

A step by step series of examples that tell you have to get a development env running

Clone the repository to your $GOPATH/src directory.

```
git clone http://github.com/juanjalvarez/anonchat $GOPATH/src/anonchat && cd $GOPATH/src/anonchat
```

Install the frontend dependencies.

```
cd frontend && npm install && cd ..
```

Install the backend dependencies.

```
cd backend && dep ensure && cd ..
```

## Running the project

Since AnonChat requires two separate web servers, you will need to follow the two next instructions in separate terminal windows/tabs

1. Frontend
```
cd frontend && npm start
```

2. Backend
```
cd backend && ./run.sh
```

## Contributing

Make sure to write code according to the following guidelines.

1. [Frontend](https://standardjs.com/rules.html)
2. [Backend](https://github.com/golang/go/wiki/CodeReviewComments)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
