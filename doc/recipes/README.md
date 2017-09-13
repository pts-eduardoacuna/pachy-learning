# Repo recipes

## Building the pipeline

### Step 1

#### Create the `mnist` repo

```sh
pachctl create-repo mnist
```

```sh
pachctl start-commit mnist master
```

```sh
cd data
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

#### Create the `mnist-csv` repo

```sh
cd ../spec
```

```sh
pachctl create-pipeline -f parse.json
```
