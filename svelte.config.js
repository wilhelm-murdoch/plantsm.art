import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';
import { vitePreprocess } from '@sveltejs/kit/vite';

import fs from 'fs';
import path from 'path'

const entries = fs.readdirSync('./src/lib/data/plant').filter(file => path.extname(file) === '.json').map(file => "/plant/" + path.parse(file).name);

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: [
		preprocess({
			postcss: true
		})
	],
	kit: {
		adapter: adapter({
			pages: 'build',
			assets: 'build',
			fallback: null,
			precompress: false,
			strict: true
		}),
		alias: {
			'$components': path.resolve('./src/lib/components'),
			'$lib': path.resolve('./src/lib'),
			'$utils': path.resolve('./src/lib/utils'),
			'$static': path.resolve('./static')
		},
		prerender: {
			entries: [
				"/",
				...entries
			],
		},
	},

	preprocess: vitePreprocess()
};

export default config;
