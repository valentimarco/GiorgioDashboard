import type { Config } from 'tailwindcss'

export default {
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
  important: true
} satisfies Config