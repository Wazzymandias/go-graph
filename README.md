# go-graph
Graph implementation in Golang

go-graph provides a graph interface as well as implementation of
a Simple Graph on top of the provided interface.

Simple Graph implementation conceptually uses a set of vertices,
where each element in the set maps to its set of neighbors.

When an edge is associated between two vertices in a simple graph,
each vertex will have the other in its neighbor set.

A priority queue implementation is provided, which is used to implement
Dijkstra's algorithm.

The priority queue uses an indexed min heap, where each vertex
is inserted into the queue based on the vertex ID.

It is assumed that vertices will have distinct, unique IDs to avoid
collisions. Examples provided maintain that invariant.
Current implementation uses monotonically increasing integers to
identify vertices in the graph.

A Fibonacci Heap provides a more optimal solution, but takes longer to implement.

An indexed min heap requires auxiliary data structures to maintain
vertex index mappings, and performs additional operations to maintain
the heap and mappings. As a result, greater space is used, but operations
such as DecreaseWeight can access elements in constant time, improving runtime speed.

In terms of Big-O complexity, where V is the number of vertices, the space usage is roughly O(V) for mappings,
which is less than O(V^2) required to maintain the set of vertices and their neighbors.

However, in real world scenarios the additional space might be significant depending
on the number of vertices.

Overall time complexity of current implementation of Dijkstra's is `O(|E + V|log(V))`,
where E represents the total number of Edges in the graph, and V represents the total number of vertices in the graph.

In the worst case, all edges from vertices may perform the `DecreaseWeight` operation,
which is used to adjust priority in current implementation.

The decrease weight operation takes `log(V)` time, so for all edges that would be `E log V` time.

In the worst case, all vertices are processed and retrieved through the `ExtractMin` operation,
which requires `log(V)` time to maintain the heap invariants, yielding a runtime of `V log V`

### Requirements
  - [Go 1.10+](https://golang.org/dl/)

### Test
```
# ensure Go is installed

> go version

go version go1.10.3 darwin/amd64

# ensure GOPATH is set, or set with command below
> export GOPATH=$HOME/go

# clone repository to correct directory in GOPATH
> export DST_DIR=$GOPATH/src/github.com/Wazzymandias/go-graph
> mkdir -p $DST_DIR
> git clone https://github.com/Wazzymandias/go-graph $DST_DIR

# run all tests in repository
> cd $DST_DIR
> go test ./...
```

### Build
Build will compile the binary, which creates an example Simple Graph
and runs Dijkstra's on it. To build:
```
# ensure Go is installed

> go version

go version go1.10.3 darwin/amd64

# ensure GOPATH is set, or set with command below
> export GOPATH=$HOME/go

# clone repository to correct directory in GOPATH
> export DST_DIR=$GOPATH/src/github.com/Wazzymandias/go-graph
> mkdir -p $DST_DIR
> git clone https://github.com/Wazzymandias/go-graph $DST_DIR

# Build binary
> cd $DST_DIR
> go build -v

# Execute binary
> ./go-graph
```