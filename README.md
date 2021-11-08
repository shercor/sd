# Tarea 2 Sistemas Distribuidos - Squid Game

- Álvaro Fuentes (Par. 200) 201611539-0
- Christian Trujillo (Par. 201) 201673582-8
- Sebastián Herrera (Par. 200) 201551551-4

Para ejecutar los procesos con 16 bots, en cada maquina: 
- Navegar a carpeta /sd
- Ejecutar comando make dependiendo del tipo de proceso

Formato: Proceso: Maquina - IP 

## Lider: dist113 - 10.6.43.101

```
make lider_full
```

## DataNode1: dist113 - 10.6.43.101

```
make datanode
```

## Pozo: dist114 -10.6.43.102

```
make pozo
```

## DataNode2: dist114 - 10.6.43.102

```
make datanode
```

## DataNode3:  dist115 -10.6.43.103

```
make datanode
```

## NameNode: dist116 -10.6.43.104

```
make namenode
```

## Jugadores: dist116 -10.6.43.104 (16 jugadores bot)

```
make jugadores_full
```

Para ejecutar con 15 bots y un jugador por consola, no ejecutar jugadores_full
y ejecutar lo siguiente:

## Jugadores: dist116 -10.6.43.104 (15 jugadores bot)

```
make jugadores
```

## Jugadores: dist116 -10.6.43.104 (1 jugador por consola)

```
make jugador
```