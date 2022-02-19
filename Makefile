.PHONY: test clean

test: deploy.done
	curl -fsSL -D - "$$(terraform output -raw url)?name=Lambda"

clean:
	terraform destroy
	rm -f init.done deploy.done ip_info.zip ip_info

init.done:
	terraform init
	touch $@

deploy.done: init.done main.tf ip_info.zip
	terraform apply
	touch $@

ip_info.zip: ip_info
	zip $@ $<

ip_info: main.go
	go get .
	GOOS=linux GOARCH=amd64 go build -o $@