default:
	# This is my 2021 Advent of Code project

init:
	mkdir day-$(day)
	sed 's/#DAY#/$(day)/g' main.go.template | sed 's/#TITLE#/$(title)/g' > day-$(day)/main.go
	sed 's/#DAY#/$(day)/g' main_test.go.template > day-$(day)/main_test.go
	touch day-$(day)/input.txt;

build:
	cd day-$(day); go run . log=$(log) part=$(part)

test:
	cd day-$(day); go test

test-utils:
	go test ./utils/*