package pipline

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	startPipCase()
	//waitRoutine := sync.WaitGroup{}
	//waitRoutine.Add(1)
	//waitRoutine.Wait()
	time.Sleep(time.Second * 100)
}

func pip1(arg interface{}) interface{} {
	log.Println("P1 get : ", arg)
	s := "I'm pip1 return " + arg.(string)
	return s
}

func pip2(arg interface{}) interface{} {
	log.Println("P2 get : ", arg)
	s := "I'm pip2 return " + arg.(string)
	time.Sleep(time.Second * 10)
	return s
}

func pip3(arg interface{}) interface{} {
	log.Println("P3 get : ", arg)
	if arg == "timeout" {
		arg = "P3 is timeout"
		log.Println("P3 timeout, do some ")
	}

	s := "I'm pip3 return " + arg.(string)
	return s
}

func pip4(arg interface{}) interface{} {
	log.Println("P4 get : ", arg)
	if arg == "timeout" {
		//log.Println("P3 get : ", timeout)
	}

	s := "I'm pip4 return " + arg.(string)
	return s
}
func createPip() (txPip Pipeline) {
	pipNodeSlice := make([]*Node, 0)
	pipNodeSlice = append(pipNodeSlice, &Node{Target: pip1, Name: "pip1", Capacity: 5})
	pipNodeSlice = append(pipNodeSlice, &Node{Target: pip2, Name: "pip2", RoutineNum: 2})
	pipNodeSlice = append(pipNodeSlice, &Node{Target: pip3, Name: "pip3", Timeout: 1})
	pipNodeSlice = append(pipNodeSlice, &Node{Target: pip4, Name: "pip4", RoutineNum: 2})
	txPip = Pipeline{
		Nodes: pipNodeSlice,
	}
	return txPip
}

func startPipCase() {
	txPip := createPip()
	indata := startProduceData()
	outData := startProcessData()
	txPip.Setup(indata, outData)
	//txPip.Setup(indata, nil)
	txPip.Start()

	//waitRoutine := sync.WaitGroup{}
	//waitRoutine.Add(1)
	//waitRoutine.Wait()
}

type preNode struct {
	node Node
	note string
}

func produceData(node *Node) {
	//note you can init some datas before start produce
	for i := 0; i < 20; i++ {
		s := "produce data : " + strconv.Itoa(i)
		log.Println(s)
		node.Output <- s
		time.Sleep(time.Second)
	}
}
func startProduceData() *Node {
	preNode := &Node{
		Name: "just do nothing",
	}
	go produceData(preNode)
	return preNode
}

type afterNode struct {
	node   Node
	messge string
}

func processResult(node *Node) {
	for {
		s := <-node.Input
		log.Println("get final data : ", s)
		time.Sleep(time.Second)
	}
}
func startProcessData() *Node {
	afterNode := &Node{
		Name: "just do nothing",
	}
	go processResult(afterNode)
	return afterNode
}
