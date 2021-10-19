module.exports = {
  important: true,
  purge: ['./public/**/*.html', './src/**/*.{vue, js}'],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      backgroundImage: {
        'header': "url('/src/assets/header-pizza.jpg')",
      }
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
