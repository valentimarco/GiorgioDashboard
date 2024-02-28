/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.templ"
  ],
  theme: {
    extend: {},
    fontFamily: {},
  },
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: ["sunset"]
  },
};