hello:
	echo "Hello"
jugadores:			
	for i in 1 2 3 4 5; do \
        go run jugador.go bot & \
	done 

lider: 
	go run lider.go 5

	