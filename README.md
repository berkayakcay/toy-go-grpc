# toy-go-grpc


[gRPC.io > GO > Quick start](https://grpc.io/docs/languages/go/quickstart/)
```sh
brew install protoc-gen-go
brew install protoc-gen-go-grpc
```


### gRPC vs REST

|**gRPC**               |**REST** 
|---	                |---
|Protocol Buffers   	|JSON
|HTTP 2   	            |HTTP 1
|Streaming   	        |Unary
|Bi directional   	    |Client->Server
|Free design   	        |GET/POST/UPDATE/DELETE/...



### Makefile (Windows)

Install Chocolatey (https://chocolatey.org/install)

```sh
choco install make
```

You should then be able to use make command in the directory that contains the Makefile.

This is optional, since you can still build the .proto files by hand by running the following commands:

```protoc -I${PROJECT}/proto --go_opt=module=${YOUR_MODULE} --go_out=. ${PROJECT}/proto/*.proto```
where ```${YOUR_MODULE}``` is the name of your go module (you can find that in your go.mod file) and ```${PROJECT}``` is one of the projects name (greet, calculator, blog).



### Notes

#### gRPC force to secure connection

https://grpc.io/docs/guides/auth/
```
Failed to connect: grpc: no transport security set (use grpc.WithTransportCredentials(insecure.NewCredentials()) explicitly or set credentials) 
```

#### vscode-proto3 extension configuration

Go into ``settings > Extensions > vscode-proto3 configuration`` and then click ``Edit in settings.json``. (you can just edit ``.vscode/settings.json`` too.)

After that, give ``--proto_path`` options like below codes.

```js
{
    "protoc": {
        "options": [
            "--proto_path=<path of your proto files>"
        ]
    }
}
```


# Resources
* https://husobee.github.io/golang/rest/grpc/2016/05/28/golang-rest-v-grpc.html
* https://www.slideshare.net/borisovalex/grpc-vs-rest-let-the-battle-begin-81800634
* https://learn.microsoft.com/en-us/aspnet/core/grpc/comparison?view=aspnetcore-3.0
