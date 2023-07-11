package main

import (
	"fmt"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func main() {
	s := op.NewScope()
	c := op.Const(s, "Hell;o from Tensorflow version "+tf.Version())
	graph, err := s.Finalize()
	if err != nil {
		panic(err)
	}

	//Execute the graph in a session
	output, err := sess.Run(nil, []tf.Output{c}, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("result: " + output[0].Value())
}
