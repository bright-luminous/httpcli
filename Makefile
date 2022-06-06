all: run

run:
	go run main.go "http://localhost:3000" --query internal --query personal --header user=hellichoper

get:
	go run main.go get "http://localhost:3000"

post:
	go run main.go post "http://localhost:3000" --json='{"task": "finish the assigment", "description": "by tmr maybe?"}'

put:
	go run main.go put "http://localhost:3000" --json="{'task': 'go back to work', 'description': 'using bus'}" --ID 225d692a-1d43-4ef8-ac48-056fa71e3531

delete:
	go run main.go delete "http://localhost:3000" --ID 01f87aaf-cb92-4e74-bf6a-6c2f4687e532