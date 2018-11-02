// MIT License

// Copyright (c) 2018 Akhil Indurti

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package graph

type bellmanford struct {
    dist map[Vertex]int
    prev map[Vertex]Vertex
    neg bool
}

func (bf bellmanford) Distance(u, v Vertex) int { return bf.dist[v] }
func (bf bellmanford) Path(u, v Vertex) (path []Vertex) {
    if bf.dist[v] == Infinity {
        return
    }
    path = []Vertex{v}
    for bf.prev[v] >= 0 {
        v = bf.prev[v]
        path = append([]Vertex{v}, path...)
    }
    return path
}
func (bf bellmanford) NegativeCycle() bool { return bf.neg }

// SingleNegative finds the shortest path from source to
// every vertex in a graph that has negative edge weights.
func SingleNegative(g Interface, source Vertex) ShortestPath {
    bf := bellmanford{
        dist: make(map[Vertex]int),
        prev: make(map[Vertex]Vertex),
        neg: false,
    }
    vert := g.Vertices()
    for _, u := range vert {
        bf.dist[u] = Infinity
        bf.prev[u] = Uninitialized
    }
    bf.dist[source] = 0
    q := []Vertex{source}
    var u Vertex
    for len(q) != 0 {
        u, q = q[0], q[1:]
        L:
        for _, w := range g.Neighbors(u) {
            wt := g.Weight(u, w)
            if bf.dist[u] == Infinity || bf.dist[w] > bf.dist[u] + wt {
                bf.dist[w] = bf.dist[u] + wt
                bf.prev[w] = u
                for _, t := range q {
                    if t == w {
                        continue L
                    }
                }
                q = append(q, w)
            }
        }
    }
    for _, u := range vert {
        for _, v := range g.Neighbors(u) {
            wt := g.Weight(u, v)
            if bf.dist[u] < Infinity && bf.dist[u] + wt < bf.dist[v] {
                bf.neg = true
            }
        }
    }
    return bf
}
