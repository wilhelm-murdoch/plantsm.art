<script lang="ts">
	import type { FilterItem } from './filters';

	export let filter: FilterItem;
</script>

<span class="border border-green-800/50 bg-green-100 m-1 inline-flex items-center rounded-md shadow-sm text-green-800 py-1 pl-3 pr-2 text-sm font-medium">
	<svg id="emoji" class="w-6 h-6 mr-1" viewBox="0 0 72 72" xmlns="http://www.w3.org/2000/svg">
		<g id="color">
			<path fill="#B1CC33" stroke="none" d="M36.8349,18.0781c1.7393,1.6583,1.6481,4.5775,1.6481,4.5775s-2.9114,0.2303-4.651-1.428 c-1.7396-1.6582-1.6481-4.5775-1.6481-4.5775S35.0958,16.4213,36.8349,18.0781z"/>
			<path fill="#5C9E31" stroke="none" d="M42.0884,54.3627c-4.6114,1.465-9.3345-1.6943-9.3345-1.6943s2.7546-5.5346,7.3663-7.0003 s9.3345,1.6943,9.3345,1.6943S46.6971,52.8966,42.0884,54.3627z"/>
			<path fill="#B1CC33" stroke="none" d="M49.0484,35.6919c-3.6542,1.0489-7.2941-1.5366-7.2941-1.5366s2.2856-4.2862,5.9401-5.3357 s7.2941,1.5366,7.2941,1.5366S52.7005,34.642,49.0484,35.6919z"/>
			<path fill="#5C9E31" stroke="none" d="M27.0179,32.1509c3.0883,2.3003,3.5096,6.9616,3.5096,6.9616s-4.5864,0.9311-7.6753-1.3691 c-3.0889-2.3002-3.5096-6.9616-3.5096-6.9616S23.9302,29.8529,27.0179,32.1509z"/>
		</g>
		<g id="hair"/>
		<g id="skin"/>
		<g id="skin-shadow"/>
		<g id="line">
			<path fill="none" stroke="#000000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M36.8349,18.0781 c1.7393,1.6583,1.6481,4.5775,1.6481,4.5775s-2.9114,0.2303-4.651-1.428c-1.7396-1.6582-1.6481-4.5775-1.6481-4.5775 S35.0958,16.4213,36.8349,18.0781z"/>
			<path fill="none" stroke="#000000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27.0179,32.1509 c3.0883,2.3003,3.5096,6.9616,3.5096,6.9616s-4.5864,0.9311-7.6753-1.3691c-3.0889-2.3002-3.5096-6.9616-3.5096-6.9616 S23.9302,29.8529,27.0179,32.1509z"/>
			<path fill="none" stroke="#000000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M42.0884,54.3627 c-4.6114,1.465-9.3345-1.6943-9.3345-1.6943s2.7546-5.5346,7.3663-7.0003s9.3345,1.6943,9.3345,1.6943 S46.6971,52.8966,42.0884,54.3627z"/>
			<path fill="none" stroke="#000000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M49.0484,35.6919 c-3.6542,1.0489-7.2941-1.5366-7.2941-1.5366s2.2856-4.2862,5.9401-5.3357s7.2941,1.5366,7.2941,1.5366 S52.7005,34.642,49.0484,35.6919z"/>
			<path fill="none" stroke="#000000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M39.9046,25.7818 c0,4.7785-2.8879,10.0608-4.8815,13.1378c-0.8136,1.2558-1.7383,2.4273-2.7031,3.5711c-2.5753,3.0534-5.2237,8.2176-3.876,15.8715"/>
		</g>
	</svg>
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
