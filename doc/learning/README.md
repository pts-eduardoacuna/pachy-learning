

# learning
`import "github.com/pts-eduardoacuna/pachy-learning/learning"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package learning allows training and inference with ML models.




## <a name="pkg-index">Index</a>
* [func ArgMax(xs []float64) int](#ArgMax)
* [func ComputeError(net *NeuralNetwork, inputs, expected *mat.Dense) (float64, error)](#ComputeError)
* [func DecodeTargets(targets []float64) float64](#DecodeTargets)
* [func EncodeAttributes(xs []float64) []float64](#EncodeAttributes)
* [func EncodeTarget(target float64) []float64](#EncodeTarget)
* [func Infer(net *NeuralNetwork, attributesSet *mat.Dense) (*mat.Dense, error)](#Infer)
* [func SplitTrainingValidation(validationPercent int, inputs, outputs *mat.Dense) (*mat.Dense, *mat.Dense, *mat.Dense, *mat.Dense)](#SplitTrainingValidation)
* [func Train(net *NeuralNetwork, attributesSet, targetsSet *mat.Dense) error](#Train)
* [func ValidateNeuralNetwork(tInputs, tOutputs, vInputs, vOutputs *mat.Dense, rate float64, arch []int) (float64, float64, error)](#ValidateNeuralNetwork)
* [func WriteAnalysis(analysis *Analysis, path string) error](#WriteAnalysis)
* [type Analysis](#Analysis)
  * [func NewAnalysis(validationData []AnalysisValidation, selected Model) *Analysis](#NewAnalysis)
  * [func ReadAnalysis(path string) (*Analysis, error)](#ReadAnalysis)
* [type AnalysisValidation](#AnalysisValidation)
* [type AnalysisValidationResult](#AnalysisValidationResult)
* [type Model](#Model)
  * [func NewModel(rate float64, arch []int) Model](#NewModel)
  * [func SelectBestModel(validationData []AnalysisValidation) (Model, error)](#SelectBestModel)
* [type NeuralNetwork](#NeuralNetwork)
  * [func NewNeuralNetwork(learningRate float64, arch []int) (*NeuralNetwork, error)](#NewNeuralNetwork)


#### <a name="pkg-files">Package files</a>
[ann.go](/src/github.com/pts-eduardoacuna/pachy-learning/learning/ann.go) [encoding.go](/src/github.com/pts-eduardoacuna/pachy-learning/learning/encoding.go) [model.go](/src/github.com/pts-eduardoacuna/pachy-learning/learning/model.go) [validation.go](/src/github.com/pts-eduardoacuna/pachy-learning/learning/validation.go) 





## <a name="ArgMax">func</a> [ArgMax](/src/target/encoding.go?s=885:914#L23)
``` go
func ArgMax(xs []float64) int
```
ArgMax takes an array of numbers and returns the index of the maximum value.



## <a name="ComputeError">func</a> [ComputeError](/src/target/validation.go?s=1880:1963#L54)
``` go
func ComputeError(net *NeuralNetwork, inputs, expected *mat.Dense) (float64, error)
```
ComputeError checks how many unsuccessful predictions the neural network makes.



## <a name="DecodeTargets">func</a> [DecodeTargets](/src/target/encoding.go?s=721:766#L18)
``` go
func DecodeTargets(targets []float64) float64
```
DecodeTargets takes an array of numbers and returns the index of the maximum value.



## <a name="EncodeAttributes">func</a> [EncodeAttributes](/src/target/encoding.go?s=168:213#L1)
``` go
func EncodeAttributes(xs []float64) []float64
```
EncodeAttributes applies a hard threshold on the [0,255] range of values of an array such that
any number other than 0 is transformed into a 1.



## <a name="EncodeTarget">func</a> [EncodeTarget](/src/target/encoding.go?s=489:532#L9)
``` go
func EncodeTarget(target float64) []float64
```
EncodeTarget takes a number between [0,9] and computes an array of zeros with a 1 in the position
of the input.



## <a name="Infer">func</a> [Infer](/src/target/ann.go?s=4073:4149#L120)
``` go
func Infer(net *NeuralNetwork, attributesSet *mat.Dense) (*mat.Dense, error)
```
Infer user the network to evaluate each row in the attributes dataset.



## <a name="SplitTrainingValidation">func</a> [SplitTrainingValidation](/src/target/validation.go?s=179:307#L1)
``` go
func SplitTrainingValidation(validationPercent int, inputs, outputs *mat.Dense) (*mat.Dense, *mat.Dense, *mat.Dense, *mat.Dense)
```
SplitTrainingValidation chooses a random sample of training attributes for constructing a validation set.



## <a name="Train">func</a> [Train](/src/target/ann.go?s=2643:2717#L78)
``` go
func Train(net *NeuralNetwork, attributesSet, targetsSet *mat.Dense) error
```
Train adjusts the parameters of a neural network to fit the attributes dataset with
the targets dataset.

Both datasets must have the same number of rows, and their columns should match the
dimension of the first and last layer of the network.



## <a name="ValidateNeuralNetwork">func</a> [ValidateNeuralNetwork](/src/target/validation.go?s=1291:1418#L29)
``` go
func ValidateNeuralNetwork(tInputs, tOutputs, vInputs, vOutputs *mat.Dense, rate float64, arch []int) (float64, float64, error)
```
ValidateNeuralNetwork runs a simple training and inference check on a training and validation set.



## <a name="WriteAnalysis">func</a> [WriteAnalysis](/src/target/model.go?s=2406:2463#L71)
``` go
func WriteAnalysis(analysis *Analysis, path string) error
```
WriteAnalysis writes a neural network analysis in a JSON encoding to a file.




## <a name="Analysis">type</a> [Analysis](/src/target/model.go?s=368:481#L5)
``` go
type Analysis struct {
    Data []AnalysisValidation `json:"validation"`
    Best Model                `json:"model"`
}
```
Analysis holds the results of training a variety of neural network architectures with different learning rates.

It holds the result of the best architecture and learning rate according to a validation analysis, and also
the data associated with each model that was analyzed.







### <a name="NewAnalysis">func</a> [NewAnalysis](/src/target/model.go?s=579:658#L11)
``` go
func NewAnalysis(validationData []AnalysisValidation, selected Model) *Analysis
```
NewAnalysis creates an Analysis reference from the given validation data and selected model.


### <a name="ReadAnalysis">func</a> [ReadAnalysis](/src/target/model.go?s=2775:2824#L91)
``` go
func ReadAnalysis(path string) (*Analysis, error)
```
ReadAnalysis reads a neural network analysis in a JSON encoding from a file.





## <a name="AnalysisValidation">type</a> [AnalysisValidation](/src/target/model.go?s=1180:1334#L33)
``` go
type AnalysisValidation struct {
    LearningRate float64                    `json:"learningRate"`
    Models       []AnalysisValidationResult `json:"models"`
}
```
AnalysisValidation holds the validation results associated with a learning rate.










## <a name="AnalysisValidationResult">type</a> [AnalysisValidationResult](/src/target/model.go?s=1450:1635#L39)
``` go
type AnalysisValidationResult struct {
    Architecture    []int   `json:"architecture"`
    TrainingError   float64 `json:"trainingError"`
    ValidationError float64 `json:"validationError"`
}
```
AnalysisValidationResult hold the training and validation error associated with a neural network architecture.










## <a name="Model">type</a> [Model](/src/target/model.go?s=782:891#L19)
``` go
type Model struct {
    LearningRate float64 `json:"learningRate"`
    Architecture []int   `json:"architecture"`
}
```
Model describes a description of a neural network.







### <a name="NewModel">func</a> [NewModel](/src/target/model.go?s=983:1028#L25)
``` go
func NewModel(rate float64, arch []int) Model
```
NewModel creates a neural network Model from the given learning rate and architecture.


### <a name="SelectBestModel">func</a> [SelectBestModel](/src/target/model.go?s=1757:1829#L46)
``` go
func SelectBestModel(validationData []AnalysisValidation) (Model, error)
```
SelectBestModel receives a slice of AnalysisValidation and returns the model which has the minimum validation error.





## <a name="NeuralNetwork">type</a> [NeuralNetwork](/src/target/ann.go?s=227:547#L4)
``` go
type NeuralNetwork struct {
    Signals         []*mat.Dense
    Outputs         []*mat.Dense
    Weights         []*mat.Dense
    Deltas          []*mat.Dense
    Gradients       []*mat.Dense
    ErrorHistory    *mat.Dense
    LayerCount      int
    LearningRate    float64
    AttributesSize  int
    PredictionsSize int
    TestingError    float64
}
```
NeuralNetwork is a supervised learning model for classification.







### <a name="NewNeuralNetwork">func</a> [NewNeuralNetwork](/src/target/ann.go?s=850:929#L23)
``` go
func NewNeuralNetwork(learningRate float64, arch []int) (*NeuralNetwork, error)
```
NewNeuralNetwork creates a NeuralNetwork with the given learning rate and architecture.

The architecture consists of at least two elements, where each element specifies the amount of
nodes in the layers. The first number corresponds to the input layer and the last to the output
layer.









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
