package datastructure

type BFSPath struct {
	Source int
	Marked map[int]bool
	EdgeTo map[int]int
	G      *Graph
}

func NewBFSPath(g *Graph, source int) *BFSPath {
	bfsPath := &BFSPath{
		Marked: make(map[int]bool),
		EdgeTo: make(map[int]int),
		G:      g,
		Source: source,
	}
	bfs(source)
	return bfsPath
}

func (b *BFSPath) bfs(v int) {
	queue := NewQueueLinkedList()
	b.Marked[v] = true
	queue.Enqueue(v)
	for {
		if queue.IsEmpty() {
			break
		}
		d := queue.Dequeue()
		for r := range b.G.Adj(d) {
			if !b.Marked[r] {
				queue.Enqueue(r)
				b.EdgeTo[r] = d
				b.Marked[r] = true
			}
		}
	}
}
