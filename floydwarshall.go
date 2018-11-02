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

type floydwarshall struct {
    dist map[Vertex]map[Vertex]int
    next map[Vertex]map[Vertex]*Vertex
    neg bool
}

func (fw floydwarshall) Distance(u, v Vertex) int { return fw.dist[u][v] }
func (fw floydwarshall) Path(u, v Vertex) (path []Vertex) {
    if u == v {
        return []Vertex{u}
    }
    if fw.next[u][v] == nil {
        return
    }
    path = []Vertex{u}
    for u != v {
        u = *fw.next[u][v]
        path = append(path, u)
    }
    return path
}
func (fw floydwarshall) NegativeCycle() bool { return fw.neg }

// AllPairs finds the shortest paths between all vertex pairs
// in a graph that has negative edge weights.
func AllPairs(g Graph) ShortestPath {
    vert := g.Vertices()
    fw := floydwarshall{
        dist: make(map[Vertex]map[Vertex]int),
        next: make(map[Vertex]map[Vertex]*Vertex),
        neg: false,
    }
    for _, u := range vert {
        fw.dist[u] = make(map[Vertex]int)
        fw.next[u] = make(map[Vertex]*Vertex)
        for _, v := range vert {
            fw.dist[u][v] = Infinity
        }
        fw.dist[u][u] = 0
        for _, v := range g.Neighbors(u) {
            v := v
            fw.dist[u][v] = g.Weight(u, v)
            fw.next[u][v] = &v
        }
    }
    for _, k := range vert {
        for _, i := range vert {
            for _, j := range vert {
                if fw.dist[i][k] < Infinity && fw.dist[k][j] < Infinity {
                    if fw.dist[i][j] > fw.dist[i][k]+fw.dist[k][j] {
                        fw.dist[i][j] = fw.dist[i][k] + fw.dist[k][j]
                        fw.next[i][j] = fw.next[i][k]
                    }
                }
            }
        }
    }
    for _, v := range vert {
        if fw.dist[v][v] < 0 {
            fw.neg = true
        }
    }
    return fw
}
