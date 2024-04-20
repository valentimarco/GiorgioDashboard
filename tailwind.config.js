/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.templ"
  ],
  theme: {
    extend: {},
    fontFamily: {},
  },
  plugins: [
    require("@tailwindcss/forms"),
    require("@tailwindcss/typography"),
    require("daisyui"),
    require("tailwind-htmx")
  ],
  daisyui: {
    themes: ["sunset"]
  },
};