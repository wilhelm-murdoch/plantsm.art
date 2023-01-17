<script lang="ts">
	import { getByAnimal } from '$utils/animals';

	export let plant: any = {};
</script>

<div class="flex flex-col overflow-hidden rounded-lg shadow-lg mb-4">
	<div class="flex-shrink-0">
		<h3 class="unstyled font-sans mb-3 p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500">Affects</h3>
		<p class="px-2.5 pb-3 text-md font-medium space-x-1 space-y-1">
			{#each plant.animals as animal}
				<span class="inline-flex items-center rounded-md bg-{getByAnimal(animal).background} px-2.5 py-1 text-{getByAnimal(animal).foreground}">{getByAnimal(animal).emoji} {animal}</span>
			{/each}
		</p>

		<h3 class="unstyled font-sans mb-3 p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500">Common Names</h3>
		<div class="prose text-md px-2.5 pb-2.5 text-gray-400">
			{#if plant.common}
				<ul>
					{#each plant.common as name, i}
						<li>{name.name}</li>
					{/each}
				</ul>
			{:else}
				<div class="rounded-md bg-slate-100 p-4">
					<p class="text-md">
						We currently don't have any other common names for this entry on record. Feel free to learn about how you can <a href="/contribute" title="Missing details? Submit your own changes for approval." class="border-b hover:text-slate-600">contribute</a> to make this page more complete.
					</p>
				</div>
			{/if}
		</div>
	</div>
	<h3 class="unstyled font-sans p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500">Symptoms</h3>
	<div class="prose text-md px-2.5 pb-2.5 text-gray-400">
		<ul>
			{#each plant.symptoms as symptom, i}
				<li>
					<span title="View other plants with this symptom.">{symptom.name}</span>
				</li>
			{/each}
		</ul>
	</div>
</div>
