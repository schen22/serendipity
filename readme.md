dart lang
flutter framework

server:
go lang through gRPC send and server-side streaming
protobuf to define API for client-server communication

following: https://github.com/amsokol/flutter-grpc-tutorial

## Setup

`brew install go`
`brew tap dart-lang/dart`
`brew install dart`

[flutter](https://flutter.io/docs/get-started/install)

- Extract file in desired location and unzip
- add flutter to path (usually within .bashrc):

install latest [protobuf version](https://github.com/protocolbuffers/protobuf/releases/tag/v3.6.1)

Install Go code generator plugin for Proto compiler

`go get -u github.com/golang/protobuf/protoc-gen-go`

Install protoc plugin for Dart
`pub global activate protoc_plugin`

Example modified ~/.bashrc:

```
export PATH=$PATH:$(go env GOPATH)/bin
export PATH="$HOME/protoc-3.6.1-osx-x86_64/bin/:$PATH"
export PATH="$HOME/flutter/bin/:$PATH"
export PATH="$PATH":"$HOME/.pub-cache/bin"
```

compiling [proto code](https://grpc.io/docs/tutorials/basic/go.html#generating-client-and-server-code):
`protoc -I routeguide/ routeguide/route_guide.proto --go_out=plugins=grpc:routeguide`

