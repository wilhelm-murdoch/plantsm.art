<script lang="ts">
	import { lazy } from '$lib/utils/lazy';
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';

	export let plant: any = {};

	let lightboxImage: any = plant.images[0];
	let lightboxImageIndex = 0;
	let lightboxVisible = false;

	let toggleScrollLock = () => {};

	function handlePreviousClick() {
		let previous = lightboxImageIndex - 1;
		if (lightboxImageIndex == 0) {
			previous = plant.images.length - 1;
		} else if (lightboxImageIndex + 1 > plant.images.length) {
			previous = 0;
		}

		lightboxImageIndex = previous;
		lightboxImage = plant.images[previous];
	}

	function handleNextClick() {
		let next = lightboxImageIndex + 1;
		if (lightboxImageIndex == plant.images.length - 1 || next > plant.images.length) {
			next = 0;
		}

		lightboxImageIndex = next;
		lightboxImage = plant.images[next];
	}

	function onLightboxKeyDown(e: any) {
		switch (e.key) {
			case 'Escape':
				lightboxImage = {};
				lightboxImageIndex = 0;
				lightboxVisible = false;

				toggleScrollLock();
				break;
			case 'ArrowLeft':
				handlePreviousClick();
				break;
			case 'ArrowRight':
				handleNextClick();
				break;
		}
	}

	function closeLightbox() {
		lightboxImage = {};
		lightboxImageIndex = 0;
		lightboxVisible = false;

		toggleScrollLock();
	}

	function openLightbox(image: any, index: number) {
		lightboxVisible = true;
		lightboxImage = image;
		lightboxImageIndex = index;

		toggleScrollLock();
	}

	onMount(() => {
		const defaultScroll = document.body.style.overflow;
		toggleScrollLock = () => {
			if (lightboxVisible) {
				document.body.style.overflow = 'hidden';
			} else {
				document.body.style.overflow = defaultScroll;
			}
		};

		toggleScrollLock();
	});
</script>

<div class="flex flex-col overflow-hidden rounded-lg shadow-lg mb-4">
	<div class="flex-shrink-0">
		<h3 class="unstyled font-sans p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500">Images</h3>
		<div class="grid grid-cols-4">
			{#each plant.images as image, i}
				<div class="overflow-hidden background-fallback bg-green-50">
					<img on:click={() => openLightbox(image, i)} on:keydown alt={image.relative_path} class="is-lazy inline-block object-cover object-center h-32 w-full hover:scale-110 ease-in-out duration-100 cursor-pointer" use:lazy={'https://plantsm.art/cdn-cgi/imagedelivery/qnkf0SBjwoY9e50dCCmY3Q/' + image.relative_path + '/thumbnail'} />
				</div>
			{/each}
		</div>
		<h3 class="unstyled font-sans p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500">Classification</h3>

		<div class="flex bg-gray-100 border-b">
			<div class="flex w-0 flex-1">
				<span class="p-4 text-slate-400">Kingdom</span>
			</div>
			<div class="-ml-px flex w-0 flex-1">
				<span class="p-4"> <a href="https://en.wikipedia.org/wiki/{plant.classification.kingdom}" title={plant.classification.kingdom} class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500">{plant.classification.kingdom}</a></span>
			</div>
		</div>

		{#if plant.classification.clades}
			{#each plant.classification.clades as clade}
				<div class="flex border-b">
					<div class="flex w-0 flex-1">
						<span class="p-4 text-slate-400">— Clade</span>
					</div>
					<div class="-ml-px flex w-0 flex-1">
						<span class="p-4"> <a href="https://en.wikipedia.org/wiki/{clade}" title={clade} class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500">{clade}</a></span>
					</div>
				</div>
			{/each}
		{/if}

		<div class="flex bg-gray-100 border-b">
			<div class="flex w-0 flex-1">
				<span class="p-4 text-slate-400">Order</span>
			</div>
			<div class="-ml-px flex w-0 flex-1">
				<span class="p-4"> <a href="https://en.wikipedia.org/wiki/{plant.classification.order}" title={plant.classification.order} class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500">{plant.classification.order}</a></span>
			</div>
		</div>

		<div class="flex border-b">
			<div class="flex w-0 flex-1">
				<span class="p-4 text-slate-400">Family</span>
			</div>
			<div class="-ml-px flex w-0 flex-1">
				<span class="p-4"> <a href="https://en.wikipedia.org/wiki/{plant.classification.family}" title={plant.classification.family} class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500">{plant.classification.family}</a></span>
			</div>
		</div>

		<div class="flex bg-gray-100 border-b">
			<div class="flex w-0 flex-1">
				<span class="p-4 text-slate-400">Genus</span>
			</div>
			<div class="-ml-px flex w-0 flex-1">
				<span class="p-4"> <a href="https://en.wikipedia.org/wiki/{plant.classification.genus}" title={plant.classification.genus} class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500">{plant.classification.genus}</a></span>
			</div>
		</div>
		{#if plant.classification.species}
			<div class="flex border-b">
				<div class="flex w-0 flex-1">
					<span class="p-4 text-slate-400">Species</span>
				</div>
				<div class="-ml-px flex w-0 flex-1">
					<span class="p-4"> <a href="https://en.wikipedia.org/wiki/{plant.classification.species}" title={plant.classification.species} class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500">{plant.classification.species}</a></span>
				</div>
			</div>
		{/if}
	</div>
	<div class="flex divide-x -mt-px divide-gray-200">
		<div class="flex w-0 flex-1">
			<a href="https://github.com/wilhelm-murdoch/plantsm.art/issues/new?title=Suggesting%20changes%20for:%20{plant.name}" title="Create an issue for this entry on GitHub." target="_blank" rel="noreferrer" class="text-slate-400 hover:text-gray-500 relative -mr-px inline-flex w-0 flex-1 items-center justify-center rounded-bl-lg border border-transparent py-4 text-sm font-medium">
				Issue
				<svg class="w-5 h-5 ml-1" stroke-width="1.5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" color="currentColor">
					<path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2 2 6.477 2 12s4.477 10 10 10z" stroke-linecap="round" stroke-linejoin="round" stroke="currentColor" />
					<path d="M14.333 19v-1.863c.025-.31-.018-.62-.126-.913a2.18 2.18 0 00-.5-.781c2.093-.227 4.293-1 4.293-4.544 0-.906-.358-1.778-1-2.434a3.211 3.211 0 00-.06-2.448s-.787-.227-2.607.961a9.152 9.152 0 00-4.666 0c-1.82-1.188-2.607-.96-2.607-.96A3.211 3.211 0 007 8.464a3.482 3.482 0 00-1 2.453c0 3.519 2.2 4.291 4.293 4.544a2.18 2.18 0 00-.496.773 2.134 2.134 0 00-.13.902V19M9.667 17.702c-2 .631-3.667 0-4.667-1.948" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" /></svg
				>
			</a>
		</div>
		<div class="-ml-px flex w-0 flex-1">
			<a href={plant.wikipedia_url} title="Read more about this entry." class=" text-slate-400 inline-flex w-0 flex-1 items-center justify-center rounded-br-lg border border-transparent py-2 text-md hover:text-gray-500">
				Wikipedia
				<svg class="w-5 h-5 ml-1" stroke-width="1.5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" color="currentColor">
					<path d="M14 11.998C14 9.506 11.683 7 8.857 7H7.143C4.303 7 2 9.238 2 11.998c0 2.378 1.71 4.368 4 4.873a5.3 5.3 0 001.143.124" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" />
					<path d="M10 11.998c0 2.491 2.317 4.997 5.143 4.997h1.714c2.84 0 5.143-2.237 5.143-4.997 0-2.379-1.71-4.37-4-4.874A5.304 5.304 0 0016.857 7" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" />
				</svg>
			</a>
		</div>
	</div>
</div>

{#if lightboxVisible}
	<div in:fade out:fade class="fixed top-0 left-0 z-800 w-screen h-screen bg-black/70 overflow-y-scroll">
		<div class="w-100 bg-gradient-to-tr from-green-700 to-emerald-900 top-0">
			<a on:click={() => closeLightbox()} href="#top" class="float-right mt-4 mr-4">
				<svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
				</svg>
			</a>
			<div class="p-2.5 text-xl font-extralight text-white text-center">
				{lightboxImageIndex + 1} of {plant.images.length}
			</div>
		</div>
		<span on:click={handlePreviousClick} on:keypress={onLightboxKeyDown} class="fixed top-1/2 text-white bg-black/20 hover:bg-black/40 p-2.5 rounded-full ml-2 cursor-pointer">
			<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" /></svg>
		</span>
		<span on:click={handleNextClick} on:keypress={onLightboxKeyDown} class="fixed top-1/2 right-4 text-white bg-black/20 hover:bg-black/40 p-2.5 rounded-full mr-2 cursor-pointer">
			<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" /></svg>
		</span>
		<div class="items-center m-16 border-2">
			<img alt={lightboxImage.relative_path} src="https://plantsm.art/cdn-cgi/imagedelivery/qnkf0SBjwoY9e50dCCmY3Q/{lightboxImage.relative_path}/large" class="w-full" />
			<div class="p-2.5 bg-white text-center text-slate-500">{lightboxImage.attribution}</div>
		</div>
	</div>
{/if}

<svelte:window on:keydown={onLightboxKeyDown} />
