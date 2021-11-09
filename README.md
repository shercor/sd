# Tarea 2 Sistemas Distribuidos - Squid Game

- Álvaro Fuentes (Par. 200) 201611539-0
- Christian Trujillo (Par. 201) 201673582-8
- Sebastián Herrera (Par. 200) 201551551-4

Versión GO: go1.17.3 linux/amd64 (ya esta actualizado en VM's)

NOTA IMPORTANTE: La comunicación asíncrona con el pozo funciona (tecnología RabbitMQ implementada),
pero detectamos que produce un mal funcionamiento de la Etapa 1. Tras horas de pruebas se decide  
"comentar" el contenido de la función aumentarPozo de lider.go para efectos de revisar todas las etapas. 

Si se descomenta el contenido de la función aumentarPozo de lider.go se puede apreciar el correcto funcionamiento de RabbitMQ durante 
las primeras rondas. 

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

## Jugadores: dist116 -10.6.43.104 (16 jugadores bot) (OBSERVACIÓN: demora en ejecutarse)

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