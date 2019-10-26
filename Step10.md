# Step 10. Update tests

Try to run

```shell script
go test example.com/realworld/cmd
```

Update **stor/stor.go**

```go
func (s *Storage) Clear() error {
	if err := s.db.Delete(&User{}).Error; err != nil {
		return err
	}
	if err := s.db.Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
```


Update **httpservice/arcticles.go**

```go
func (s *Service) ArticleList(c echo.Context) error {
	// ...
	resp.ArticlesCount = total
	resp.Articles = make([]article, 0, len(articles))
	// ..
	return c.JSON(http.StatusOK, resp)
}
```


Update **cmd/main_test.go**

```go
func TestMain(m *testing.M) {
	gofakeit.Seed(0)
	e := echo.New()
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

func TestArticles(t *testing.T) {
	e := httpexpect.New(t, serverURL)
	r := e.GET("/api/articles").Expect().JSON()
	r.Path("$.articles").Array().Length().Equal(0)
}

```