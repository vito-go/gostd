check:
	@go build -o /tmp/gostd-server ./cmd/server && rm /tmp/gostd-server
fmt:
	@gofmt -w -s ./
acp:fmt check
ifndef m
	@$(error error: 需要提交说明 请指定参数m, 例如 make acp m=fix)
else
	git add . && git commit -m '$(m)'  && git push
endif
wire:
	cd logic/wireinject && wire
run:
	go run ./cmd/server -out