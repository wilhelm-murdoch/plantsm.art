<script lang="ts">
	import { lazy } from '$lib/utils/lazy';
	import { getByAnimal, normalizeAnimal } from '$utils/animals';
	import type { PlantSlim } from '$lib/types/plant';
	import { getImageUrl } from '$lib/utils/urls';
	import { currentlySelectedPlant } from '$utils/stores/card';
	import { goto } from '$app/navigation';

	export let plant: PlantSlim;

	const selectPlant = (pid: string) => {
		currentlySelectedPlant.set(pid);
		goto(`/plant/${pid}`);
	}
</script>

<div on:keyup={() => selectPlant(plant.pid)} on:click={() => selectPlant(plant.pid)} role="button" tabindex="0" class="flex flex-col overflow-hidden rounded-lg shadow-xl fade-in-animation ring-2 ring-emerald-700 border-2 border-emerald-50 transition hover:scale-102 ease-in-out duration-250">
	<div class="flex-shrink-0 relative">
		<div class="overflow-hidden background-fallback relative bg-green-50">
			<span class="absolute top-1 right-1 inline-flex items-center rounded-md bg-black px-2.5 py-1 text-xs font-medium text-white opacity-80 z-[1]">
				<svg class="w-6 h-6 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
					<path stroke-linecap="round" stroke-linejoin="round" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
				</svg>+{plant.image_total - 1}
			</span>
			<a href="/plant/{plant.pid}" title="Read more about {plant.name}.">
				<img class="is-lazy h-52 w-full object-cover" use:lazy={getImageUrl(plant.cover_image_url, 'medium')} alt="Cover image for {plant.name}." />
			</a>
		</div>

		<h3 class="unstyled bg-gradient-to-tr from-green-700 to-emerald-900 p-2.5 text-xl font-extralight text-white">
			{plant.name}
		</h3>
	</div>
	<div class="flex flex-1 flex-col justify-between bg-white">
		<div class="flex-1">
			<h3 class="unstyled font-sans mb-3 p-2.5 border-b text-sm bg-slate-50">Affects</h3>
			<p class="px-2.5 pb-2.5 text-sm font-medium space-x-1 space-y-1">
				{#each plant.animals as animal}
					<span class="inline-flex bg-{getByAnimal(animal).background} p-1 rounded-md cursor-pointer border-solid border border-{getByAnimal(animal).foreground}/50" title="{normalizeAnimal(animal)}">
						<svg id="emoji" class="h-8 w-8" viewBox="0 0 72 72" xmlns="http://www.w3.org/2000/svg">
							{@html getByAnimal(animal).svg}
						</svg>
					</span>
				{/each}
			</p>
			<h3 class="unstyled font-sans mb-3 p-2.5 border-b text-sm bg-slate-50">Common Names</h3>
			<p class="px-2.5 pb-2.5 text-sm text-slate-500">
				{#if plant.common != null}
					{#each plant.common as common, i}
						{#if plant.common.length == 2}
							{common}{#if i == 0}&nbsp;&amp;&nbsp;{/if}
						{:else}
							{common}{#if plant.common.length != 1 && i != 2},&nbsp;{/if}
						{/if}
					{/each}
					{#if plant.common_total > 3}
						<span class="inline-flex items-center rounded-sm bg-gray-100 px-2.5 py-0.5 text-xs font-medium text-gray-800">+{plant.common_total - 3} more </span>
					{/if}
				{:else}
					<span class="text-slate-400 text-center"><i>None listed. Maybe <a href="/" title="Missing details? Submit your own changes for approval." class="border-b hover:text-slate-500">add some</a>?</i></span>
				{/if}
			</p>
			<h3 class="unstyled font-sans mb-3 p-2.5 border-b text-sm bg-slate-50">Symptoms</h3>
			<p class="px-2.5 pb-2.5 text-sm text-slate-500">
				{#each plant.symptoms as symptom, i}
					{#if plant.symptoms.length == 2}
						{symptom}{#if i == 0}&nbsp;&amp;&nbsp;{/if}
					{:else}
						{symptom}{#if plant.symptoms.length != 1 && i != 2},&nbsp;{/if}
					{/if}
				{/each}
				{#if plant.symptoms_total > 3}
					<span class="inline-flex items-center rounded-sm bg-gray-100 px-2.5 py-0.5 text-xs font-medium text-gray-800">+{plant.symptoms_total - 3} more </span>
				{/if}
			</p>
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
