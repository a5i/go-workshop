# Step 5. Add web UI

## download static UI files

You can find the source code here https://github.com/gothinkster/vue-realworld-example-app

And download prebuild here https://github.com/a5i/go-workshop/blob/master/static.zip

Extract files to **realworld/static**

```
static/
  img/
  js/
  index.html
  ...
```

**cmd/main.go**

```go
package main

import (
	"log"

	"example.com/realworld/httpservice"
	"github.com/labstack/echo"
)

func mainImpl() error {
	e := echo.New()
	s := httpservice.Service{}
	if err := s.SetupAPI(e); err != nil {
		return err
	}
	e.Static("/", "static")
	return e.Start(":3333")
}

func main() {
	if err := mainImpl(); err != nil {
		log.Fatal(err)
	}
}
```

## Run

```shell script
go run cmd/main.go
```

Open a browser http://127.0.0.1:3333