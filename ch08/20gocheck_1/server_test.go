package main

import (
	"encoding/json"
	"gopkg.in/check.v1"
	_ "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PostTestSuite struct {}

func init() {
	check.Suite(&PostTestSuite{})
}

func Test(t *testing.T) {check.TestingT(t)}

func (s *PostTestSuite) TestHandleGet(c *check.C) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	var post Post
	mux.ServeHTTP(writer, request)
	c.Check(writer.Code, check.Equals, 200)

	json.Unmarshal(writer.Body.Bytes(), &post)
	c.Check(post.Id, check.Equals, 1)
}