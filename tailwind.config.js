/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./templates/*.html",
    "./public/**/*.js"
  ],
  theme: {
    extend: {
      colors: {
        primary: '#2563eb',
        secondary: '#0ea5e9',
        accent: '#16a34a',
        neutral: '#1e293b',
        "base-100": "#ffffff",
        "base-200": "#f3f4f6",
      },
    },
  },
  plugins: [require("daisyui")],
}
