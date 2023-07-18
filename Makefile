# to run docker
run:
	docker build -t ascii .
	docker run -dp 8081:8081 ascii:latest
	$(info running on http://localhost:8081)

# to save on git
s:                     
	git add .
	git commit -m "comment"
	git push