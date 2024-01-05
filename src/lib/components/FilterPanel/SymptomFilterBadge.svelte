<script lang="ts">
	import type { FilterItem } from './filters';

	export let filter: FilterItem;
</script>

<span class="m-1 bg-green-700 border border-green-700 inline-flex items-center rounded-md shadow-sm text-white py-1.5 pl-3 pr-2 text-sm font-medium">
	<span>{filter.term}</span>
	<button
		on:click
		type="button"
		class="ml-1 inline-flex h-4 w-4 flex-shrink-0 p-1 text-white opacity-50 hover:opacity-100"
	>
		<svg class="h-2 w-2" stroke="currentColor" fill="none" viewBox="0 0 8 8">
			<path stroke-linecap="round" stroke-width="1.5" d="M1 1l6 6m0-6L1 7" />
		</svg>
	</button>
</span>
