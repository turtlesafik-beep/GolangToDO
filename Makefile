include .env
export


export PROJECT_ROOT=$(shell pwd)


env-up:
	docker compose up -d todoapp-postgres

env-down:
	docker compose down todoapp-postgres
	
env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres && \
		sudo rm -rf out/pgdata && \
		echo "Файлы окружения очищины"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутсвует необходисый параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi;	\
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"
