docker_build:
	docker build -t todoapp .

docker_run:
	docker run -it -p 8000:8000 -v $(pwd):/app todoapp
