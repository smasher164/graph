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

import "container/heap"

type dijkstra struct {
	dist  map[Vertex]int
	prev  map[Vertex]Vertex
	items []Vertex
	// value to index
	m map[Vertex]int
	// value to priority
	pr map[Vertex]int
}

func (dj *dijkstra) Distance(u, v Vertex) int { return dj.dist[v] }
func (dj *dijkstra) Path(u, v Vertex) (path []Vertex) {
	if dj.dist[v] == Infinity {
	    return
	}
	path = []Vertex{v}
	for dj.prev[v] >= 0 {
		v = dj.prev[v]
		path = append([]Vertex{v}, path...)
	}
	return path
}
func (dj *dijkstra) NegativeCycle() bool { return false }
func (dj *dijkstra) Len() int            { return len(dj.items) }
func (dj *dijkstra) Less(i, j int) bool  { return dj.pr[dj.items[i]] < dj.pr[dj.items[j]] }
func (dj *dijkstra) Swap(i, j int) {
	dj.items[i], dj.items[j] = dj.items[j], dj.items[i]
	dj.m[dj.items[i]] = i
	dj.m[dj.items[j]] = j
}
func (dj *dijkstra) Push(x interface{}) {
	n := len(dj.items)
	item := x.(Vertex)
	dj.m[item] = n
	dj.items = append(dj.items, item)
}
func (dj *dijkstra) Pop() interface{} {
	old := dj.items
	n := len(old)
	item := old[n-1]
	dj.m[item] = -1
	dj.items = old[0 : n-1]
	return item
}
func (dj *dijkstra) update(item Vertex, priority int) {
	dj.pr[item] = priority
	heap.Fix(dj, dj.m[item])
}
func (dj *dijkstra) addWithPriority(item Vertex, priority int) {
	heap.Push(dj, item)
	dj.update(item, priority)
}

// Single finds the shortest path from source to every
// vertex in a graph.
func Single(g Graph, source Vertex) ShortestPath {
	q := &dijkstra{
		dist:  make(map[Vertex]int),
		prev:  make(map[Vertex]Vertex),
		items: []Vertex{},
		m:     make(map[Vertex]int),
		pr:    make(map[Vertex]int),
	}
	q.dist[source] = 0
	for _, v := range g.Vertices() {
		if v != source {
			q.dist[v] = Infinity
		}
		q.prev[v] = Uninitialized
		q.addWithPriority(v, q.dist[v])
	}
	for len(q.items) != 0 {
		u := heap.Pop(q).(Vertex)
		for _, v := range g.Neighbors(u) {
			alt := q.dist[u] + g.Weight(u, v)
			if alt < q.dist[v] {
				q.dist[v] = alt
				q.prev[v] = u
				q.update(v, alt)
			}
		}
	}
	return q
}
