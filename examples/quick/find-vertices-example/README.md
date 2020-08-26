# find-vertices-example

The basics of querying for vertices using the `VertexIDs` and `VertexByID` functions from the `quick` package.

## Description

**find-vertices-example** demonstrates how to find vertices on the graph using the quick package. Specifically this example shows how use the `VertexIDs` and `VertexByID` functions which return vertex IDs based on provided labels and properties, and then return the entire vertex object based on the resulting ID, respectively.

---

## Prerequisites

- go 1.12
- Git
- Elastic Search
- Cassandra
  - Java 8

---

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

---

## Steps

### General steps

- Create a [zap](https://github.com/uber-go/zap) logger to help explain what's going on in the test and display the results.
- Adds two test vertices to the graph using `AddVertex` found in the `quick` package.
  - For testing this was created using JanusGraph. This can be run in the `/scripts` directory.
- `VertexIDs` returns the IDs for the vertices based on the label and properties provided.
- `VertexByID` returns the entire vertex object based on the IDs captured from `VertexIDs`.
- Drop all of the possible interfering vertices that were already on the graph.
- Defer a drop of all the testing vertices. This is done as clean up.

---

### Test specific steps

- Adds testing vertices to the graph
- Shows how to gather them from the graph using labels, properties, and IDs
- Logs the vertices
