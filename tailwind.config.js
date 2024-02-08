/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      "./templates/**/*.{html,js,templ,go}",
      "./templates/common/**/*.{html,js,templ,go}",
      "./templates/components/**/*.{html,js,templ,go}",
    ],
    theme: {
      extend: {},
      fontFamily: {},
    },
    plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography"), require("daisyui")],
  };