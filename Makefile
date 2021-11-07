hello:
	echo "Hello"
jugadores:			
	for i in 1 2 3 4 5; do \
        gnome-terminal -- go run jugador.go; \
	done 
lider: 
	go run lider.go 5
	