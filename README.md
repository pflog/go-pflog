# High Performance Logger For Human And Machine

Log formats are commonly designed for either humans or machines. Pflog
provides a log format (called *pformat*) that is easy to encode, easy
to read and easy to parse.

Pflog is designed follow the methodology of the [12-Factor App](https://12factor.net/),
where the app itself never concerns about the routing and storage of its
output stream.
*Pformat* is 100% structured and therefore able to get interpreted
and qualified by any other party. This allows you to build high
performance logging routers in your infrastructure.

Pflog provides a simple and intuitive API compatible with [google/glog](https://github.com/golang/glog).

## Repo layout

```
packages

pflog
- cmd                       // command starters
- log                       // log client
- containers                // predefined container types
- pkg/cli/...               // cli command code
- pkg/encoding/pformat      // format related stuff
- pkg/encoding/binary       // format related stuff
- pkg/encoding/json         // format related stuff
- pkg/logging/              // interfaces, Item/Entry, ... (or logger, or root?)
- pkg/logger/               // standard logger, noop logger
- pkg/registry

```

## Features

* 100% structured log format with typed contextual data
* language agnostig format
* optimized for reading (humans and standard tools like grep, awk, ...)
* optimized for parsing (capture qualified information of [12-Factor Apps](https://12factor.net/))
* lower storage size than JSON for long term archives
* able to get rid of unperformant format strings and blank `interface{}`'s


# Roadmap

- [ ] Pformat documentation
- [ ] Basic logger (client library) implementation
- [ ] Compare and optimize performance with [rs/logbench](https://github.com/rs/logbench)
- [ ] Further output formats
  - [ ] Binary (Compressed) for transportation and long term archives
  - [ ] JSON for analysis tools like Splunk or Elasticsearch
- [ ] CLI Tool for output format conversion
- [ ] Log router (basic implementation) to route logs to different destinations
- [ ] Log scraper, to read applications Stdout/Stderr and pass it to the router
