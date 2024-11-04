clean:
	rm -rf foobar bin/gdextension

build: clean
	go build -o bin/gdextension main.go

run: build
	./bin/gdextension

create: build
	./bin/gdextension -no-clean -all -jobs 18 foobar gd421

create43: build
	./bin/gdextension -no-clean -all -jobs 18 foobar gd43

