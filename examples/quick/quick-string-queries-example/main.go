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

	"github.com/williamsia/grammes/examples/exampleutil"
	"github.com/williamsia/grammes/quick"

	"go.uber.org/zap"
)

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

	// ------------------------------------- Executing Queries using QuickExecuteStringQuery
	logger.Info("Executing Basic String Queries...")

	{
		// Drop the vertices from the graph beforehand for no interference.
		quick.ExecuteStringQuery(addr, "g.V().drop()")

		// Adding a Vertex with traversal through QuickExecuteStringQuery
		quick.ExecuteStringQuery(addr, "g.addV('traversalVertex')")

		// Adding a Vertex with graph through QuickExecuteStringQuery
		quick.ExecuteStringQuery(addr, "graph.addVertex(T.label, 'graphingVertex')")

		// Storing a result byte array after executing a query.
		responses, err := quick.ExecuteStringQuery(addr, "g.V().label()")
		if err != nil {
			logger.Fatal("Error executing query", zap.Error(err))
		}

		// Log the resulting vertices.
		for _, res := range responses {
			logger.Info("executed string query", zap.ByteString("result", res))
		}
	}
}
