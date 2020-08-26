# execute-bound-query

The basics of executing Gremlin queries that are bound using the `ExecuteBoundQuery` function in the Grammes client.

## Description

**execute-bound-query** demonstrates how to execute Gremlin queries that are bound. Specifically this example shows how to execute them using the `ExecuteBoundQuery` function in the Grammes client which takes in a `Query` object, bindings and re-bindings as `map[string]string`. Then returns a `[]byte` and `error`. The `[]byte` can be used to *unmarshal* into any kind of expected result such as a list of vertices.

## Prerequisites

- go 1.12
- Git
- Elastic Search
- Cassandra
  - Java 8

## Running

To run this test you will need a TinkerPop server running and a graph database to connect to locally. This example was tested while using JanusGraph which can be used by locating yourself to the root directory of the Grammes project.

``` sh
cd $GOPATH/src/github.com/williamsia/grammes
```

After locating yourself here then you may change directory to the `/scripts` folder.

``` sh
cd scripts
```

Finally you may run the `janusgraph.sh` script to begin a local instance of JanusGraph. This will begin the TinkerPop server for you as well.

``` sh
./janusgraph.sh
```

For further instructions please find yourself to the root [README.md](../../README.md) file.

## Steps

### General steps

- Create a [zap](https://github.com/uber-go/zap) logger to help explain what's going on in the test and display the results.
- Creates a Grammes client that connects to a locally hosted [TinkerPop](http://tinkerpop.apache.org/) server with a WebSocket.
  - For testing this was created using JanusGraph. This can be run in the `/scripts` directory.
- Drop all of the possible interfering vertices that were already on the graph.
- Defer a drop of all the testing vertices. This is done as clean up.

---

### Test specific steps

- Adds a vertex using a Query object and bindings to assign the label
