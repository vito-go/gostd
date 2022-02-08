init:
	@git config --global url."http://vitogo.tpddns.cn:9000".insteadOf "https://gitea.com"
check:
	go build -o /tmp/openblog ./cmd/openblog && rm /tmp/openblog
	go build -o /tmp/user ./cmd/user && rm /tmp/user
fmt:
	@gofmt -w -s ./
acp:fmt check
ifndef m
	@$(error error: 需要提交说明 请指定参数m, 例如 make acp m=fix)
else
	git add . && git commit -m '$(m)'  && git push origin --all
endif
wire:
	cd wireinject && wire
run-openblog:
	go run ./cmd/openblog
run-user:
	go run ./cmd/user
lint:
	@golangci-lint run