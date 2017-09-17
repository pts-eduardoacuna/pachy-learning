# Repo recipes

## Building the pipeline

```sh
PACHY_HOME=$GOPATH/src/github.com/pts-eduardoacuna/pachy-learning
```

### Step 1 `mnist`

#### Create an empty `mnist` repo

```sh
pachctl create-repo mnist
```

#### Upload to it the training and testing MNIST binary files

```sh
pachctl start-commit mnist master
```

```sh
cd $PACHY_HOME/data
```

```sh
pachctl put-file mnist <commit-id> train-images-idx3-ubyte -f train-images-idx3-ubyte
pachctl put-file mnist <commit-id> train-labels-idx1-ubyte -f train-labels-idx1-ubyte
pachctl put-file mnist <commit-id> t10k-images-idx3-ubyte -f t10k-images-idx3-ubyte
pachctl put-file mnist <commit-id> t10k-labels-idx1-ubyte -f t10k-labels-idx1-ubyte
```

```sh
pachctl finish-commit mnist <commit-id>
```

### Step 2 `mnist-csv`

#### Create a pipeline using the `parse` program

```sh
cd $PACHY_HOME/spec
```

```sh
pachctl create-pipeline -f parse.json
```

#### Inspect what's going on

```sh
pachctl list-repo
```

```sh
watch -n1 kubectl get all
```

```sh
watch -n1 pachctl list-job
```

```sh
pachctl list-repo
```

#### Trace the computed files to its origins

```sh
pachctl list-commit mnist-csv
```

```sh
pachctl list-file mnist-csv <commit-id>
```

```sh
pachctl inspect-commit mnist-csv <commit-id>
```

### Step 2 `analysis`

#### Create a pipeline using the `stats` program

```sh
cd $PACHY_HOME/spec
```

```sh
pachctl create-pipeline -f stats.json
```
