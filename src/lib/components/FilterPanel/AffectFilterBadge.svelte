<script lang="ts">
	import { getByAnimal, normalizeAnimal } from '$lib/utils/animals';
	import type { FilterItem } from './filters';

	export let filter: FilterItem;
</script>

<span
	class="m-1 inline-flex items-center rounded-md border-solid border border-{getByAnimal(filter.term).foreground}/50 text-{getByAnimal(filter.term).foreground} bg-{getByAnimal(filter.term).background} py-1.5 pl-3 pr-2 text-sm font-medium"
>
	<svg id="emoji" class="h-6 w-6 mr-1" viewBox="0 0 72 72" xmlns="http://www.w3.org/2000/svg">
		{@html getByAnimal(filter.term).svg}
	</svg>
	<span>{normalizeAnimal(filter.term)}</span>

	<button
		on:click
		type="button"
		class="ml-1 inline-flex h-4 w-4 flex-shrink-0 rounded-full p-1 text-{getByAnimal(filter.term)
			.foreground} opacity-50 hover:opacity-100"
	>
		<svg class="h-2 w-2" stroke="currentColor" fill="none" viewBox="0 0 8 8">
			<path stroke-linecap="round" stroke-width="1.5" d="M1 1l6 6m0-6L1 7" />
		</svg>
	</button>
</span>
