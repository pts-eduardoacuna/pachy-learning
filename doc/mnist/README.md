

# mnist
`import "github.com/pts-eduardoacuna/pachy-learning/mnist"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package mnist allows working with images and labels MNIST binary files.




## <a name="pkg-index">Index</a>
* [type ImageParser](#ImageParser)
  * [func NewImageParser(file *os.File) (*ImageParser, error)](#NewImageParser)
  * [func (p *ImageParser) Parse() ([]int, error)](#ImageParser.Parse)
* [type LabelParser](#LabelParser)
  * [func NewLabelParser(file *os.File) (*LabelParser, error)](#NewLabelParser)
  * [func (p *LabelParser) Parse() (int, error)](#LabelParser.Parse)


#### <a name="pkg-files">Package files</a>
[read.go](/src/github.com/pts-eduardoacuna/pachy-learning/mnist/read.go) 






## <a name="ImageParser">type</a> [ImageParser](/src/target/read.go?s=1113:1213#L50)
``` go
type ImageParser struct {
    Count   int
    Rows    int
    Columns int
    // contains filtered or unexported fields
}
```
ImageParser holds the state of an image parsing process.







### <a name="NewImageParser">func</a> [NewImageParser](/src/target/read.go?s=1278:1334#L59)
``` go
func NewImageParser(file *os.File) (*ImageParser, error)
```
NewImageParser creates an image parser from the given file.





### <a name="ImageParser.Parse">func</a> (\*ImageParser) [Parse](/src/target/read.go?s=1913:1957#L85)
``` go
func (p *ImageParser) Parse() ([]int, error)
```
Parse reads an image from the parser file and returns it as a slice of integers.

The elements of the slice are the image's pixels grayscale values in the range [0,255].




## <a name="LabelParser">type</a> [LabelParser](/src/target/read.go?s=263:331#L6)
``` go
type LabelParser struct {
    Count int
    // contains filtered or unexported fields
}
```
LabelParser holds the state of a label parsing process.







### <a name="NewLabelParser">func</a> [NewLabelParser](/src/target/read.go?s=395:451#L13)
``` go
func NewLabelParser(file *os.File) (*LabelParser, error)
```
NewLabelParser creates a label parser from the given file.





### <a name="LabelParser.Parse">func</a> (\*LabelParser) [Parse](/src/target/read.go?s=804:846#L34)
``` go
func (p *LabelParser) Parse() (int, error)
```
Parse reads a label from the parser file.

The returned value is in the range [0,9].








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
