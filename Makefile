## css: install tailwindcss
.PHONY: css
install-css:
	npm install -D tailwindcss
	npm i -D daisyui@latest
	npm i -D @tailwindcss/forms
	npm i -D @tailwindcss/typography

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	npx tailwindcss -i ./css/input.css -o ./css/output.css --watch

## ln htmx
.PHONY: install-htmx
install-htmx:
	npm i htmx
	mkdir static
	cd static
	ln -s ../node_modules/htmx.org/dist htmxlib