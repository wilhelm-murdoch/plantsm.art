<script lang="ts">
	import { fade } from 'svelte/transition';
	import { type Item, items } from './items';

	const toggleOpen = (i: number) => {
		items[i].open = !items[i].open;
	};
</script>

<div class="relative bg-gray-50 mb-4 py-4 border-y">
	<div class="relative mx-auto max-w-7xl">
		<div
			class="overflow-hidden bg-white shadow-lg mb-4 pb-8 rounded-none sm:rounded-none lg:rounded-lg"
		>
			<h2
				class="unstyled bg-gradient-to-tr from-green-700 to-emerald-900 p-2.5 text-2xl font-extralight text-white text-center mb-8"
			>
				faq
			</h2>
			<div class="prose prose-lg prose-slate px-4 mx-auto text-gray-500">
				<h2
					class="block text-center text-3xl font-bold leading-8 tracking-tight text-gray-900 sm:text-4xl"
				>
					You've got questions? We've got answers!
				</h2>

				<p class="text-xl leading-8 text-gray-500 text-justify">
					We've compiled a list of common questions you may have about the site. Though, not unlike
					a healthy plant or pet, it requires constant care. If you notice something is missing,
					feel free to <a
						href="https://github.com/wilhelm-murdoch/plantsm.art/issues"
						title="Let us know on Github.">let us know</a
					> and we'll be sure to add it to the page.
				</p>
				<div class="prose">
					{#each items as item, i}
						<h3
							class="border-t pt-8 cursor-pointer"
							on:keydown={() => toggleOpen(i)}
							on:click={() => toggleOpen(i)}
						>
							{item.question}
						</h3>
						<p class:hidden={!item.open}>{@html item.answer}</p>
					{/each}
				</div>
			</div>
		</div>
	</div>
</div>
