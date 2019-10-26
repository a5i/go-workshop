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
	"example.com/realworld/stor"
	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo"
	"gopkg.in/gavv/httpexpect.v2"
)

var serverURL string

func TestMain(m *testing.M) {
	gofakeit.Seed(0)
	e := echo.New()
	s := httpservice.Service{}
	st, err := stor.New()
	if err != nil {
		log.Panic(err)
	}
	if err := st.Migrate(); err != nil {
		log.Panic(err)
	}
	if err := st.Clear(); err != nil {
		log.Panic(err)
	}
	s := httpservice.Service{Stor: st}
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
func TestArticles(t *testing.T) {
	e := httpexpect.New(t, serverURL)
	r := e.GET("/api/articles").Expect().JSON()
	r.Path("$.articles").Array().Length().Equal(0)
}
```


```shell script
go test example.com/realworld/cmd
```