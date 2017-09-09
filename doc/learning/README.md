

# learning
`import "./learning"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func Infer(net *NeuralNetwork, attributesSet *mat.Dense) *mat.Dense](#Infer)
* [func Train(net *NeuralNetwork, attributesSet, targetsSet *mat.Dense)](#Train)
* [type NeuralNetwork](#NeuralNetwork)
  * [func NewNeuralNetwork(learningRate float64, arch []int) *NeuralNetwork](#NewNeuralNetwork)


#### <a name="pkg-files">Package files</a>
[ann.go](/src/target/ann.go) 





## <a name="Infer">func</a> [Infer](/src/target/ann.go?s=2534:2601#L82)
``` go
func Infer(net *NeuralNetwork, attributesSet *mat.Dense) *mat.Dense
```
Infer is a function :)



## <a name="Train">func</a> [Train](/src/target/ann.go?s=1641:1709#L53)
``` go
func Train(net *NeuralNetwork, attributesSet, targetsSet *mat.Dense)
```
Train is a function :)




## <a name="NeuralNetwork">type</a> [NeuralNetwork](/src/target/ann.go?s=113:436#L1)
``` go
type NeuralNetwork struct {
    Signals      []*mat.Dense
    Outputs      []*mat.Dense
    Weights      []*mat.Dense
    Deltas       []*mat.Dense
    Gradients    []*mat.Dense
    ErrorHistory *mat.Dense
    LayerCount   int
    LearningRate float64
    // contains filtered or unexported fields
}
```
NeuralNetwork is a structure :)







### <a name="NewNeuralNetwork">func</a> [NewNeuralNetwork](/src/target/ann.go?s=475:545#L16)
``` go
func NewNeuralNetwork(learningRate float64, arch []int) *NeuralNetwork
```
NewNeuralNetwork is a function :)









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)