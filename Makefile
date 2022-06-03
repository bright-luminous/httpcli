run:
	go run main.go "http://localhost:3000"
	go run main.go get "http://localhost:3000"
	go run main.go post "http://localhost:3000" --json="{'task': 'finish the assigment', 'description': 'by tmr maybe?'}"
	go run main.go put "http://localhost:3000" --json="{'task': 'go back to work', 'description': 'using bus'}"
	go run main.go delete "http://localhost:3000" --ID 24cdbf0a-d034-4d49-99c8-037efdf0d14d