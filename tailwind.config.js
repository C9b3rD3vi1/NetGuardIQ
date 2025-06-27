/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./templates/**/*.html",
    "./public/**/*.js",
  ],
  theme: {
    extend: {},
  },
  plugins: [require("rippleui")],
}
