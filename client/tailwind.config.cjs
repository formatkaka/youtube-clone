const plugin = require('tailwindcss/plugin');

/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
      colors: {
        gray: {
          900: 'var(--gray-900)',
          700: 'var(--gray-700)',
          500: 'var(--gray-500)',
          400: 'var(--gray-400)',
          300: 'var(--gray-300)',
          200: 'var(--gray-200)',
          100: 'var(--gray-100)',
        },
      },
      width: {
        left_nav: 'var(--left-nav)',
        feed: 'calc(100vw - var(--left-nav))',
      },
      height: {
        header: 'var(--header-height)',
        left_nav: 'calc(100vh - var(--header-height))',
      },
      margin: {
        left_nav: 'var(--left-nav)',
      },
    }
	},
	plugins: [
		require('@tailwindcss/line-clamp'),
		plugin(function ({ addUtilities }) {
			addUtilities({
				'.no-scrollbar': {
					'scrollbar-width': 'none',
					'&::-webkit-scrollbar': {
						display: 'none'
					}
				}
			});
		})
	]
};
