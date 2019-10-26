# Step 4. Add testing

**cmd/main_test.go**

```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"example.com/realworld/httpservice"
	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo"
)

var serverURL string

func TestMain(m *testing.M) {
	gofakeit.Seed(0)
	e := echo.New()
	s := httpservice.Service{}
	if err := s.SetupAPI(e); err != nil {
		log.Panic(err)
	}
	srv := httptest.NewServer(e)
	serverURL = srv.URL
	resultCode := m.Run()
	srv.Close()
	os.Exit(resultCode)
}

func TestHello(t *testing.T) {
	r, err := http.Get(serverURL)
	if err != nil {
		t.Errorf("get error %s", err)
		t.FailNow()
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("read body error %s", err)
		t.FailNow()
	}
	if string(body) != "RealWorld!" {
		t.Errorf("waits \"RealWorld!\" got %q", body)
		t.FailNow()
	}
}
```


```shell script
go test example.com/realworld/cmd
```