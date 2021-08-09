# rest-server

A REST test client to test X-Road services

## Getting started
You can download compiled binaries for your architecture:

* [Linux x64](https://github.com/digitaliceland/xrd-rest-client/raw/master/release/linux/rest-client)
* [MacOS x64](https://github.com/digitaliceland/xrd-rest-client/raw/master/release/macos/rest-client)
* [MacOS M1](https://github.com/digitaliceland/xrd-rest-client/raw/master/release/macosm1/rest-client)
* [Windows x64](https://github.com/digitaliceland/xrd-rest-client/raw/master/release/windows/rest-client.exe)

## Parameter Options

```
  -client string
    	Your X Road Client ID (default "CS/ORG/1111/TestClient")
  -cmd string
    	method to call on your REST server
  -loop
    	repeatedly call the function every second
  -service string
    	Your X Road Service ID (default "CS/ORG/1111/TestService/TEST123")
  -ss string
    	Your X Road Security Server URL without trailing / (default "http://localhost:80")
```

## Usage examples
Run the script and view help information:

```console
$ ./rest-client 
```

Full example calling the time service in a loop

```console
$ ./rest-client -client CS/ORG/1111/TestClient -cmd time -loop -service CS/ORG/1111/TestService/TEST123 -ss http://localhost:80
```


## Debugging

You can run code directly with go like this:
```console
$ go run main.go
```

## Building

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`

Then to compile and run it just:

```console
$ go build main.go
$ ./main
```

To build for multiple architectures we also provide a build script, build.sh
