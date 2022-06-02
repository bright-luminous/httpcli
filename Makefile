run:
	go run main.go "http://localhost:3000/"
	go run main.go get "http://localhost:3000/todos"
	go run main.go post "http://localhost:3000/todos" --json-task="finish this game" --json-description="in 2 day"
	# go run main.go put
	# go run main.go delete