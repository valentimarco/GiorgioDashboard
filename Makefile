## css: install tailwindcss
.PHONY: css
install-css:
	bun i -D tailwindcss
	bun i -D daisyui@latest
	bun i -D @tailwindcss/forms
	bun i -D @tailwindcss/typography

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	bun run watch

## ln htmx
.PHONY: install-htmx
install-htmx:
	bun i htmx
	mkdir -p static
	ln -s ../node_modules/htmx.org/dist ./static/htmxlib

## ln jquery
.PHONY: install-jquery
install-jquery:
	bun i jquery
	mkdir -p static
	ln -s ../node_modules/jquery/dist ./static/jquery