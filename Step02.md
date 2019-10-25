# Step 2. Add the server code

## Create the main.go

**cmd/main.go**

```go
package main

import (
	"net/http"
	
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "RealWorld!")
	})
	e.Logger.Fatal(e.Start(":3333"))
}
```

## Run

```shell script
go run cmd/main.go
```