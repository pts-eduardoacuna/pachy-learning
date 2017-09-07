all: compile docker push clean

compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o train .

docker:
	docker build --no-cache=true --force-rm=true -t eduardoacye/pachy-learning-train .

push:
	docker push eduardoacye/pachy-learning-train

clean:
	rm train
