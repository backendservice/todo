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
	docker run -it --rm -e AWS_ACCESS_KEY_ID=AKIAINVNBB4S7ZT6N76Q -e AWS_SECRET_ACCESS_KEY=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 -p 50080:50080 todos