# chatstodo-client

## 1. Install Go (skip if installed)

To install Go, head over to the [official Go website](https://go.dev/dl/) and follow the instructions to install it.

## 2. Create and edit .env file

Copy the file ".env.example" into a new filed named ".env". Open up the file using a text editor and add in the relevant information.

## 3. Install dependencies & Build

Run the following commands to install depdencies and build the server.

```
go install && go build
```

You should see a new executable file called "chatstodo-web" (or something similar, depending on OS).

## 4. Run the server

Run the executable and head over to the "LISTEN_ADDR" as specified in the ".env" file or in the terminal.