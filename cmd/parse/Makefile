all: compile docker push clean

compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o parse .

docker:
	docker build --no-cache=true --force-rm=true -t eduardoacye/pachy-learning-parse .

push:
	docker push eduardoacye/pachy-learning-parse

clean:
	rm parse

local:
	go build -o parse .
