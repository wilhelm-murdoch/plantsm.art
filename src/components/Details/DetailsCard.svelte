<script lang="ts">
	export let plant: any = {};

	export let symptoms: string[] = [];
	export let common: string[] = [];
</script>

<div class="flex flex-col overflow-hidden rounded-lg shadow-lg mb-4">
	<div class="flex-shrink-0">
		<h3
			class="unstyled font-sans mb-3 p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500"
		>
			Other Names
		</h3>
		<div class="text-md px-2.5 pb-2.5 text-gray-400">
			{#if common}
				{#each common as name, i}
					{name}{#if i != common.length - 1},&nbsp;{/if}
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
									free to check out how you can <a
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
		{#each symptoms as symptom, i}
			<a
				href="/"
				title="View other plants with this symptom."
				class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
				>{symptom
					.split(' ')
					.map((x) => x[0].toUpperCase() + x.substring(1))
					.join(' ')}</a
			>{#if i != symptoms.length - 1},&nbsp;{/if}
		{/each}
	</div>
	<h3 class="unstyled font-sans p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500">
		Images
	</h3>
	<div class="grid grid-cols-4">
		{#each plant.images as image}
			<a href={image.source_url} title="">
				<img
					alt="gallery"
					class="inline-block object-cover object-center h-32 w-full hover:opacity-75 cursor-pointer"
					src="/{image.url}"
				/>
			</a>
		{/each}
	</div>
</div>
