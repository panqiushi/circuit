module.exports = {
  // content: ['./templates/**/*.html', 'node_modules/preline/dist/*.js'],
  content: ['./frontend/**/*.html', './frontend/**/*.js'],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/typography'),
    require('@tailwindcss/forms'),
    require('preline/plugin'),
  ],
};
