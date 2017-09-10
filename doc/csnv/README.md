

# csnv
`import "github.com/pts-eduardoacuna/pachy-learning/csnv"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func ReadFloats(r *csv.Reader) ([]float64, error)](#ReadFloats)
* [func ReadInts(r *csv.Reader) ([]int, error)](#ReadInts)
* [func WriteFloats(w *csv.Writer, floats []float64) error](#WriteFloats)
* [func WriteInts(w *csv.Writer, ints []int) error](#WriteInts)


#### <a name="pkg-files">Package files</a>
[numeric.go](/src/github.com/pts-eduardoacuna/pachy-learning/csnv/numeric.go) 





## <a name="ReadFloats">func</a> [ReadFloats](/src/target/numeric.go?s=143:192#L1)
``` go
func ReadFloats(r *csv.Reader) ([]float64, error)
```
ReadFloats is a wrapper around the standard CSV reader for records containing float64.



## <a name="ReadInts">func</a> [ReadInts](/src/target/numeric.go?s=835:878#L26)
``` go
func ReadInts(r *csv.Reader) ([]int, error)
```
ReadInts is a wrapper around the standard CSV reader for records containing int.



## <a name="WriteFloats">func</a> [WriteFloats](/src/target/numeric.go?s=543:598#L16)
``` go
func WriteFloats(w *csv.Writer, floats []float64) error
```
WriteFloats is a wrapper around the standard CSV writer for record containing float64.



## <a name="WriteInts">func</a> [WriteInts](/src/target/numeric.go?s=1219:1266#L43)
``` go
func WriteInts(w *csv.Writer, ints []int) error
```
WriteInts is a wrapper around the standard CSV writer for record containing int.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)