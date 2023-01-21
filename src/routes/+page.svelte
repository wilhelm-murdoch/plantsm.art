<script lang="ts">
	import SvelteSeo from 'svelte-seo';
	import { onDestroy } from 'svelte';
	import { Card, FilterPanel } from '$components';
	import InfiniteScroll from 'svelte-infinite-scroll';
	import type { PlantSlim, PlantsWrapped } from '$lib/types/plant';
	import { createSearchStore, searchHandler } from '$utils/stores/search';
	import { filters as filterStore } from '$components/FilterPanel/filters';

	export let data: PlantsWrapped;

	let page = 0;
	let size = 12;
	let plants: PlantSlim[] = [];

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

	const twitterMeta = {
		site: '@wilhelm',
		title: 'Plant Smart &middot; Helping plants and pets peacefully coexist.',
		description: 'A free service that aims to provide a detailed listing of dangerous plants for your pets.',
		image: 'https://plantsm.art/images/og-cover.png',
		imageAlt: 'Plant Smart'
	};

	const openGraphMeta = {
		title: 'Plant Smart &middot; Helping plants and pets peacefully coexist.',
		description: 'A free service that aims to provide a detailed listing of dangerous plants for your pets.',
		type: 'article',
		url: 'https://plantsm.art',
		images: [
			{
				url: 'https://plantsm.art/images/og-cover.png',
				alt: 'Plant Smart'
			}
		]
	};
</script>

<FilterPanel resultCount={$searchStore.filtered.length} />

<div class="relative bg-gray-50 pb-5 mb-8 px-6 sm:pt-4 lg:pt-4 lg:pb-4 border-b">
	<div class="relative mx-auto max-w-7xl">
		{#if $searchStore.filtered.length == 0}
			<div class="mx-auto flex-shrink-0 text-center py-16 prose">
				<h2 class="text-base font-semibold text-green-600">Oh, no! 🙀</h2>
				<h3 class="text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">No matching results!</h3>
				<div class="text-left">
					<ol>
						<li>Try loosening your search or <a href="/" title="Clear all filters." on:click={() => {$filterStore = []}}>clearing out</a> your filters.</li>
						<li>The search mechanism doesn't account for <strike>speeling missed steaks</strike> spelling mistakes. Double-check your spelling.</li>
						<li>Should this plant be listed? <a href="https://github.com/wilhelm-murdoch/plantsm.art/issues/new?title=Update%20Request" title="Create a request on GitHub.">Send us the details</a> and we'll look into adding it to our database.</li>
					</ol>
				</div>
			</div>
		{:else}
			<div class="mx-auto grid max-w-lg gap-5 lg:max-w-none lg:grid-cols-3 md:grid-cols-2 md:max-w-none">
				{#each plants as plant (plant.pid)}
					<Card {plant} />
				{/each}
				<InfiniteScroll window={true} threshold={150} on:loadMore={() => page++} />
			</div>
		{/if}
	</div>
</div>

<SvelteSeo title="Plant Smart &middot; Helping plants and pets peacefully coexist." description="A free service that aims to provide a detailed listing of dangerous plants for your pets." canonical="https://plantsm.art" twitter={twitterMeta} openGraph={openGraphMeta} />
