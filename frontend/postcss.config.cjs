const tailwind = require('tailwindcss');
const tailwindConfig = require('./tailwind.config.cjs')
const autoprefixer = require('autoprefixer');

const config = {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
};

module.exports = config;
