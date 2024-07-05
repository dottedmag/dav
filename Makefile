APP=dav
BIN=$(APP).linux-amd64
GB=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
H=rooibos

build:
	$(GB) -o $(BIN) .

deploy: build
	ssh $H mkdir -p $(APP) .config/systemd/user
	rsync $(BIN) $H:$(APP)/$(APP)
	rsync $(APP).service $H:.config/systemd/user

	ssh $H systemctl --user daemon-reload
	ssh $H systemctl --user enable $(APP).service
	ssh $H systemctl --user restart $(APP).service

clean:
	rm -f $(BIN)

