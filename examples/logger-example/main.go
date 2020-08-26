// Copyright (c) 2018 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"flag"

	"go.uber.org/zap"

	"github.com/williamsia/grammes"
	"github.com/williamsia/grammes/examples/exampleutil"
)

// This example is to show how to
// switch out the default logger
// in the Grammes client with a custom
// one. In this instance we are going
// to use a zap logger.

// CustomLogger is our new logger
// to print using zap.
type CustomLogger struct {
	logger *zap.Logger
}

// PrintQuery will print the query out
// using the zap library rather than log.
func (c *CustomLogger) PrintQuery(q string) { c.logger.Info("QUERY", zap.String("cmd", q)) }

// Debug debugs the logs to stdout.
func (c *CustomLogger) Debug(msg string, fields map[string]interface{}) {
	arguments := []zap.Field{}
	for k, v := range fields {
		arguments = append(arguments, zap.Any(k, v))
	}
	c.logger.Debug(msg, arguments...)
}

// Error handles errors
func (*CustomLogger) Error(string, error) {}

// Fatal handles errors that are fatal
func (*CustomLogger) Fatal(string, error) {}

var (
	// addr is used for holding the connection IP address.
	// for example this could be, "ws://127.0.0.1:8182"
	addr string
)

func main() {
	flag.StringVar(&addr, "h", "", "Connection IP")
	flag.Parse()

	logger := exampleutil.SetupLogger()
	defer logger.Sync()

	if addr == "" {
		logger.Fatal("No host address provided. Please run: go run main.go -h <host address>")
		return
	}

	// Create a client with the custom logger.
	client, err := grammes.DialWithWebSocket(addr,
		grammes.WithLogger(&CustomLogger{logger: logger}),
	)
	if err != nil {
		logger.Fatal("Failed to create client", zap.Error(err))
	}

	// Try executing a query to see what the logger looks like.
	_, err = client.ExecuteStringQuery("g.addV('testvertex')")
	if err != nil {
		logger.Fatal("Error while adding vertex", zap.Error(err))
	}

	// Drop the testing vertex.
	err = client.DropAll()
	if err != nil {
		logger.Fatal("Error while dropping vertices", zap.Error(err))
	}
}
