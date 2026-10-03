.PHONY: boid filesearch winddirection

boid:
	go -C ./boid-simulation run .

filesearch:
	go -C ./filesearch run .

winddirection:
	go -C ./winddirection run .
