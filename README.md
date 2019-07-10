# Multi-Layered Perceptron
Learning `Go` programming language by re-implementing a `Connectionist Computing` assignment. 

# Description
The software should be able to:

··* Create a new MLP with any given number of inputs, any number of outputs (can be sigmoidal or linear), and any number of hidden units in a single layer.
··* Initialise the weights of the MLP to small random values
··* Predict the outputs corresponding to an input vector
··* Implement learning by backpropagation

# Tasks:

## XOR
Train an MLP with __2 inputs__, __two hidden units__ and __one output__ on the following examples (XOR function): 
__inputs__: [x1 x2] (1,0)
__output__: XOR(x1, x2)

### Example XOR: 
··* ((x1, x2), o1)
··* ((0, 0), 0) ((0, 1), 1) ((1, 0), 1) ((1, 1), 0)
At the end of training, check if the MLP predicts correctly all the examples.

## Sin(x)
Generate 50 vectors containing 4 components each. The value of each component should be a random number between -1 and 1. These will be your input vectors. The corresponding output for each vector should be the sin() of the sum of the components. 
__inputs__: [x1 x2 x3 x4] the (single component) 
__output__ : sin(x1+x2+x3+x4) 

Now train an MLP with 4 inputs, at least 5 hidden units and one output on 40 of these examples and keep the remaining 10 for testing.

## Image Recongnition

Train an MLP on the letter recognition set available in the UCI Machine Learning repository: http://archive.ics.uci.edu/ml/machine-learning-databases/letterrecognition/letter-recognition.data. 
The first entry of each line is the letter to be recognised (i.e. the target) and the following numbers are attributes extracted from images of the letters (i.e. your input). 
You can find a description of the set here: http://archive.ics.uci.edu/ml/datasets/Letter+Recognition

Split the dataset in a training part containing approximately 4/5 of the records, and a testing part containing the rest. Your MLP should have as many inputs as there are attributes (17), as many hidden units as you want (I suggest ~10) and 26 outputs (one for each letter of the alphabet).
You should train your MLP for __at least 1000 epochs__. After training, check how well you can classify the data reserved for testing. 

# Methods: 
1. randomise() 
    Initialises W1 and W2 to small random values. Also don’t forget to set dW1 and dW2 to all zeroes – this could be a good place to do it. 
2. forward(double* I) Forward pass. Input vector I is processed to produce an output, which is stored in O[]. 
3. double backwards(double* t) Backwards pass. Target t is compared with output O, deltas are computed for the upper layer, and are multiplied by the inputs to the layer (the values in H) to produce the weight updates which are stored in dW2 (added to it, as you may want to store these for many examples). Then deltas are produced for the lower layer, and the same process is repeated here, producing weight updates to be added to dW1. Returns the error on the example. 
4. updateWeights(double learningRate) this simply does (component by component, i.e. within for loops): W1 += learningRatedW1; W2 += learningRatedW2; dW1 = 0; dW2 = 0; 

 Training would proceed along these lines (pseudo-C code): 
 ```C++
    for (int e=0; e<maxEpochs; e++) 
        { 
            error = 0; 
            for (int p=0; p< numExamples; p++) { 
                NN.forward(example[p].input); error += NN.backwards(example[p].output); 
                every now and then { 
                    updateWeights(some_small_value); 
                } 
            } 
            cout << “Error at epoch “ << e << “ is “ << error << “\n”; 
        } 
```

The error that is output should, ideally, get smaller at every epoch. You may have to try different learning rates (too big and training will explode, too small and learning will be very slow), and different “every now and then” – to play safe you can even do the update only once every epoch

# Results