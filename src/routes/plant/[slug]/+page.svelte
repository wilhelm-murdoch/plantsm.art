<script lang="ts">
	import Carousel from 'svelte-carousel';
	import { animalToBadge } from '../../../components/Card/Card';

	export let data: any;
</script>

<article>
	<div class="shadow-lg">
		{#if data.plant.symptoms.includes('death')}
			<h1
				class="unstyled bg-black p-2.5 text-2xl font-extralight text-white rounded-tr-md rounded-tl-md"
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					x="0px"
					y="0px"
					viewBox="0 0 512 512"
					class="svg-icon inline-block outline-none fill-current w-4 h-4 -translate-y-0.5"
					focusable="false"
					data-testid="svg-icon"
					><path
						d="M416 400V464C416 490.5 394.5 512 368 512H320V464C320 455.2 312.8 448 304 448C295.2 448 288 455.2 288 464V512H224V464C224 455.2 216.8 448 208 448C199.2 448 192 455.2 192 464V512H144C117.5 512 96 490.5 96 464V400C96 399.6 96 399.3 96.01 398.9C37.48 357.8 0 294.7 0 224C0 100.3 114.6 0 256 0C397.4 0 512 100.3 512 224C512 294.7 474.5 357.8 415.1 398.9C415.1 399.3 416 399.6 416 400V400zM160 192C124.7 192 96 220.7 96 256C96 291.3 124.7 320 160 320C195.3 320 224 291.3 224 256C224 220.7 195.3 192 160 192zM352 320C387.3 320 416 291.3 416 256C416 220.7 387.3 192 352 192C316.7 192 288 220.7 288 256C288 291.3 316.7 320 352 320z"
					/></svg
				>
				{data.plant.name}
			</h1>
		{:else}
			<h1
				class="unstyled bg-gradient-to-tr from-green-700 to-emerald-900 p-2.5 text-2xl font-extralight text-white rounded-tr-md rounded-tl-md"
			>
				{data.plant.name}
			</h1>
		{/if}
		<Carousel
			autoplay
			autoplayProgressVisible
			autoplayDuration={5000}
			particlesToShow={3}
			dots={false}
			let:showPrevPage
			let:showNextPage
		>
			<div
				slot="prev"
				on:keydown={showPrevPage}
				on:click={showPrevPage}
				class="custom-arrow custom-arrow-prev"
			>
				<i />
			</div>
			{#each data.plant.images as image}
				<img src="/{image.url}" alt="" class="h-64 w-full object-cover" />
			{/each}
			<div
				slot="next"
				on:keydown={showNextPage}
				on:click={showNextPage}
				class="custom-arrow custom-arrow-next"
			>
				<i />
			</div>
		</Carousel>
		{#if data.plant.symptoms.includes('death')}
			<div class="h-4 bg-black rounded-br-md rounded-bl-md" />
		{:else}
			<div
				class="h-4 bg-gradient-to-tr from-green-700 to-emerald-900 rounded-br-md rounded-bl-md"
			/>
		{/if}
	</div>
	<div class="rounded-md bg-red-800 text-white p-4 mt-4">
		<div class="flex">
			<div class="flex-shrink-0">
				<!-- Heroicon name: mini/x-circle -->
				<svg
					xmlns="http://www.w3.org/2000/svg"
					x="0px"
					y="0px"
					viewBox="0 0 512 512"
					class="svg-icon inline-block outline-none fill-current w-5 h-5 -translate-y-0.5"
					focusable="false"
					data-testid="svg-icon"
					><path
						d="M416 400V464C416 490.5 394.5 512 368 512H320V464C320 455.2 312.8 448 304 448C295.2 448 288 455.2 288 464V512H224V464C224 455.2 216.8 448 208 448C199.2 448 192 455.2 192 464V512H144C117.5 512 96 490.5 96 464V400C96 399.6 96 399.3 96.01 398.9C37.48 357.8 0 294.7 0 224C0 100.3 114.6 0 256 0C397.4 0 512 100.3 512 224C512 294.7 474.5 357.8 415.1 398.9C415.1 399.3 416 399.6 416 400V400zM160 192C124.7 192 96 220.7 96 256C96 291.3 124.7 320 160 320C195.3 320 224 291.3 224 256C224 220.7 195.3 192 160 192zM352 320C387.3 320 416 291.3 416 256C416 220.7 387.3 192 352 192C316.7 192 288 220.7 288 256C288 291.3 316.7 320 352 320z"
					/></svg
				>
			</div>
			<div class="ml-3">
				<h3 class="text-md font-medium text-white">
					This plant is known to be potentially deadly!
				</h3>
				<div class="mt-2 text-sm">
					<ul class="list-disc space-y-1 pl-5">
						<li>
							If you're comfortable with the risk, ensure this plant is out of reach of your animal.
						</li>
						<li>
							If you suspect your animal has ingested portions of this plant, immediately consult
							your veternarian.
						</li>
						<li>
							If your animal is experiencing any of the following symptoms, immediately contact your
							veternarian's emergency number for further instructions.
						</li>
					</ul>
				</div>
			</div>
		</div>
	</div>
	<div class="mx-auto grid max-w-lg gap-5 mt-4 lg:max-w-none lg:grid-cols-2">
		<div class="flex flex-col overflow-hidden rounded-lg shadow-lg">
			<div class="flex-shrink-0">
				<h3
					class="unstyled font-sans mb-3 p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500"
				>
					Other Names
				</h3>
				<div class="text-md">
					{#if data.plant.common}
						<ul>
							{#each data.plant.common as name}
								<li>{name}</li>
							{/each}
						</ul>
					{:else}{/if}
				</div>
			</div>
			<h3
				class="unstyled font-sans mb-3 p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500"
			>
				Symptoms
			</h3>
			<div class="text-md">
				<ul class="list-decimal">
					{#each data.plant.symptoms as symptom}
						<li>
							<a
								href="/"
								title="View other plants with this symptom."
								class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
								>{symptom}</a
							>
						</li>
					{/each}
				</ul>
			</div>
		</div>

		<div class="flex flex-col overflow-hidden rounded-lg shadow-lg">
			<div class="flex-shrink-0">
				<h3
					class="unstyled font-sans mb-3 p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500"
				>
					Affects
				</h3>
				<p class="px-2.5 pb-2.5 text-md font-medium space-x-1 space-y-1">
					{#each data.plant.animals as animal}
						<span
							class="inline-flex items-center rounded-md {animalToBadge[animal]
								.background} p-1 {animalToBadge[animal].foreground}"
							>{animalToBadge[animal].emoji} {animal}s</span
						>{/each}
				</p>
				<h3
					class="unstyled font-sans p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500"
				>
					Classification
				</h3>

				<div class="flex bg-gray-100 border-b">
					<div class="flex w-0 flex-1">
						<span class="p-4 text-slate-400">Kingdom</span>
					</div>
					<div class="-ml-px flex w-0 flex-1">
						<span class="p-4">
							<a
								href="/"
								title={data.plant.classification.kingdom}
								class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
								>{data.plant.classification.kingdom}</a
							></span
						>
					</div>
				</div>

				{#each data.plant.classification.clades as clade}
					<div class="flex border-b">
						<div class="flex w-0 flex-1">
							<span class="p-4 text-slate-400">— Clade</span>
						</div>
						<div class="-ml-px flex w-0 flex-1">
							<span class="p-4">
								<a
									href="/"
									title={clade}
									class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
									>{clade}</a
								></span
							>
						</div>
					</div>
				{/each}

				<div class="flex bg-gray-100 border-b">
					<div class="flex w-0 flex-1">
						<span class="p-4 text-slate-400">Order</span>
					</div>
					<div class="-ml-px flex w-0 flex-1">
						<span class="p-4">
							<a
								href="/"
								title={data.plant.classification.order}
								class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
								>{data.plant.classification.order}</a
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
								href="/"
								title={data.plant.classification.family}
								class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
								>{data.plant.classification.family}</a
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
								href="/"
								title={data.plant.classification.genus}
								class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
								>{data.plant.classification.genus}</a
							></span
						>
					</div>
				</div>
				{#if data.plant.classification.species}
					<div class="flex border-b">
						<div class="flex w-0 flex-1">
							<span class="p-4 text-slate-400">Species</span>
						</div>
						<div class="-ml-px flex w-0 flex-1">
							<span class="p-4">
								<a
									href="/"
									title={data.plant.classification.species}
									class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
									>{data.plant.classification.species}</a
								></span
							>
						</div>
					</div>
				{/if}
			</div>
			<div class="flex">
				<a
					href={data.plant.wikipedia_url}
					title="Read more about {data.plant.name}."
					class=" text-slate-400 relative inline-flex w-0 flex-1 items-center justify-center rounded-br-lg border border-transparent py-2 text-md hover:bg-slate-50 hover:text-gray-500"
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
	</div>

	<div class="mt-4 flex flex-col shadow-lg">
		<div class="-my-2 -mx-4 overflow-x-auto sm:-mx-6 lg:-mx-8">
			<div class="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
				<div class="overflow-hidden md:rounded-lg">
					<table class="min-w-full divide-y divide-gray-300">
						<thead class=" bg-slate-100">
							<tr>
								<th
									scope="col"
									class="py-3.5 pl-4 pr-3 text-left text-md font-semibold  text-slate-500 sm:pl-6"
									>Image Attribution</th
								>
								<th scope="col" class="px-3 py-3.5 text-left text-md font-semibold  text-slate-500"
									>License</th
								>

								<th scope="col" class="relative py-3.5 pl-3 pr-4 sm:pr-6">
									<span class="sr-only">Edit</span>
								</th>
							</tr>
						</thead>
						<tbody class="divide-y divide-gray-200 bg-white">
							{#each data.plant.images as image}
								<tr>
									<td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm sm:pl-6">
										<div class="flex items-center">
											<div class="h-10 w-10 flex-shrink-0">
												<img class="h-10 w-10 rounded-full" src="/{image.url}" alt="" />
											</div>
											<div class="ml-4">
												<div class="text-gray-500">{image.attribution}</div>
											</div>
										</div>
									</td>
									<td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
										<div class="text-slate-500">
											<a
												href="https://creativecommons.org/licenses/{image.license.replace(
													'cc-',
													''
												)}/2.0/"
												title="">{image.license}</a
											>
										</div>
									</td>
									<td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
										<span
											class="inline-flex rounded-full bg-green-100 px-2 text-xs font-semibold leading-5 text-green-800"
											>Active</span
										>
									</td>
								</tr>
							{/each}
						</tbody>
					</table>
				</div>
			</div>
		</div>
	</div>
</article>

<style>
	.custom-arrow {
		width: 20px;
		background-color: #000000;
		opacity: 0.3;
		position: absolute;
		top: 0;
		bottom: 0;
		z-index: 1;
		transition: opacity 150ms ease;
		display: flex;
		align-items: center;
		justify-content: center;
		cursor: pointer;
		-webkit-tap-highlight-color: transparent;
	}

	.custom-arrow:hover {
		opacity: 0.5;
	}

	.custom-arrow > i {
		border: solid white;
		border-width: 0 5px 5px 0;
		padding: 5px;
		position: relative;
	}

	.custom-arrow-prev {
		left: 0;
	}

	.custom-arrow-prev > i {
		transform: rotate(135deg);
		right: -4px;
	}

	.custom-arrow-next {
		right: 0;
	}

	.custom-arrow-next > i {
		transform: rotate(-45deg);
		left: -4px;
	}
</style>
