build:
	docker build -t arxiv-insanity/backend-service .
run:
	docker run -it -p 8080:8080 arxiv-insanity/backend-service