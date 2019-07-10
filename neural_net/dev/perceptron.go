package main

import ("time"
 		"fmt"
		"math/rand"
		)

type Neuron struct {
	weights []float64
	bias float64
	learning_rate float 64
	input float64 
	deltas []float64
}

type Layer struct {
	neurons []Neuron
	size int
}

func forward(input_vector *float64){

}

func backward(){

}

func update_weights(n Neuron) {
	for(i := 0; i< len(n.weights); i++){
		n.weights[i] += n.learning_rate
		n.deltas[i] := 0
	}
}


func init_weights(n Neuron, hidden_size) {
	for(i := 0; i < hidden_size; i++){
		n.weights[i] := rand.float64() * n.bias
		n.deltas[i] := 0
	}
}

func activate(){

}