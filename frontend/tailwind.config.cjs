/** @type {import('tailwindcss').Config} */

const config = {
  theme: {
    extend: {
      fontFamily: ['Inter var']
    },
  },
  purge: ["./index.html",'./src/**/*.{svelte,js,ts}'], //for unused css
  plugins: [
    require("daisyui")
  ],
  daisyui: {
    themes: [
      {
        ifcctheme: {
          "color-scheme": "dark",
          "primary": "oklch(65.69% 0.196 275.75)",
          "secondary": "oklch(74.8% 0.26 342.55)",
          "accent": "oklch(74.51% 0.167 183.61)",
          "neutral": "#2a323c",
          "neutral-content": "#A6ADBB",
          "base-100": "#1d232a",
          "base-200": "#191e24",
          "base-300": "#15191e",
          "base-content": "#A6ADBB",
          "primary-content": "#000000"
        },
      },
    ], 
  },
  variants: {
    extend: {},
  },
  darkmode: false, // or 'media' or 'class'
}

module.exports = config;