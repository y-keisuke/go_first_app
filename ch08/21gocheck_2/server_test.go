package main

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	"gopkg.in/check.v1"
	_ "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("Testing with Ginkgo", func() {
})

type PostTestSuite struct {
	mux    *http.ServeMux
	post   *FakePost
	writer *httptest.ResponseRecorder
}

func init() {
	check.Suite(&PostTestSuite{})
}

func Test(t GinkgoTInterface) { check.TestingT(t) }

func (s *PostTestSuite) SetUpTest(c *check.C) {
	s.post = &FakePost{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/post/", handleRequest(s.post))
	s.writer = httptest.NewRecorder()
}

func (s *PostTestSuite) TestGetPost(c *check.C) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, check.Equals, 200)
	var post Post
	json.Unmarshal(s.writer.Body.Bytes(), &post)
	c.Check(s.post.Id, check.Equals, 1)
}

func (s *PostTestSuite) TestPutPost(c *check.C) {
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, check.Equals, 200)
	c.Check(s.post.Id, check.Equals, 1)
	c.Check(s.post.Content, check.Equals, "Updated post")
}
