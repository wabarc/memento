# A Golang and Command-Line Interface to Time Travel Service

This package is a command-line tool named `memento` saving webpage to [Time Travel Service](http://timetravel.mementoweb.org/), it also supports imports as a Golang package for a programmatic. Please report all bugs and issues on [Github](https://github.com/wabarc/memento/issues).

## Installation

From source:

```sh
$ go get github.com/wabarc/memento
```

From [gobinaries.com](https://gobinaries.com):

```sh
$ curl -sf https://gobinaries.com/wabarc/memento | sh
```

From [releases](https://github.com/wabarc/memento/releases)

## Usage

#### Command-line

```sh
$ memento https://example.com https://example.org
https://example.com/ => https://arquivo.pt/wayback/20210416163608mp_/http://example.com/
https://example.org/ => http://wayback.archive-it.org/all/20210415130220/http://example.org/
```

#### Go package interfaces

```go
package main

package ia

import (
        "fmt"

        "github.com/wabarc/memento"
)

func main() {
        mem := &memento.Memento{}
        archives, _ := mem.Mementos(args)
        for orig, dest := range archives {
                fmt.Println(orig, "=>", dest)
        }
}

// Output:
https://example.com/ => https://arquivo.pt/wayback/20210416163608mp_/http://example.com/
https://example.org/ => http://wayback.archive-it.org/all/20210415130220/http://example.org/
```

## License

This software is released under the terms of the GNU General Public License v3.0. See the [LICENSE](https://github.com/wabarc/memento/blob/main/LICENSE) file for details.
