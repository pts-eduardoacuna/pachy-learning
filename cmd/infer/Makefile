all: compile docker push clean

compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o infer .

docker:
	docker build --no-cache=true --force-rm=true -t eduardoacye/pachy-learning-infer .

push:
	docker push eduardoacye/pachy-learning-infer

clean:
	rm infer

local:
	go build -o infer .
