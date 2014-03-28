all:
			go tool yacc -o mutan.go mutan.y
			go build
