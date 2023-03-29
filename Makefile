build:
	docker build -t gagankshetty/arxiv-insanity:backend-service .
run:
	docker run -it -p 8080:8080 gagankshetty/arxiv-insanity:backend-service
upload:
	docker push gagankshetty/arxiv-insanity:backend-service