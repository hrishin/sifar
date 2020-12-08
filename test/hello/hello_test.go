package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/hrishin/sifar/pkg/hello"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHelloWorldAPI(t *testing.T) {

	Convey("Invoking /hello endpoint", t, func() {

		response := invoke("/hello")

		Convey("It should return 200 OK HTTP status", func() {
			So(response.Result().StatusCode, ShouldEqual, http.StatusOK)
		})

		Convey("It should return 'hello world' in response", func() {
			So(response.Body.String(), ShouldEqual, "hello world")
		})
	})
}

func invoke(uri string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	routes := mux.NewRouter()
	hello.MountRoutes(routes)
	rr := httptest.NewRecorder()
	routes.ServeHTTP(rr, req)

	return rr
}
