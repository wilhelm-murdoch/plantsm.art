import { sveltekit } from '@sveltejs/kit/vite';
import { exec } from 'child_process'
import { promisify } from 'util'

const pexec = promisify(exec)

let version = (
	await Promise.allSettled([
		pexec('git rev-parse --short=8 HEAD'),
	])
).map(v => v.value?.stdout.trim())

/** @type {import('vite').UserConfig} */
const config = {
	define: {
		__BUILD_NUMBER__: version
	},
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
};

export default config;
