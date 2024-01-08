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
			{#if plant.severity.level == 1}
				<span class="absolute top-1 left-1 inline-flex items-center rounded-md bg-green-900 px-2.5 py-1 text-xs font-medium text-white opacity-80 z-[1]">
					<svg class="w-6 h-6 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
						<path stroke-linecap="round" stroke-linejoin="round" d="m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z" />
					</svg>
					{plant.severity.label} Symptoms
				</span>
			{:else if plant.severity.level == 2}
				<span class="absolute top-1 left-1 inline-flex items-center rounded-md bg-yellow-900 px-2.5 py-1 text-xs font-medium text-white opacity-80 z-[1]">
					<svg class="w-6 h-6 mr-1" stroke="currentColor" viewBox="0 0 72 72" xmlns="http://www.w3.org/2000/svg">
						<g id="line">
						  <path fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M32.5222,13.005c0.6977-1.2046,1.9862-2.0244,3.4778-2.0244c1.4916,0,2.7801,0.8198,3.4778,2.0244l20.9678,41.9351 C60.7889,55.5339,61,56.2136,61,56.9483c0,2.2272-1.8051,4.0323-4.0323,4.0323l-41.9354,0.0173 C12.8051,60.9979,11,59.192,11,56.9657c0-0.7356,0.211-1.4145,0.5544-2.0083L32.5222,13.005"/>
						  <path fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M37.6129,47.2709c0,0.8907-0.7222,1.6129-1.6129,1.6129c-0.8907,0-1.6129-0.7222-1.6129-1.6129V23.8925 c0-0.8907,0.7222-1.6129,1.6129-1.6129c0.8907,0,1.6129,0.7222,1.6129,1.6129V47.2709z"/>
						  <circle cx="36" cy="54.529" r="1.6129" fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/>
						</g>
					</svg>
					{plant.severity.label} Symptoms
				</span>
			{:else if plant.severity.level == 3}
				<span class="absolute top-1 left-1 inline-flex items-center rounded-md bg-orange-900 px-2.5 py-1 text-xs font-medium text-white opacity-80 z-[1]">
					<svg class="w-6 h-6 mr-1" stroke="currentColor" viewBox="0 0 72 72" xmlns="http://www.w3.org/2000/svg">
						<path fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M36,11c-13.8066,0-25,11.1926-25,25c0,13.8066,11.1934,25,25,25c13.8065,0,25-11.1934,25-25C61,22.1926,49.8065,11,36,11z"/>
						<path fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M55.5195,39.4297c0.1986-1.1263,0.3081-2.2831,0.3081-3.4659c0-1.171-0.1078-2.316-0.3014-3.4306H16.4738 c-0.1936,1.1146-0.3014,2.2596-0.3014,3.4306c0,1.1828,0.1095,2.3396,0.3081,3.4659H55.5195z"/>
					</svg>
					{plant.severity.label} Symptoms
				</span>
			{:else if plant.severity.level == 4}
				<span class="absolute top-1 left-1 inline-flex items-center rounded-md bg-red-900 px-2.5 py-2 text-xs font-medium text-white opacity-80 z-[1]">
					<svg class="w-4 h-4 mr-1" fill="currentColor" viewBox="0 0 512 512" xmlns="http://www.w3.org/2000/svg">
						<path d="M416 400V464C416 490.5 394.5 512 368 512H320V464C320 455.2 312.8 448 304 448C295.2 448 288 455.2 288 464V512H224V464C224 455.2 216.8 448 208 448C199.2 448 192 455.2 192 464V512H144C117.5 512 96 490.5 96 464V400C96 399.6 96 399.3 96.01 398.9C37.48 357.8 0 294.7 0 224C0 100.3 114.6 0 256 0C397.4 0 512 100.3 512 224C512 294.7 474.5 357.8 415.1 398.9C415.1 399.3 416 399.6 416 400V400zM160 192C124.7 192 96 220.7 96 256C96 291.3 124.7 320 160 320C195.3 320 224 291.3 224 256C224 220.7 195.3 192 160 192zM352 320C387.3 320 416 291.3 416 256C416 220.7 387.3 192 352 192C316.7 192 288 220.7 288 256C288 291.3 316.7 320 352 320z" />
					</svg>
					{plant.severity.label} Symptoms
				</span>
			{/if}
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
