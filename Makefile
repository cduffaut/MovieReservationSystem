rundb: 
	docker run --name testdb -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root -e POSTGRES_DB=root -d -p 5432:5432 postgres

destroydb: 
	docker stop testdb && docker rm testdb