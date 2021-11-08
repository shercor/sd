hello:
	echo "Hello"

jugadores:			
	for i in 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15; do \
        go run jugador.go bot & \
	done 

jugador: 
	go run jugador.go

jugadores_full:			
	for i in 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16; do \
        go run jugador.go bot & \
	done 

lider: 
	go run lider.go 5

lider_full: 
	go run lider.go 16

pozo:
	go run pozo.go

datanode:
	go run datanode.go

namenode:
	go run namenode.go

clean:
	rm *.txt