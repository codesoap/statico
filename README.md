Minimalistic static file server using Go's
[http.FileServer()](https://golang.org/pkg/net/http/#FileServer).

# Usage
```console
$ statico -help
Usage of statico:
  -p int
        Port to listen on. (default 8080)

$ statico
Serving files from directory /home/vaidik on port 8080
^C

$ cd movies && statico -p 4000
Serving files from directory /home/vaidik/movies on port 4000
```

# Installation
```shell
git clone 'https://github.com/codesoap/statico.git'
cd statico
go install
# The statico binary is now installed at $HOME/go/bin/ .
```

If you don't want to install from source, you can download binaries from
the [releases page](https://github.com/codesoap/statico/releases).

# License
`statico` is MIT licensed. See
[LICENSE.md](https://github.com/vaidik/statico/blob/master/LICENSE.md).
