build_image:
	docker build . -t tr1stanzhi/prometheus-demo:v1

push_image:
	docker push tr1stanzhi/prometheus-demo:v1