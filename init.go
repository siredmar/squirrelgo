package main

func initEntities(board *Board) {

	var i int

	for i = 0; i < boardX; i++ {
		board.AddEntity(createWall(xy{i, 0}), xy{i, 0})
		board.AddEntity(createWall(xy{i, boardY - 1}), xy{i, boardY - 1})
	}
	for i = 0; i < boardX; i++ {
		board.AddEntity(createWall(xy{0, i}), xy{0, i})
		board.AddEntity(createWall(xy{boardX - 1, i}), xy{boardX - 1, i})
	}

	board.AddEntity(createWall(xy{5, 1}), xy{5, 1})
	board.AddEntity(createWall(xy{5, 2}), xy{5, 2})
	board.AddEntity(createWall(xy{5, 3}), xy{5, 3})
	board.AddEntity(createWall(xy{5, 4}), xy{5, 4})
	board.AddEntity(createWall(xy{4, 4}), xy{4, 4})
	board.AddEntity(createWall(xy{3, 4}), xy{3, 4})
	board.AddEntity(createWall(xy{17, 8}), xy{17, 8})
	board.AddEntity(createWall(xy{17, 8}), xy{17, 7})
	board.AddEntity(createWall(xy{17, 6}), xy{17, 6})
	board.AddEntity(createWall(xy{17, 5}), xy{17, 5})
	board.AddEntity(createWall(xy{16, 5}), xy{16, 5})
	board.AddEntity(createWall(xy{15, 5}), xy{15, 5})
	board.AddEntity(createWall(xy{14, 4}), xy{14, 5})
	board.AddEntity(createWall(xy{14, 4}), xy{14, 4})
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")

	player := createMasterSquirrel(xy{10, 5})
	board.AddEntity(player, xy{10, 5})
	board.addPlayer(player)
}
