<script lang="ts">
	import type { CardPlant } from './Card';
	import { affectedAnimals } from '$utils/animals';
	import { isDeadly } from './Card';
	import { filters } from '$components/FilterPanel/filters';

	export let plant: CardPlant;
</script>

<div class="flex flex-col overflow-hidden rounded-lg shadow-lg fade-in-animation">
	<div class="flex-shrink-0 relative">
		<span
			class="absolute top-1 right-1 inline-flex items-center rounded-md bg-black px-2.5 py-1 text-xs font-medium text-white opacity-75"
			><svg
				class="w-4 h-4 mr-1"
				fill="none"
				stroke="currentColor"
				viewBox="0 0 24 24"
				xmlns="http://www.w3.org/2000/svg"
				><path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
				/></svg
			>+{plant.images.length - 1}</span
		>
		<a href="/plant/{plant.pid}" title="Read more about {plant.name}.">
			<img
				class="h-48 w-full object-cover"
				src="images/{plant.images[0].relative_path.replace('original', 'medium')}"
				alt="Cover image for {plant.name}."
			/>
			{#if isDeadly(plant)}
				<h3 class="unstyled bg-black p-2.5 text-xl font-extralight text-white">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						x="0px"
						y="0px"
						viewBox="0 0 512 512"
						class="svg-icon inline-block outline-none fill-current w-4 h-4 mr-2 -translate-y-0.5"
						focusable="false"
						data-testid="svg-icon"
						><path
							d="M416 400V464C416 490.5 394.5 512 368 512H320V464C320 455.2 312.8 448 304 448C295.2 448 288 455.2 288 464V512H224V464C224 455.2 216.8 448 208 448C199.2 448 192 455.2 192 464V512H144C117.5 512 96 490.5 96 464V400C96 399.6 96 399.3 96.01 398.9C37.48 357.8 0 294.7 0 224C0 100.3 114.6 0 256 0C397.4 0 512 100.3 512 224C512 294.7 474.5 357.8 415.1 398.9C415.1 399.3 416 399.6 416 400V400zM160 192C124.7 192 96 220.7 96 256C96 291.3 124.7 320 160 320C195.3 320 224 291.3 224 256C224 220.7 195.3 192 160 192zM352 320C387.3 320 416 291.3 416 256C416 220.7 387.3 192 352 192C316.7 192 288 220.7 288 256C288 291.3 316.7 320 352 320z"
						/></svg
					>
					{plant.name}
				</h3>
			{:else}
				<h3
					class="unstyled bg-gradient-to-tr from-green-700 to-emerald-900 p-2.5 text-xl font-extralight text-white"
				>
					{plant.name}
				</h3>
			{/if}
		</a>
	</div>
	<div class="flex flex-1 flex-col justify-between bg-white">
		<div class="flex-1">
			<h3 class="unstyled font-sans mb-3 p-2.5 border-b text-sm bg-slate-50">Affects</h3>
			<p class="px-2.5 pb-2.5 text-sm font-medium space-x-1 space-y-1">
				{#each plant.animals.slice(0, 4) as animal}
					<span
						class="inline-flex items-center rounded-md bg-{affectedAnimals[animal]
							.background} px-2.5 py-1 text-{affectedAnimals[animal].foreground}"
						>{affectedAnimals[animal].emoji} {animal}</span
					>{/each}{#if plant.animals.length > 4}<span
						class="inline-flex items-center rounded-md bg-gray-100 py-1 px-2.5 text-gray-800"
						>+{plant.animals.length - 4} more
					</span>
				{/if}
			</p>
			<h3 class="unstyled font-sans mb-3 p-2.5 border-b text-sm bg-slate-50">Other Names</h3>
			<p class="px-2.5 pb-2.5 text-sm text-slate-500">
				{#if plant.common != null}
					{#each plant.common.slice(0, 3) as common, i}
						{common.name}{#if i != plant.common.slice(0, 3).length - 1},&nbsp;{/if}
					{/each}
					{#if plant.common.length > 3}
						<span
							class="inline-flex items-center rounded-sm bg-gray-100 px-2.5 py-0.5 text-xs font-medium text-gray-800"
							>+{plant.common.length - 3} more
						</span>
					{/if}
				{:else}
					<span class="text-slate-400 text-center"
						><i
							>None listed. Maybe <a
								href="/"
								title="Missing details? Submit your own changes for approval."
								class="border-b hover:text-slate-500">add some</a
							>?</i
						></span
					>
				{/if}
			</p>
			<h3 class="unstyled font-sans mb-3 p-2.5 border-b text-sm bg-slate-50">Symptoms</h3>
			<p class="px-2.5 pb-2.5 text-sm text-slate-500">
				{#each plant.symptoms.slice(0, 3) as symptom, i}
					<a
						href="/?f=s:{symptom.name}"
						title="View other plants with this symptom."
						class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
						>{symptom.name}</a
					>{#if i != plant.symptoms.slice(0, 3).length - 1},&nbsp;{/if}
				{/each}
				{#if plant.symptoms.length > 3}
					<span
						class="inline-flex items-center rounded-sm bg-gray-100 px-2.5 py-0.5 text-xs font-medium text-gray-800"
						>+{plant.symptoms.length - 3} more
					</span>
				{/if}
			</p>
		</div>

		<div class="-mt-px flex border-t">
			<a
				href="/plant/{plant.pid}"
				title="Read more about {plant.name}."
				class="relative inline-flex w-0 flex-1 items-center justify-center rounded-br-lg border border-transparent py-1 text-sm font-medium text-gray-700 hover:bg-slate-50 hover:text-gray-500"
			>
				<svg
					class="h-6 w-6 text-slate-400"
					fill="none"
					stroke="currentColor"
					viewBox="0 0 24 24"
					xmlns="http://www.w3.org/2000/svg"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M5 12h.01M12 12h.01M19 12h.01M6 12a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0z"
					/></svg
				>
			</a>
		</div>
	</div>
</div>

<style>
	.fade-in-animation {
		-webkit-animation: fade-in-animation 1s cubic-bezier(0.39, 0.575, 0.565, 1) both;
		animation: fade-in-animation 1s cubic-bezier(0.39, 0.575, 0.565, 1) both;
	}

	@-webkit-keyframes fade-in-animation {
		0% {
			opacity: 0;
		}
		100% {
			opacity: 1;
		}
	}
	@keyframes fade-in-animation {
		0% {
			opacity: 0;
		}
		100% {
			opacity: 1;
		}
	}
</style>
