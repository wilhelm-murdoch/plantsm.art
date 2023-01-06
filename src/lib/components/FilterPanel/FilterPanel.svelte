<script lang="ts">
	import { browser } from '$app/environment';
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';
	import { getAllAnimals, getByAnimal } from '$utils/animals';
	import { filters } from './filters';
	import type { FilterItem, SymptomItem } from './filters';
	import AffectFilterBadge from './AffectFilterBadge.svelte';
	import SymptomFilterBadge from './SymptomFilterBadge.svelte';
	import TextFilterBadge from './TextFilterBadge.svelte';

	export let resultCount: number = 0;

	let affectsOpen = false;
	let symptomsOpen = false;
	let apiSymptoms: SymptomItem[] = [];

	onMount(async () => {
		fetch('/symptoms.json')
			.then((response) => response.json())
			.then((data) => {
				apiSymptoms = data.sort(function (a: SymptomItem, b: SymptomItem) {
					if (a.slug < b.slug) return -1;
					if (a.slug > a.slug) return 1;
					return 0;
				});
			})
			.catch((error) => {
				return [];
			});
	});

	let symptomFilterSearch = '';

	$: filteredSymptoms = apiSymptoms.filter((s) =>
		s.name.toLowerCase().includes(symptomFilterSearch.toLowerCase())
	);

	export function addFilter(type: string, term: string) {
		if (!['affects', 'symptoms', 'text'].includes(type)) {
			return;
		}

		let filter: FilterItem = {
			type: type,
			term: term
		};

		if (filterExists(filter)) {
			return;
		}

		$filters = [...$filters, filter];
	}

	export function filterExists(filter: FilterItem): boolean {
		return $filters.some((f) => {
			return f.type == filter.type && f.term == filter.term;
		});
	}

	function handleCheckboxFilter(type: string, term: string) {
		let filter: FilterItem = {
			type: type,
			term: term
		};

		return filterExists(filter) ? removeFilter(filter) : addFilter(type, term);
	}

	function removeFilter(filter: FilterItem) {
		let index = $filters.findIndex((f) => {
			return f.type == filter.type && f.term == filter.term;
		});

		if (index != -1) {
			$filters.splice(index, 1);
		}

		$filters = $filters;
	}

	function handleSearchInput(e: KeyboardEvent): void {
		if (e.key == 'Enter') {
			let searchInputElement = <HTMLInputElement>document.getElementById('filter-search');
			if (searchInputElement.value.trim() != '') {
				addFilter('text', searchInputElement.value.trim());
				searchInputElement.value = '';
			}
		}
	}

	function onWindowKeydown(e: KeyboardEvent): void {
		if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
			e.preventDefault();
			document.getElementById('filter-search')?.focus();
		}
	}

	function focus(n: HTMLInputElement) {
		n.focus();
	}
</script>

<section aria-labelledby="filter-heading">
	<div class="border-b border-gray-200 bg-white py-2">
		<div class="mx-auto flex max-w-7xl items-center justify-between">
			<div class="relative inline-block text-left w-1/2">
				<div class="relative mt-1 flex items-center">
					<div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
						<svg
							class="-ml-1 mr-3 h-4 w-4 text-gray-400"
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
						type="text"
						name="filter-search"
						id="filter-search"
						class="block h-12 w-full rounded-md text-gray-400 border-gray-300 pl-8 pr-14 shadow-sm focus:border-green-400 focus:ring-green-400 sm:text-sm lg:text-base"
						use:focus
						on:keydown={handleSearchInput}
						on:click={() => {
							affectsOpen = false;
							symptomsOpen = false;
						}}
					/>

					<div class="absolute inset-y-0 right-0 flex py-1.5 pr-1.5">
						{#if browser}
							<kbd
								class="shortcut inline-flex items-center rounded border border-gray-200 px-2 font-sans text-sm font-medium text-gray-400"
								>{navigator.userAgent.search('Mac') !== -1 ? '⌘' : 'Ctrl'}+K</kbd
							>
						{/if}
					</div>
				</div>
			</div>

			<div class="sm:block">
				<div class="flow-root">
					<div class="-mx-4 flex items-center divide-x divide-gray-200">
						<div class="relative inline-block px-4 text-left">
							<button
								type="button"
								class="group inline-flex justify-center text-sm font-medium text-gray-700 hover:text-gray-900"
								aria-expanded="false"
								on:click={() => {
									affectsOpen = !affectsOpen;
									symptomsOpen = false;
								}}
							>
								<span>Affects</span>

								<span
									class="ml-1.5 rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700"
									>{$filters.filter((f) => {
										return f.type == 'affects';
									}).length}</span
								>

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
								class:hidden={!affectsOpen}
								class="absolute right-0 z-10 mt-2 origin-top-right rounded-md bg-white p-4 shadow-2xl ring-1 ring-black ring-opacity-5 focus:outline-none"
							>
								<form class="space-y-4">
									{#each getAllAnimals() as animal}
										<div class="flex items-center">
											<input
												id="filter-affects-{animal}"
												name="affects[]"
												value={animal}
												type="checkbox"
												class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
												on:click={() => handleCheckboxFilter('affects', animal)}
												checked={$filters.some((f) => {
													return f.type == 'affects' && f.term == animal;
												})}
											/>
											<label
												for="filter-affects-{animal}"
												class="ml-3 whitespace-nowrap pr-6 text-sm font-medium text-gray-900 cursor-pointer"
											>
												<span
													class="inline-flex items-center rounded-md bg-{getByAnimal(animal)
														.background} px-2.5 py-1 text-{getByAnimal(animal).foreground}"
													>{getByAnimal(animal).emoji} {animal}</span
												>
											</label>
										</div>
									{/each}
								</form>
							</div>
						</div>

						<div class="relative inline-block px-4 text-left">
							<button
								type="button"
								class="group inline-flex justify-center text-sm font-medium text-gray-700 hover:text-gray-900"
								aria-expanded="false"
								on:click={() => {
									symptomsOpen = !symptomsOpen;
									affectsOpen = false;
								}}
							>
								<span>Symptoms</span>
								<span
									class="ml-1.5 rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700"
									>{$filters.filter((f) => {
										return f.type == 'symptoms';
									}).length}</span
								>
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
								class:hidden={!symptomsOpen}
								class="absolute right-0 z-10 mt-2 origin-top-right rounded-md bg-white shadow-2xl ring-1 ring-black ring-opacity-5 focus:outline-none max-h-96 w-64 overflow-y-scroll"
							>
								<form>
									<div class="relative flex items-center border-b text-center p-4 mb-4">
										<div
											class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3"
										>
											<svg
												class="ml-4 h-4 w-4 text-gray-400"
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
														>{symptom.count}</span
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

						<div class="relative inline-block px-4 text-left">
							<button
								type="button"
								class="group inline-flex justify-center text-sm border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
								aria-expanded="false"
								on:click={() => {
									$filters = [];
									symptomsOpen = false;
									affectsOpen = false;
								}}
							>
								<span>Clear</span>
							</button>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>

	<div class="bg-gray-50">
		<div class="mx-auto max-w-7xl py-3 sm:flex sm:items-center">
			<h3 class="text-sm font-medium text-gray-500">Filters</h3>

			<div aria-hidden="true" class="hidden h-5 w-px bg-gray-300 sm:ml-4 sm:block" />

			<div class="mt-2 sm:mt-0 sm:ml-4">
				<div class="-m-1 flex flex-wrap items-center">
					{#if $filters.length}
						{#each $filters as filter, i}
							{#if filter.type == 'affects'}
								<AffectFilterBadge {filter} on:click={() => removeFilter(filter)} />
							{:else if filter.type == 'symptoms'}
								<SymptomFilterBadge {filter} on:click={() => removeFilter(filter)} />
							{:else}
								<TextFilterBadge {filter} on:click={() => removeFilter(filter)} />
							{/if}
						{/each}
						{#if $filters.length}
							<span
								class="ml-1.5 rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700"
							>
								{#if resultCount == 0}
									no matches
								{:else if resultCount == 1}
									1 match
								{:else}
									{resultCount} matches
								{/if}
							</span>
						{/if}
					{:else}
						<span class="text-slate-400 text-sm italic m-1 py-1.5">none specified &hellip;</span>
					{/if}
				</div>
			</div>
		</div>
	</div>
</section>

<svelte:window on:keydown={onWindowKeydown} />

<style>
	@keyframes fade-in {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}

	.shortcut {
		animation: fade-in 0.2s;
	}
</style>
