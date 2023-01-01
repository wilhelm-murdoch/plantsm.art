<script lang="ts">
	import { onMount } from 'svelte';
	import type { SymptomItem } from './filters';
	import { filters } from './filters';

	export let isOpen: boolean = false;
</script>

<div class="relative inline-block px-4 text-left">
	<button
		type="button"
		class="group inline-flex justify-center text-sm font-medium text-gray-700 hover:text-gray-900"
		aria-expanded="false"
		on:click={() => {
			isOpen = !isOpen;
		}}
	>
		<span>Symptoms</span>
		<span
			class="ml-1.5 rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700"
			>{$filters.filter((f) => {
				return f.type == 'symptoms';
			}).length}</span
		>
		<!-- Heroicon name: mini/chevron-down -->
		<svg
			class="-mr-1 ml-1 h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500"
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 20 20"
			fill="currentColor"
			aria-hidden="true"
		>
			<path
				fill-rule="evenodd"
				d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z"
				clip-rule="evenodd"
			/>
		</svg>
	</button>

	<div
		class:hidden={!isOpen}
		class="absolute right-0 z-10 mt-2 origin-top-right rounded-md bg-white shadow-2xl ring-1 ring-black ring-opacity-5 focus:outline-none max-h-96 w-64 overflow-scroll"
	>
		<form>
			<div class="relative flex items-center border-b text-center w-64 p-4 mb-4">
				<div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
					<svg
						class="ml-6 h-4 w-4 text-gray-400"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
						/>
					</svg>
				</div>
				<input
					bind:value={symptomFilterSearch}
					type="text"
					id="filter-symptom-search"
					class="w-52 rounded-md text-gray-400 border-gray-300 pl-9 shadow-sm focus:border-green-400 focus:ring-green-400 sm:text-sm lg:text-base mx-auto"
				/>
			</div>
			{#if filteredSymptoms.length}
				{#each filteredSymptoms as symptom, i}
					<div class="flex items-center px-4 mb-4">
						<input
							id="filter-symptom-{i}"
							name="symptom[]"
							value="white"
							type="checkbox"
							class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
							on:click={() => handleCheckboxFilter('symptoms', symptom.name)}
							checked={$filters.some((f) => {
								return f.type == 'symptoms' && f.term == symptom.name;
							})}
						/>
						<label
							for="filter-symptom-{i}"
							class="ml-3 whitespace-nowrap text-sm font-medium text-gray-900 cursor-pointer"
							>{symptom.name}
							<span
								class="rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700"
								>{symptom.plants}</span
							></label
						>
					</div>
				{/each}
			{:else}
				<div class="items-center mb-4 text-center text-gray-400">no matches</div>
			{/if}
		</form>
	</div>
</div>
