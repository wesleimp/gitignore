# gitignore

![ci](https://github.com/wesleimp/gitignore/workflows/ci/badge.svg)

.gitignore generator for terminal

## Install

### Pre-compiled binary

**shell script**

```sh
curl -sf https://gobinaries.com/wesleimp/gitignore | sh
```

**manually**

Download the pre-compiled binaries from the [releases page](https://github.com/wesleimp/gitignore/releases) and copy to the desired location

### Docker

```sh
$ docker run -it --rm wesleimp/gitignore list
```

### Compiling from source

Clone the repository

```sh
$ git clone git@github.com:wesleimp/gitignore.git

$ cd gitignore
```

download dependencies

```sh
$ go mod download
```

build

```sh
$ go build -o gitignore main.go
```

verify it works

```sh
$ gitignore --help
```

## Usage

**list templates**

```sh
$ gitignore list
```

**generate**

```sh
$ gitignore generate --lang Go --lang yarn
```

**generate in another workspace**

```sh
$ gitignore generate --lang Go --lang yarn --path ./my/another/workspace
```

**append new language**

```sh
$ gitignore append --lang Go 
```

**append new custom text**

```sh
$ gitignore append --text "mybin.exe"
```

## License

[MIT](https://github.com/wesleimp/gitignore/blob/master/LICENSE)
