# Step 4. Add testing

**cmd/main_test.go**

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

	return e.Start(":3333")
}

func main() {
	if err := mainImpl(); err != nil {
		log.Fatal(err)
	}
}
```


```shell script
go test example.com/realworld/cmd
```