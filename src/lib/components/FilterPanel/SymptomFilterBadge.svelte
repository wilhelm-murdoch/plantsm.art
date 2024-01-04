<script lang="ts">
	import type { FilterItem } from './filters';

	export let filter: FilterItem;
</script>

<span
	class:bg-green-700={filter.term != 'Death'}
	class:bg-black={filter.term == 'Death'}
	class="m-1 border border-green-700 inline-flex items-center rounded-md shadow-sm text-white py-1 pl-3 pr-2 text-sm font-medium"
>
	<svg
		class:hidden={filter.term != 'Death'}
		xmlns="http://www.w3.org/2000/svg"
		viewBox="0 0 512 512"
		class="inline-block outline-none fill-current w-4 h-4 mr-2"
		><path
			d="M416 400V464C416 490.5 394.5 512 368 512H320V464C320 455.2 312.8 448 304 448C295.2 448 288 455.2 288 464V512H224V464C224 455.2 216.8 448 208 448C199.2 448 192 455.2 192 464V512H144C117.5 512 96 490.5 96 464V400C96 399.6 96 399.3 96.01 398.9C37.48 357.8 0 294.7 0 224C0 100.3 114.6 0 256 0C397.4 0 512 100.3 512 224C512 294.7 474.5 357.8 415.1 398.9C415.1 399.3 416 399.6 416 400V400zM160 192C124.7 192 96 220.7 96 256C96 291.3 124.7 320 160 320C195.3 320 224 291.3 224 256C224 220.7 195.3 192 160 192zM352 320C387.3 320 416 291.3 416 256C416 220.7 387.3 192 352 192C316.7 192 288 220.7 288 256C288 291.3 316.7 320 352 320z"
		/></svg
	>
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
