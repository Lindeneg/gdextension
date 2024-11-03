build:
	go build -o bin/gdextension *.go

run: build
	./bin/gdextension

create: build
	rm -rf foobar && ./bin/gdextension -all -jobs 18 foobar gd421

