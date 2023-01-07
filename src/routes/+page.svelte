<script lang="ts">
	import { onDestroy } from 'svelte';
	import { Card, FilterPanel } from '$components';
	import InfiniteScroll from 'svelte-infinite-scroll';
	import type { Plant, PlantsWrapped } from '$lib/types/plant';
	import { createSearchStore, searchHandler } from '$utils/stores/search';
	import { filters as filterStore } from '$components/FilterPanel/filters';

	export let data: PlantsWrapped;

	let page = 0;
	let size = 9;
	let plants: Plant[] = [];

	const searchStore = createSearchStore(data.plants);
	const unsubscribePlants = searchStore.subscribe((s) => searchHandler(s));
	const unsubscribeFilters = filterStore.subscribe((f) => {
		page = 0;
		plants = [];
		$searchStore.search = '^(?=.*\\b' + f.map((t) => t.term).join('\\b)(?=.*\\b') + '\\b).*$';
	});

	$: plants = [...plants, ...$searchStore.filtered.slice(size * page, size * (page + 1))];

	onDestroy(() => {
		unsubscribePlants();
		unsubscribeFilters();
	});
</script>

<FilterPanel resultCount={$searchStore.filtered.length} />

<div class="relative bg-gray-50 pb-5 mb-8 sm:px-6 sm:pt-4 lg:px-8 lg:pt-4 lg:pb-4 border-b">
	<div class="relative mx-auto max-w-7xl">
		{#if $searchStore.filtered.length == 0}
			<div class="my-auto flex-shrink-0 text-center py-16">
				<p class="text-base font-semibold text-green-600">Wow.</p>
				<h1 class="mt-2 text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
					No results!
				</h1>
				<p class="mt-2 text-base text-gray-500">
					Looks like you went a bit too hard on the filter. Try loosening your search a bit.
				</p>
				<div class="mt-6">
					<a
						href="/"
						on:click={() => {
							$filterStore = [];
						}}
						class="text-base font-medium text-green-600 hover:text-green-500"
					>
						Or, clear your filters ...
					</a>
				</div>
			</div>
		{:else}
			<div class="mx-auto grid max-w-lg gap-5 lg:max-w-none lg:grid-cols-3">
				{#each plants as plant}
					<Card {plant} />
				{/each}
				<InfiniteScroll window={true} threshold={150} on:loadMore={() => page++} />
			</div>
		{/if}
	</div>
</div>
