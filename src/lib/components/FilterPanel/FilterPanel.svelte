<script lang="ts">
	import { onMount } from 'svelte';
	import { getAllAnimals, getByAnimal, normalizeAnimal } from '$utils/animals';
	import { filters } from './filters';
	import type { FilterItem, SymptomItem, FamilyItem } from './filters';
	import AffectFilterBadge from './AffectFilterBadge.svelte';
	import SymptomFilterBadge from './SymptomFilterBadge.svelte';
	import TextFilterBadge from './TextFilterBadge.svelte';
	import FamilyFilterBadge from './FamilyFilterBadge.svelte';
	import { slide } from 'svelte/transition';

	export let resultCount: number = 0;

	let affectsOpen = false;
	let symptomsOpen = false;
	let familyOpen = false;
	let severityOpen = false

	let apiSymptoms: SymptomItem[] = [];
	let apiFamilies: FamilyItem[] = [];

	const sortBySlug = (a: any , b: any) => {
		if (a.slug < b.slug) return -1;
		if (a.slug > a.slug) return 1;
		return 0;
	}

	const filterBlanks = (s: any) => {
		return s.slug != "";
	}

	onMount(async () => {
		const [responseSymptoms, responseFamilies] = await Promise.all([
			fetch('/api/symptoms.json'),
			fetch('/api/families.json'),
		])

		if ( responseSymptoms.ok && responseFamilies.ok ) {
			apiSymptoms = (await responseSymptoms.json())
				.filter(filterBlanks)
				.sort(sortBySlug);

			apiFamilies = (await responseFamilies.json())
				.filter(filterBlanks)
				.sort(sortBySlug);
		}
	});

	let symptomFilterSearch = '';
	$: filteredSymptoms = apiSymptoms.filter((s) => s.name.toLowerCase().includes(symptomFilterSearch.toLowerCase()));

	let familyFilterSearch = '';
	$: filteredFamilies = apiFamilies.filter((s) => s.name.toLowerCase().includes(familyFilterSearch.toLowerCase()));

	export const addFilter = (type: string, term: string) => {
		if (!['affects', 'symptoms', 'text', 'families'].includes(type)) {
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

	export const filterExists = (filter: FilterItem): boolean => {
		return $filters.some((f) => {
			return f.type == filter.type && f.term == filter.term;
		});
	}

	const handleCheckboxFilter = (type: string, term: string) => {
		let filter: FilterItem = {
			type: type,
			term: term
		};

		return filterExists(filter) ? removeFilter(filter) : addFilter(type, term);
	}

	const handleRadioFilter = (type: string, term: string) => {
		let filter: FilterItem = {
			type: type,
			term: term
		};

		if ( filterExists(filter) ) {
			removeFilter(filter)
			return
		}

		removeFilterByType(filter.type);

		return addFilter(type, term);
	}

	const removeFilterByType = (type: string) => {
		$filters = $filters.filter((f) => {
			return f.type != type;
		});
	}

	const removeFilter = (filter: FilterItem) => {
		let index = $filters.findIndex((f) => {
			return f.type == filter.type && f.term == filter.term;
		});

		if (index != -1) {
			$filters.splice(index, 1);
		}

		$filters = $filters;
	}

	const handleSearchInput = (e: KeyboardEvent): void => {
		if (e.key != 'Enter') {
			return;
		}

		let searchInputElement = <HTMLInputElement>document.getElementById('filter-search');

		if (searchInputElement.value.trim() != '') {
			addFilter('text', searchInputElement.value.trim());
			searchInputElement.value = '';
		}
	}

	const onWindowKeydown = (e: KeyboardEvent): void => {
		if (e.key !== '/') {
			return;
		}

		e.preventDefault();
		document.getElementById('filter-search')?.focus();
	}

	const focus = (n: HTMLInputElement) => {
		n.focus();
	}

	const handleFamilyDropdownUnfocus = ({ relatedTarget, currentTarget }) => {
		familyOpen = handleFilterDropdownUnfocus({ relatedTarget, currentTarget });
	}

	const handleAffectsDropdownUnfocus = ({ relatedTarget, currentTarget }) => {
		affectsOpen = handleFilterDropdownUnfocus({ relatedTarget, currentTarget });
	}

	const handleSymptomsDropdownUnfocus = ({ relatedTarget, currentTarget }) => {
		symptomsOpen = handleFilterDropdownUnfocus({ relatedTarget, currentTarget });
	}

	const handleSeverityDropdownUnfocus = ({ relatedTarget, currentTarget }) => {
		severityOpen = handleFilterDropdownUnfocus({ relatedTarget, currentTarget });
	}

	const handleFilterDropdownUnfocus = ({ relatedTarget, currentTarget }) => {
		return currentTarget.contains(relatedTarget);
	}
</script>

<section aria-labelledby="filter-heading">
	<div class="border-b border-gray-200 bg-white py-2 px-6">
		<div class="mx-auto flex max-w-7xl items-center justify-between">
			<div class="relative inline-block text-left w-full md:w-1/2 lg:w-1/2">
				<div class="relative my-2 flex items-center mr-2">
					<div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
						<svg class="-ml-1 mr-3 h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
						</svg>
					</div>
					<input
						type="text"
						name="filter-search"
						id="filter-search"
						class="block h-12 w-full rounded-md text-gray-400 border-gray-300 pl-8 pr-14 shadow-sm focus:border-green-400 focus:ring-green-400 sm:text-sm lg:text-base"
						use:focus
						on:keydown={handleSearchInput}
					/>

					<div class="absolute inset-y-0 right-0 py-1.5 pr-1.5 hidden md:flex lg:flex text-gray-400">
						<kbd class="ml-2 inline-flex items-center rounded border border-gray-200 px-2 font-sans text-sm font-medium text-gray-400">/</kbd>
					</div>
				</div>
			</div>

			<div class="hidden lg:block md:block">
				<div class="-mx-4 flex items-center divide-x divide-gray-200">
					<div class="relative inline-block px-4 text-left" on:focusout={handleAffectsDropdownUnfocus} >
						<button
							type="button"
							class="group inline-flex justify-center text-sm font-medium text-gray-700 hover:text-gray-900"
							aria-expanded="false"
							on:click={() => {
								affectsOpen = !affectsOpen;
								symptomsOpen = false;
								familyOpen = false;
								severityOpen = false
							}}
						>
							<span>Affects</span>

							<span class="ml-1.5 rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700">
								{$filters.filter((f) => {
									return f.type == 'affects';
								}).length}
							</span>

							<svg class="-mr-1 ml-1 h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
								<path fill-rule="evenodd" d="M10 3a.75.75 0 01.55.24l3.25 3.5a.75.75 0 11-1.1 1.02L10 4.852 7.3 7.76a.75.75 0 01-1.1-1.02l3.25-3.5A.75.75 0 0110 3zm-3.76 9.2a.75.75 0 011.06.04l2.7 2.908 2.7-2.908a.75.75 0 111.1 1.02l-3.25 3.5a.75.75 0 01-1.1 0l-3.25-3.5a.75.75 0 01.04-1.06z" clip-rule="evenodd" />
							</svg>
						</button>

						<div class:hidden={!affectsOpen} class="absolute space-y-2 p-3 right-0 z-10 mt-2 origin-top-right rounded-md bg-white shadow-2xl ring-1 ring-black ring-opacity-5 focus:outline-none">
							{#each getAllAnimals() as animal}
								<button class="w-full text-left whitespace-nowrap text-sm font-medium text-gray-900">
									<input
										id="filter-affects-{animal}"
										name="affects[]"
										value={animal}
										type="checkbox"
										class="h-4 w-4 mb-4 mr-1 rounded border-gray-300 text-green-600 focus:ring-green-500"
										on:click={() => handleCheckboxFilter('affects', animal)}
										checked={$filters.some((f) => {
											return f.type == 'affects' && f.term == animal;
										})}
									/>
									<label for="filter-affects-{animal}">
										<span class="inline-flex cursor-pointer items-center rounded-md bg-{getByAnimal(animal).background} px-2.5 py-1 text-{getByAnimal(animal).foreground} border-solid border border-{getByAnimal(animal).foreground}/50">
											<svg id="emoji" class="h-6 w-6 pr-1" viewBox="0 0 72 72" xmlns="http://www.w3.org/2000/svg">
												{@html getByAnimal(animal).svg}
											</svg> 
											{normalizeAnimal(animal)}
										</span>
									</label>
								</button>
							{/each}
						</div>
					</div>

					<div class="relative inline-block px-4 text-left" on:focusout={handleSymptomsDropdownUnfocus}>
						<button
							type="button"
							class="group inline-flex justify-center text-sm font-medium text-gray-700 hover:text-gray-900"
							aria-expanded="false"
							on:click={() => {
								symptomsOpen = !symptomsOpen;
								affectsOpen = false;
								familyOpen = false;
								severityOpen = false
							}}
						>
							<span>Symptoms</span>
							<span class="ml-1.5 rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700"
								>{$filters.filter((f) => {
									return f.type == 'symptoms';
								}).length}</span
							>
							<svg class="-mr-1 ml-1 h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
								<path fill-rule="evenodd" d="M10 3a.75.75 0 01.55.24l3.25 3.5a.75.75 0 11-1.1 1.02L10 4.852 7.3 7.76a.75.75 0 01-1.1-1.02l3.25-3.5A.75.75 0 0110 3zm-3.76 9.2a.75.75 0 011.06.04l2.7 2.908 2.7-2.908a.75.75 0 111.1 1.02l-3.25 3.5a.75.75 0 01-1.1 0l-3.25-3.5a.75.75 0 01.04-1.06z" clip-rule="evenodd" />
							</svg>
						</button>

						<div class:hidden={!symptomsOpen} class="absolute text-center right-0 z-10 mt-2 origin-top-right rounded-md bg-white shadow-2xl ring-1 ring-black ring-opacity-5 focus:outline-none max-h-96 w-64 overflow-y-scroll">
							<div class="relative border-b">
								<div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-6">
									<svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
										<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
									</svg>
								</div>
								<input bind:value={symptomFilterSearch} type="text" id="filter-symptom-search" class="my-2 pl-8 w-56 rounded-md text-gray-400 border-gray-300 shadow-sm focus:border-green-400 focus:ring-green-400 text-sm" />
							</div>
							{#if filteredSymptoms.length}
								{#each filteredSymptoms as symptom, i}
									<button class="text-left w-full px-3 py-2 whitespace-nowrap text-sm font-medium text-gray-900 block">
										<input
											id="filter-symptom-{i}"
											name="symptom[]"
											value="white"
											type="checkbox"
											class="h-4 w-4 mr-2 rounded border-gray-300 text-green-600 focus:ring-green-500"
											on:click={() => handleCheckboxFilter('symptoms', symptom.name)}
											checked={$filters.some((f) => {
												return f.type == 'symptoms' && f.term == symptom.name;
											})}
										/>
										<label for="filter-symptom-{i}" class="cursor-pointer">
											{symptom.name}
											<span class="rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700">{symptom.count}</span>
										</label>
									</button>
								{/each}
							{:else}
								<div class="items-center my-4 text-center text-gray-400">no matches</div>
							{/if}
						</div>
					</div>

					<div class="relative inline-block px-4 text-left" on:focusout={handleFamilyDropdownUnfocus}>
						<button
							type="button"
							class="group inline-flex justify-center text-sm font-medium text-gray-700 hover:text-gray-900"
							aria-expanded="false"
							on:click={() => {
								familyOpen = !familyOpen;
								affectsOpen = false;
								symptomsOpen = false;
							}}
						>
							<span>Family</span>

							<span class="ml-1.5 rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700">
								{$filters.filter((f) => {
									return f.type == 'families'
								}).length}
							</span>
							<svg class="-mr-1 ml-1 h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
								<path fill-rule="evenodd" d="M10 3a.75.75 0 01.55.24l3.25 3.5a.75.75 0 11-1.1 1.02L10 4.852 7.3 7.76a.75.75 0 01-1.1-1.02l3.25-3.5A.75.75 0 0110 3zm-3.76 9.2a.75.75 0 011.06.04l2.7 2.908 2.7-2.908a.75.75 0 111.1 1.02l-3.25 3.5a.75.75 0 01-1.1 0l-3.25-3.5a.75.75 0 01.04-1.06z" clip-rule="evenodd" />
							</svg>
						</button>

						<div class:hidden={!familyOpen} class="absolute text-center right-0 z-10 mt-2 origin-top-right rounded-md bg-white shadow-2xl ring-1 ring-black ring-opacity-5 focus:outline-none max-h-96 w-64 overflow-y-scroll">
							<div class="relative border-b">
								<div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-6">
									<svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
										<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
									</svg>
								</div>
								<input bind:value={familyFilterSearch} type="text" id="filter-symptom-search" class="my-2 pl-8 w-56 rounded-md text-gray-400 border-gray-300 shadow-sm focus:border-green-400 focus:ring-green-400 text-sm" />
							</div>
							{#if filteredFamilies.length}
								{#each filteredFamilies as family, i}
									<button class="text-left w-full px-3 py-2 whitespace-nowrap text-sm font-medium text-gray-900">
										<input
											id="filter-family-{i}"
											name="family"
											value="white"
											type="radio"
											class="h-4 w-4 mr-2 rounded border-gray-300 text-green-600 focus:ring-green-500"
											on:click={() => handleRadioFilter('families', family.name)}
											checked={$filters.some((f) => {
												return f.type == 'families' && f.term == family.name;
											})}
										/>
										<label for="filter-family-{i}" class="cursor-pointer">
											{family.name}
											<span class="rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700">{family.count}</span>
										</label>
									</button>
								{/each}
							{:else}
								<div class="items-center my-4 text-center text-gray-400">no matches</div>
							{/if}
						</div>
					</div>

					<div class="hidden relative inline-block px-4 text-left" on:focusout={handleSeverityDropdownUnfocus}>
						<button
							type="button"
							class="group inline-flex justify-center text-sm font-medium text-gray-700 hover:text-gray-900"
							aria-expanded="false"
							on:click={() => {
								severityOpen = !severityOpen;
								affectsOpen = false;
								symptomsOpen = false;
								familyOpen = false
							}}
						>
							<span>Severity</span>

							<span class="ml-1.5 rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700">
								{$filters.filter((f) => {
									return f.type == 'families'
								}).length}
							</span>
							<svg class="-mr-1 ml-1 h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
								<path fill-rule="evenodd" d="M10 3a.75.75 0 01.55.24l3.25 3.5a.75.75 0 11-1.1 1.02L10 4.852 7.3 7.76a.75.75 0 01-1.1-1.02l3.25-3.5A.75.75 0 0110 3zm-3.76 9.2a.75.75 0 011.06.04l2.7 2.908 2.7-2.908a.75.75 0 111.1 1.02l-3.25 3.5a.75.75 0 01-1.1 0l-3.25-3.5a.75.75 0 01.04-1.06z" clip-rule="evenodd" />
							</svg>
						</button>

						<div class:hidden={!severityOpen} class="absolute text-center right-0 z-10 mt-2 origin-top-right rounded-md bg-white shadow-2xl ring-1 ring-black ring-opacity-5 focus:outline-none max-h-96 w-64 overflow-y-scroll">
							{#if filteredFamilies.length}
								{#each filteredFamilies as family, i}
									<button class="text-left w-full px-3 py-2 whitespace-nowrap text-sm font-medium text-gray-900">
										<!-- <input
											id="filter-family-{i}"
											name="family"
											value="white"
											type="radio"
											class="h-4 w-4 mr-2 rounded border-gray-300 text-green-600 focus:ring-green-500"
											on:click={() => handleRadioFilter('toxicity', '')}
											checked={$filters.some((f) => {
												return f.type == 'families' && f.term == family.name;
											})}
										/> -->
										<label for="filter-family-{i}" class="cursor-pointer">
											{family.name}
											<span class="rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700">{family.count}</span>
										</label>
									</button>
								{/each}
							{:else}
								<div class="items-center my-4 text-center text-gray-400">no matches</div>
							{/if}
						</div>
					</div>

				</div>
			</div>
		</div>
	</div>

	{#if $filters.length}
		<div class="bg-gray-50 px-6">
			<div class="mx-auto max-w-7xl">
				<div transition:slide={{ duration: 300, axis: 'y' }}>
					<div class="mx-auto max-w-7xl pt-4 flex items-center">
						<h3 class="text-sm font-medium text-gray-500">Filters</h3>

						<div aria-hidden="true" class="h-5 w-px bg-gray-300 ml-4 block" />

						<div class="ml-4">
							<div class="flex flex-wrap items-center">
								{#each $filters as filter, i}
									{#if filter.type == 'affects'}
										<AffectFilterBadge {filter} on:click={() => removeFilter(filter)} />
									{:else if filter.type == 'symptoms'}
										<SymptomFilterBadge {filter} on:click={() => removeFilter(filter)} />
									{:else if filter.type == 'families'}
										<FamilyFilterBadge {filter} on:click={() => removeFilter(filter)} />
									{:else}
										<TextFilterBadge {filter} on:click={() => removeFilter(filter)} />
									{/if}
								{/each}
								{#if $filters.length && resultCount != 0}
									<span class="ml-1.5 rounded bg-gray-200 py-0.5 px-1.5 text-xs font-semibold tabular-nums text-gray-700">
										{#if resultCount == 1}
											1 match
										{:else}
											{resultCount} matches
										{/if}
									</span>
								{/if}

								<div class="ml-2 mb-1">
									<button
										type="button"
										class="group inline-flex justify-center text-xs text-green-600 hover:text-green-500"
										aria-expanded="false"
										on:click={() => {
											$filters = [];
											symptomsOpen = false;
											affectsOpen = false;
										}}
									>
										<span>clear all</span>
									</button>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	{/if}
</section>

<svelte:window on:keydown={onWindowKeydown} />