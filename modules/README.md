#Go modules

A module(dependency) has a name and usually follows `semver` of the format `vx.y.z` where

x - major (breaking)

y - feature (ideally not breaking)

z - patch

e.g, github.com/sirupsen/logrus:v1.18.0

##go.mod
Defines module name, go version, dependencies(also indirect) with version.
It also supports adding other metadata like replacing a dependency with a modified version.


```
module github.com/vishnukarthikl/gofun

go 1.16

require (
	k8s.io/klog v1.0.0
	rsc.io/quote v1.5.2
)

```

##go.sum
Lists all the dependencies with version and hashes. If the same version's code changes
in remote, the hash would change when `go mod tidy` is run.
```
github.com/go-logr/logr v0.1.0/go.mod h1:ixOQHD9gLJUVQQ2ZOR7zLEifBX6tGkNJF4QyIY7sIas=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
k8s.io/klog v1.0.0 h1:Pt+yjF5aB1xDSVbau4VsWe+dQNzA0qv1LlXdC2dF6Q8=
k8s.io/klog v1.0.0/go.mod h1:4Bi6QPql/J/LkTDqv7R/cd3hPo4k2DG6Ptcz060Ez5I=
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=

```

## Create a new module

`go mod init <module_name>`

e.g

`
go mod init github.com/vishnukarthikl/gofunc
`

## Import dependencies
If a package was imported(or removed) in the go code, run `go mod tidy` to update the `go.mod` file.
```go
package main

import (
	"k8s.io/klog"
)

func main() {
	klog.Info("Hi")
}

```
Note that this does not add/remove vendored code.

## Package/vendor dependencies
If the dependencies are to be packaged with the module, run `go mod vendor`. This will create
a directory called `vendor` and dumps all the required modules.
```
> tree vendor -L 1                                                                                                                                                                                             vishnukl@vishnukl-macbookpro

vendor
├── golang.org
├── k8s.io
├── modules.txt
└── rsc.io
```

note: `go mod tidy` does not add/remove code from vendor dir.  


## Package usage

Useful to know which module is using a specific dependency.
```
> go mod why -m k8s.io/klog                                                                                                                                                                                    vishnukl@vishnukl-macbookpro

# k8s.io/klog
github.com/vishnukarthikl/gofun
k8s.io/klog

```
