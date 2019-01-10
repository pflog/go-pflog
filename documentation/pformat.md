# Pformat

Must contain:
- timestamp
- severity (simple <> complex?)

## Examples

```
I2006-01-02T15:04:05+07:00 Hello World!;

E2006-01-02T15:04:05+07:00 ctx{123-123-123-123} http{127.0.0.1 suzy "GET /foobar HTTP/1.0" 200 1234};
I2006-01-02T15:04:05+07:00 ctx{123-123-123-123} kubernetes{(hostname) (node) (labels)};
```


## Binary representation

```
Header:
  version byte
  headerLen uint32
  recordLen uint64
  severity byte
  timestamp int64

[severity byte][timestamp int64][5 container in bytes][container 1][container 2][container 3]

ContainerListHeader:
  len int32

ContainerHeader:
  offset int64
  type byte
  body len

[severity:rune][time:int64][[headerlen:int][type:rune]]

```

_.error{fooo}


### Container examples

```
http.req{UserAgent, Client IP, Authenticated User, Requested Ressource, Body length}
http.res{Body length, Response Code, Duration}
error{Something went wrong}
error.http.req{}

```

#### HTTP Audit Log

```
http{127.0.0.1 suzy "GET /foobar HTTP/1.0" 200 1234}

type HTTPRoundTrip struct {
    Method     string
	Client     string
	UserName   string
	Request    string
	StatusCode int
	BodyLength int64
}
```

#### Unstructured content

```
u{foo="bar", meh="xyz"}

type Map map[string]string
```

```
I2006-01-02T15:04:05Z07:00 ctx{123-123-123-123} grpc{127.0.0.1 - MyAPI.Myfunction *grpc code* 1234};

E2006-01-02T15:04:05Z07:00 ctx{123-123-123-123} Something horribly went wrong!; // message{Something horribly went wrong!}
```


### Parsing of unstructured content

```
Heap Dump... Aaaaah!!!!

Foo
Yeah...
```


### In code example

```
log.With(pflog.NewHTTPRequest(req), pflog.NewError(err), pflog.NewContext(ctx)).Info("Foobar")
```



## Command line tool

```
logfoo --context -f http.req:/^POST/ -f http.res.duration:>5000 -f error.http:* --show message
```
