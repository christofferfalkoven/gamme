run_node:
	node app.js
run_go:
	go run board.go human.go main.go server.go virus.go
run:
	make -j2 run_node run_go
