.PHONY: boid filesearch winddirection threadpool

boid:
	go -C ./boid-simulation run .

filesearch:
	go -C ./filesearch run .

winddirection:
	go -C ./winddirection run .

threadpool:
	go -C ./threadpool run .
