.PHONY: boid filesearch

boid:
	go -C ./boid-simulation run .

filesearch:
	go -C ./filesearch run .
