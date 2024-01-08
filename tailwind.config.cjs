module.exports = {
	content: [
		'./src/**/*.{html,js,svelte,ts}'
	],
	theme: {
		extend: {
			scale: {
				'102': '1.02',
			  }
		}
	},
	plugins: [
		require('@tailwindcss/forms'),
		require('@tailwindcss/typography')
	],
	safelist: [
		{
		  pattern: /border-(yellow|orange|green|blue|slate|green|indigo)-(50|100|700|800)/,
		},
	  ],
};