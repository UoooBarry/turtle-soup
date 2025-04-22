watch:
	docker-compose up --watch
sh:
	docker-compose run --rm app sh
go-download:
	docker-compose run --rm app go mod donwload
