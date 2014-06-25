all:
			go tool yacc -o front/mutan.go front/mutan.y
			go test
