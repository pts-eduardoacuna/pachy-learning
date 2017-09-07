all: compile docker push clean

compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o image .

docker:
	docker build --no-cache=true --force-rm=true -t eduardoacye/pachy-learning-image .

push:
	docker push eduardoacye/pachy-learning-image

clean:
	rm image
