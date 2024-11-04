build:
	go build -o bin/gdextension main.go

run: build
	./bin/gdextension

create: build
	rm -rf foobar && ./bin/gdextension -all -jobs 18 foobar gd421

