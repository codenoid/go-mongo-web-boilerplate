# Golang Mongo Web Boilerplate 
[![Go Report Card](https://goreportcard.com/badge/github.com/codenoid/go-mongo-web-boilerplate)](https://goreportcard.com/report/github.com/codenoid/go-mongo-web-boilerplate)

WORK IN PROGRESS

## Installation

1. makesure your mongodb are running in 127.0.0.1:27017 without password, if you don't have mongo, you can install it [here](https://docs.mongodb.com/manual/administration/install-community/)
2. or you can set it up [here](https://github.com/codenoid/go-mongo-web-boilerplate/blob/644dc7a0b73e19ace25017495bc1293774155ef4/routes.go#L12)
3. run this command

```bash
git clone https://github.com/codenoid/go-mongo-web-boilerplate.git
cd go-mongo-web-boilerplate
go build
./gmwb
// open localhost:6969
```

## Usage

1. access [http://localhost:6969/setup](http://localhost:6969/setup) to save user data (for test purpose)
2. you will redirected to `/login` the username and password is `admin`
3. once success, you will redirected to `/`

## Features & TODO

- [x] Basic Project Structure
- [x] Shared Single Mongo Connection
- [x] Login/Logout page
- [ ] Redis use for session
- [ ] HTML Views
