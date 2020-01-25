Go's HTTP server for serving static files brought to the CLI.

# Usage
```console
$ statico -help
Usage of statico:
  -port string
        Port to listen on. (default "8080")

$ statico
Serving files from directory /home/vaidik on port 8080
^C

$ cd movies && statico -port 4000
Serving files from directory /home/vaidik/movies on port 4000
```

# Installation
`go get github.com/codesoap/statico`

# License
`statico` is MIT licensed. See
[LICENSE.md](https://github.com/vaidik/statico/blob/master/LICENSE.md).
