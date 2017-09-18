all: compile docker push clean

compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o stats .

docker:
	docker build --no-cache=true --force-rm=true -t eduardoacye/pachy-learning-stats .

push:
	docker push eduardoacye/pachy-learning-stats

clean:
	rm stats

local:
	go build -o stats .
