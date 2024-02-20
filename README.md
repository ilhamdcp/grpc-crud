<h1>About</h1>
<p>This repo contains Go and Java code to simulate GRPC communication between service. The Java project will act as a GRPC service accepting request from GRPC client written in Go</p>

<h1>How to Run</h1>

<h2>Running Java Code</h2>
1. Run `./gradlew build` for Unix or `gradlew build` for Windows to fetch dependencies.<br>
2. Run `./gradlew run` for Unix or `gradlew build` for Windows to start the server.<br>
3. The server will run on localhost port 8090 by default.

<h2> Running Go Code</h2>
1. Run `go build` to get dependencies.<br>
2. Run `go run main.go` to run the client. The client will call `getItemById` and `getItemsByName` and then it will exit the process.<br>
3. The server will run on localhost port 8090 by default.<br>