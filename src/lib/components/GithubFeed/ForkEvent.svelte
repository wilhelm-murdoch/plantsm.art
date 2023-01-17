<script lang="ts">
	import Time from 'svelte-time';
	import { eventIcons } from './events';

	export let index: number = 0;
	export let numEntries: number = 0;
	export let event: any;
</script>

<li class="relative py-2">
	<div class="flex items-center mb-1 pb-5">
		<div
			class:hidden={index == numEntries - 1}
			class="absolute left-0 h-full w-0.5 bg-slate-200 self-start ml-4 -translate-x-1/2 translate-y-3"
			aria-hidden="true"
		/>
		<div
			class="absolute rounded-full p-1 text-{eventIcons.ForkEvent[0]}-800 bg-{eventIcons
				.ForkEvent[0]}-100"
			aria-hidden="true"
		>
			<svg
				class="w-6 h-6"
				fill="none"
				stroke-width="1"
				viewBox="0 0 24 24"
				stroke="currentColor"
				xmlns="http://www.w3.org/2000/svg">{@html eventIcons.ForkEvent[1]}</svg
			>
		</div>
		<h3 class="text-sm font-bold text-slate-800 pl-12">
			Made a <a
				href={event.payload.forkee.html_url}
				title="The forked version of our repository."
				class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500 font-normal"
				>fork</a
			>
			of our
			<a
				href="https://github.com/wilhelm-murdoch/plantsm.art"
				title="Our repository."
				class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500 font-normal"
				>repository</a
			>
			about
			<span class="font-normal text-sm text-slate-400"
				><Time relative timestamp={event.created_at} /></span
			>
		</h3>
	</div>
</li>
