build:
	docker build -t taxicalculationservice .
run:
	docker run -p 40001:40001  -e DB_NAME=taxi_service -e DB_USER=taxi -e DB_PASSWORD=a982lsnn -e GEO_SERVICE_HOST=localhost -e GEO_SERVICE_PORT=40000 -e JAEGER_HOST=localhost -e JAEGER_PORT=6831 -e GOOGLE_MAPS_API_KEY=AIzaSyDoVIMRxxJkfA8--jFqaUoBrjTirGwDEzg taxicalculationservice
