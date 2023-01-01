<script lang="ts">
	import { affectedAnimals } from '$utils/animals';

	export let animals: string[] = [];
	export let classification: any = {};
	export let wikipedia_url = '';
</script>

<div class="flex flex-col overflow-hidden rounded-lg shadow-lg mb-4">
	<div class="flex-shrink-0">
		<h3
			class="unstyled font-sans mb-3 p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500"
		>
			Affects
		</h3>
		<p class="px-2.5 pb-3 text-md font-medium space-x-1 space-y-1">
			{#each animals as animal}
				<span
					class="inline-flex items-center rounded-md bg-{affectedAnimals[animal]
						.background} px-2.5 py-1 text-{affectedAnimals[animal].foreground}"
					>{affectedAnimals[animal].emoji} {animal}s</span
				>{/each}
		</p>
		<h3 class="unstyled font-sans p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500">
			Classification
		</h3>

		<div class="flex bg-gray-100 border-b">
			<div class="flex w-0 flex-1">
				<span class="p-4 text-slate-400">Kingdom</span>
			</div>
			<div class="-ml-px flex w-0 flex-1">
				<span class="p-4">
					<a
						href="https://en.wikipedia.org/wiki/{classification.kingdom}"
						title={classification.kingdom}
						class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
						>{classification.kingdom}</a
					></span
				>
			</div>
		</div>

		{#if classification.clades}
			{#each classification.clades as clade}
				<div class="flex border-b">
					<div class="flex w-0 flex-1">
						<span class="p-4 text-slate-400">— Clade</span>
					</div>
					<div class="-ml-px flex w-0 flex-1">
						<span class="p-4">
							<a
								href="https://en.wikipedia.org/wiki/{clade}"
								title={clade}
								class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
								>{clade}</a
							></span
						>
					</div>
				</div>
			{/each}
		{/if}

		<div class="flex bg-gray-100 border-b">
			<div class="flex w-0 flex-1">
				<span class="p-4 text-slate-400">Order</span>
			</div>
			<div class="-ml-px flex w-0 flex-1">
				<span class="p-4">
					<a
						href="https://en.wikipedia.org/wiki/{classification.order}"
						title={classification.order}
						class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
						>{classification.order}</a
					></span
				>
			</div>
		</div>

		<div class="flex border-b">
			<div class="flex w-0 flex-1">
				<span class="p-4 text-slate-400">Family</span>
			</div>
			<div class="-ml-px flex w-0 flex-1">
				<span class="p-4">
					<a
						href="https://en.wikipedia.org/wiki/{classification.family}"
						title={classification.family}
						class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
						>{classification.family}</a
					></span
				>
			</div>
		</div>

		<div class="flex bg-gray-100 border-b">
			<div class="flex w-0 flex-1">
				<span class="p-4 text-slate-400">Genus</span>
			</div>
			<div class="-ml-px flex w-0 flex-1">
				<span class="p-4">
					<a
						href="https://en.wikipedia.org/wiki/{classification.genus}"
						title={classification.genus}
						class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
						>{classification.genus}</a
					></span
				>
			</div>
		</div>
		{#if classification.species}
			<div class="flex border-b">
				<div class="flex w-0 flex-1">
					<span class="p-4 text-slate-400">Species</span>
				</div>
				<div class="-ml-px flex w-0 flex-1">
					<span class="p-4">
						<a
							href="https://en.wikipedia.org/wiki/{classification.species}"
							title={classification.species}
							class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
							>{classification.species}</a
						></span
					>
				</div>
			</div>
		{/if}
	</div>
	<div class="flex">
		<a
			href={wikipedia_url}
			title="Read more about this entry."
			class=" text-slate-400 inline-flex w-0 flex-1 items-center justify-center rounded-br-lg border border-transparent py-2 text-md hover:bg-slate-50 hover:text-gray-500"
		>
			wikipedia<svg
				class="w-5 h-5"
				fill="currentColor"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M11 3a1 1 0 100 2h2.586l-6.293 6.293a1 1 0 101.414 1.414L15 6.414V9a1 1 0 102 0V4a1 1 0 00-1-1h-5z"
				/><path
					d="M5 5a2 2 0 00-2 2v8a2 2 0 002 2h8a2 2 0 002-2v-3a1 1 0 10-2 0v3H5V7h3a1 1 0 000-2H5z"
				/></svg
			>
		</a>
	</div>
</div>
