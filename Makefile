server:
	go build -o server *.go
run: server
	AWS_REGION=us-west-2 ./server
test:
	go test
genmock:
	mkdir -p mock
	mockgen -source=services.go -destination mock/services.go
docker:
	docker build -t todos .
run-docker:
	docker run -it --rm -p 50080:50080 todos