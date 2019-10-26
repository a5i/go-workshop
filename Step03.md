__# Step 3. Organize the project

```shell script
mkdir httpservice
```

**httpservice/service.go**

```go
package httpservice

import (
	"net/http"

	"github.com/labstack/echo"
)

// Service represents the HTTP service
type Service struct {
}

// SetupAPI initializes the HTTP endpoints
func (s *Service) SetupAPI(e *echo.Echo) error {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "RealWorld!")
	})
	return nil
}
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

	return e.Start(":3333")
}

func main() {
	if err := mainImpl(); err != nil {
		log.Fatal(err)
	}
}
```
