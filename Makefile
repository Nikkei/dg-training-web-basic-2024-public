.PHONY: dev
dev_all:
	docker compose build && docker compose up

.PHONY: dev_first
dev_first:
	docker compose build tcp_server proxy_server simple_website && docker compose up tcp_server proxy_server simple_website
