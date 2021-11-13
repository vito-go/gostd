check:
	@go build -o /tmp/gostd-server ./cmd/server && rm /tmp/gostd-server
fmt:
	@gofmt -w -s ./
# acp: 三个git指令一气呵成: add commit push push前进行格式化代码代码,并进行编译性通过审查
acp:fmt check
ifndef m
	@$(error error: 需要提交说明 请指定参数m, 例如 make acp m=fix)
else
	git add . && git commit -m '$(m)'  && git push origin --all
endif
wire:
	cd logic/wireinject && wire
run:
	go run ./cmd/server
lint:
	@golangci-lint run