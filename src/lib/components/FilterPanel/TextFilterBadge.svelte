<script lang="ts">
	import type { FilterItem } from './filters';

	export let filter: FilterItem;
</script>

<span class="m-1 inline-flex items-center rounded-md border-solid border border-gray-900/50 bg-white py-1.5 pl-3 pr-2 text-sm font-medium text-gray-900">
	<span>{filter.term}</span>
	<button
		on:click
		type="button"
		class="ml-1 inline-flex h-4 w-4 flex-shrink-0 rounded-full p-1 text-gray-400 opacity-50 hover:opacity-100"
	>
		<svg class="h-2 w-2" stroke="currentColor" fill="none" viewBox="0 0 8 8">
			<path stroke-linecap="round" stroke-width="1.5" d="M1 1l6 6m0-6L1 7" />
		</svg>
	</button>
</span>
