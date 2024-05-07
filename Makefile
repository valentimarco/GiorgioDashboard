export UID=$(id -u)
export GID=$(id -g)
## css: install tailwindcss
.PHONY: css
install-css:
	bun i -D tailwindcss
	bun i -D daisyui@latest
	bun i -D @tailwindcss/forms
	bun i -D @tailwindcss/typography

## install js modules and copy htmx and jquery
.PHONY: install-htmx
install-js:
	bun i
	mkdir -p static
	cp -i ./node_modules/htmx.org/dist/htmx.min.js ./static/libraries/htmx.js
	cp -i ./node_modules/jquery/dist/jquery.min.js ./static/libraries/jquery.js


## install css and js libraries
.PHONY: install
install: install-css install-js
	

## docker compose build
.PHONY: docker-build
docker-build:
	docker-compose build
