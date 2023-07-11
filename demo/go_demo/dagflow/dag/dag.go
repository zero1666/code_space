package dag

import "fmt"

const (
	STATE_UNKNOWN = 0
	STATE_INIT    = 1
	STATE_READY   = 2
	STATE_RUNNING = 3
	STATE_DONE    = 4
)

type DAG struct {
	Vertexes  []*Vertex
	mapVertex map[string]*Vertexes
}

type Vertex struct {
	Name     string
	State    int
	Parents  []*Vertex
	Children []*Vertex
	f        func()
	result   interface{}
	err      error
}

func (v *Vertex) Run() {
	v.State = STATE_RUNNING
	f()
	v.State = STATE_DONE
}

func NewDag() *DAG {
	dag := &DAG{}
	dag.mapVertex = make(map[string]*Vertex, 0)
	dag.Vertexes = make([]*Vertex, 0, 0)

	return dag
}

func NewVertex(name string, f func()) *Vertex {
	return &Vertex{Name: name, f: f, State: STATE_INIT}
}

func (dag *DAG) AddVertex(v *Vertex) {
	dag.Vertexes = append(dag.Vertexes, v)
	dag.mapVertex[v.Name] = v
}

func (dag *DAG) AddEdge(from, to *Vertex) {
	from.Children = append(from.Children, to)
	to.Parents = append(to.Parents, from)
}

// build dag
func (dag *DAG) BuildState() {
	for _, v := range dag.Vertexes {
		if v.State == STATE_RUNNING || v.State == STATE_DONE {
			continue
		}

		state := STATE_READY
		for _, parent := range v.Parents {
			if parent.State != STATE_DONE {
				state = STATE_INIT
				break
			}
		}
		if state == STATE_READY {
			v.State = STATE_READY
		}
	}
}

// 并发运行
func (dag *DAG) Run() {
	// 1. BuildState
	// 2. 筛选 ready
	// 3. 加入并发

	//// 任何一个任务执行结束,重新执行1-3

	////等所有并发执行完毕。则整个流程执行完毕
}

// DAG图是否有环检测
func (dag *DAG) CheckCircle() {

}
