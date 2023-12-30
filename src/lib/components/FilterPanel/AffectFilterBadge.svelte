<script lang="ts">
	import { fade } from 'svelte/transition';
	import { getByAnimal, normalizeAnimal } from '$lib/utils/animals';
	import type { FilterItem } from './filters';

	export let filter: FilterItem;
</script>

<span
	in:fade
	class="m-1 inline-flex items-center rounded-md shadow-sm text-{getByAnimal(filter.term)
		.foreground} bg-{getByAnimal(filter.term).background} py-1 pl-3 pr-2 text-sm font-medium"
>
	<!-- <span class="inline-flex items-center rounded-md bg-{getByAnimal(filter.term).background} px-2.5 py-1 text-{getByAnimal(filter.term).foreground} cursor-pointer"  title="{normalizeAnimal(filter.term)}"> -->
		<svg id="emoji" class="h-5 w-5 mr-1" viewBox="0 0 72 72" xmlns="http://www.w3.org/2000/svg">
			{@html getByAnimal(filter.term).svg}
		</svg>{normalizeAnimal(filter.term)}
	<!-- </span> -->


	<button
		on:click
		type="button"
		class="ml-1 inline-flex h-4 w-4 flex-shrink-0 rounded-full p-1 text-gray-400 hover:bg-gray-200 hover:text-gray-500"
	>
		<svg class="h-2 w-2" stroke="currentColor" fill="none" viewBox="0 0 8 8">
			<path stroke-linecap="round" stroke-width="1.5" d="M1 1l6 6m0-6L1 7" />
		</svg>
	</button>
</span>
