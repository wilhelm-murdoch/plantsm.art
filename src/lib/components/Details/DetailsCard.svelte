<script lang="ts">
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';
	import { getAllAnimals, getByAnimal } from '$utils/animals';

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
		<h3
			class="unstyled font-sans mb-3 p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500"
		>
			Affects
		</h3>
		<p class="px-2.5 pb-3 text-md font-medium space-x-1 space-y-1">
			{#each plant.animals as animal}
				<span
					class="inline-flex items-center rounded-md bg-{getByAnimal(animal)
						.background} px-2.5 py-1 text-{getByAnimal(animal).foreground}"
					>{getByAnimal(animal).emoji} {animal}</span
				>{/each}
		</p>

		<h3
			class="unstyled font-sans mb-3 p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500"
		>
			Other Names
		</h3>
		<div class="text-md px-2.5 pb-2.5 text-gray-400">
			{#if plant.common}
				{#each plant.common as name, i}
					{name.name}{#if i != plant.common.length - 1},&nbsp;{/if}
				{/each}
			{:else}
				<div class="rounded-md bg-slate-100 p-4">
					<div class="flex">
						<div class="flex-shrink-0">
							<svg
								class="h-5 w-5 text-blue-400"
								xmlns="http://www.w3.org/2000/svg"
								viewBox="0 0 20 20"
								fill="currentColor"
								aria-hidden="true"
							>
								<path
									fill-rule="evenodd"
									d="M19 10.5a8.5 8.5 0 11-17 0 8.5 8.5 0 0117 0zM8.25 9.75A.75.75 0 019 9h.253a1.75 1.75 0 011.709 2.13l-.46 2.066a.25.25 0 00.245.304H11a.75.75 0 010 1.5h-.253a1.75 1.75 0 01-1.709-2.13l.46-2.066a.25.25 0 00-.245-.304H9a.75.75 0 01-.75-.75zM10 7a1 1 0 100-2 1 1 0 000 2z"
									clip-rule="evenodd"
								/>
							</svg>
						</div>
						<div class="ml-3 flex-1 md:flex md:justify-between">
							<p class="text-sm text-slate-500">
								<i
									>We currently don't have any other common names for this entry on record. Feel
									free to learn about how you can <a
										href="/"
										title="Missing details? Submit your own changes for approval."
										class="border-b hover:text-slate-600">contribute</a
									> to make this page more complete.</i
								>
							</p>
						</div>
					</div>
				</div>
			{/if}
		</div>
	</div>
	<h3 class="unstyled font-sans p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500">
		Symptoms
	</h3>
	<div class="text-md p-2.5">
		{#each plant.symptoms as symptom, i}
			<a
				href="/?f=s:{symptom.name}"
				title="View other plants with this symptom."
				class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
				>{symptom.name}</a
			>{#if i != plant.symptoms.length - 1},&nbsp;{/if}
		{/each}
	</div>
	<h3 class="unstyled font-sans p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500">
		Images
	</h3>
	<div class="grid grid-cols-4">
		{#each plant.images as image, i}
			<img
				on:click={() => openLightbox(image, i)}
				on:keydown
				alt={image.relative_path}
				class="inline-block object-cover object-center h-32 w-full hover:opacity-75 cursor-pointer"
				src="/images/{image.relative_path.replace('original', 'thumbnail')}"
			/>
		{/each}
	</div>
</div>

{#if lightboxVisible}
	<div
		in:fade
		out:fade
		class="fixed top-0 left-0 z-800 w-screen h-screen bg-black/70 overflow-scroll"
	>
		<div class="w-screen bg-gradient-to-tr from-green-700 to-emerald-900 top-0">
			<a on:click={() => closeLightbox()} href="#top" class="float-right mt-5 mr-5">
				<svg
					class="w-6 h-6 text-white"
					fill="none"
					stroke="currentColor"
					viewBox="0 0 24 24"
					xmlns="http://www.w3.org/2000/svg"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M6 18L18 6M6 6l12 12"
					/>
				</svg>
			</a>
			<div class="p-2.5 text-3xl font-extralight text-white text-center">
				{lightboxImageIndex + 1} of {plant.images.length}
			</div>
		</div>
		<span
			on:click={handlePreviousClick}
			on:keypress={onLightboxKeyDown}
			class="fixed top-1/2 text-white bg-black/20 hover:bg-black/40 p-2.5 rounded-full ml-2 cursor-pointer"
		>
			<svg
				class="w-6 h-6"
				fill="none"
				stroke="currentColor"
				viewBox="0 0 24 24"
				xmlns="http://www.w3.org/2000/svg"
				><path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M15 19l-7-7 7-7"
				/></svg
			>
		</span>
		<span
			on:click={handleNextClick}
			on:keypress={onLightboxKeyDown}
			class="fixed top-1/2 right-0 text-white bg-black/20 hover:bg-black/40 p-2.5 rounded-full mr-2 cursor-pointer"
		>
			<svg
				class="w-6 h-6"
				fill="none"
				stroke="currentColor"
				viewBox="0 0 24 24"
				xmlns="http://www.w3.org/2000/svg"
				><path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M9 5l7 7-7 7"
				/></svg
			>
		</span>
		<div class="items-center m-16 border-2">
			<img
				alt={lightboxImage.relative_path}
				src="/images/{lightboxImage.relative_path}"
				class="w-full"
			/>
			<div class="p-2.5 bg-white text-center text-slate-500">{lightboxImage.attribution}</div>
		</div>
	</div>
{/if}

<svelte:window on:keydown={onLightboxKeyDown} />
