package test

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/suite"
)

func wrapRPCResult(res string) string {
	return `{"jsonrpc":"2.0","result":` + res + `,"id":0}`
}

func wrapRPCError(err string) string {
	return `{"jsonrpc":"2.0","error":` + err + `,"id":0}`
}

type MockJSONRPCServer struct {
	*httptest.Server
	body     []byte
	Counter  int
	Response interface{}
}

func MockJSONRPC(s *suite.Suite, response interface{}) (mock *MockJSONRPCServer, close func()) {
	require := s.Require()
	mock = &MockJSONRPCServer{
		Response: response,
		Server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			curResponse := mock.Response
			if a, ok := mock.Response.([]string); ok {
				curResponse = a[mock.Counter]
			}
			mock.Counter++

			// error => send RPC error
			if e, ok := curResponse.(error); ok {
				rw.WriteHeader(400)
				rw.Write([]byte(wrapRPCError(e.Error())))
				return
			}

			// string => convert to JSON
			if s, ok := curResponse.(string); ok {
				curResponse = json.RawMessage(wrapRPCResult(s))
			}

			var err error
			mock.body, err = io.ReadAll(req.Body)
			log.Println("rpc>>", string(mock.body))
			require.NoError(err)

			// JSON input, or serializable into JSON
			var responseBody []byte
			if v, ok := curResponse.(json.RawMessage); ok {
				responseBody = v
			} else {
				responseBody, err = json.Marshal(curResponse)
				require.NoError(err)
			}
			log.Println("<<rpc", string(responseBody))
			rw.Write(responseBody)
		})),
	}
	return mock, func() { mock.Close() }
}

type MockHTTPServer struct {
	*httptest.Server
	body     []byte
	Counter  int
	Response interface{}
}

func MockHTTP(s *suite.Suite, response interface{}) (mock *MockHTTPServer, close func()) {
	require := s.Require()
	mock = &MockHTTPServer{
		Response: response,
		Server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			curResponse := mock.Response
			if a, ok := mock.Response.([]string); ok {
				curResponse = a[mock.Counter]
			}
			mock.Counter++

			// error => send RPC error
			if e, ok := curResponse.(error); ok {
				rw.WriteHeader(400)
				rw.Write([]byte(wrapRPCError(e.Error())))
				return
			}

			// string => convert to JSON
			if s, ok := curResponse.(string); ok {
				curResponse = json.RawMessage(s)
			}

			var err error
			mock.body, err = io.ReadAll(req.Body)
			log.Println("http>>", req)
			require.NoError(err)

			// JSON input, or serializable into JSON
			var responseBody []byte
			if v, ok := curResponse.(json.RawMessage); ok {
				responseBody = v
			} else {
				responseBody, err = json.Marshal(curResponse)
				require.NoError(err)
			}
			log.Println("<<http", string(responseBody))
			rw.Write(responseBody)
		})),
	}
	return mock, func() { mock.Close() }
}
