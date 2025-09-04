/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,js}"],
  theme: {
    fontFamily: {
        sans: ["Space Grotesk"],
    },
    extend: {
        colors: {
            bfg: {
                standard: 'rgb(35,35,35)'
            },
        },
    },
  },
  plugins: [],
}

