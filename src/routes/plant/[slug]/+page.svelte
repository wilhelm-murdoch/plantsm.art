<script lang="ts">
	import { ColumnLeft, ColumnRight, Warning, Seo } from '$components';
	import { isDeadly } from '$lib/types/plant';
	import { getImageUrl } from '$lib/utils/urls';

	export let data: any;
</script>

<div class="relative bg-gray-50 pb-5 mb-8 sm:px-6 sm:pt-4 lg:px-8 lg:pt-4 lg:pb-4 border-y px-6 pt-4">
	<div class="relative mx-auto max-w-7xl">
		<div class="mb-4">
			{#if isDeadly(data.plant)}
				<h1 class="unstyled bg-black p-2.5 text-4xl font-extralight text-white rounded-tr-md rounded-tl-md">
					<svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" viewBox="0 0 512 512" class="svg-icon inline-block outline-none fill-current w-5 h-5 -translate-y-0.5" focusable="false" data-testid="svg-icon">
						<path
							d="M416 400V464C416 490.5 394.5 512 368 512H320V464C320 455.2 312.8 448 304 448C295.2 448 288 455.2 288 464V512H224V464C224 455.2 216.8 448 208 448C199.2 448 192 455.2 192 464V512H144C117.5 512 96 490.5 96 464V400C96 399.6 96 399.3 96.01 398.9C37.48 357.8 0 294.7 0 224C0 100.3 114.6 0 256 0C397.4 0 512 100.3 512 224C512 294.7 474.5 357.8 415.1 398.9C415.1 399.3 416 399.6 416 400V400zM160 192C124.7 192 96 220.7 96 256C96 291.3 124.7 320 160 320C195.3 320 224 291.3 224 256C224 220.7 195.3 192 160 192zM352 320C387.3 320 416 291.3 416 256C416 220.7 387.3 192 352 192C316.7 192 288 220.7 288 256C288 291.3 316.7 320 352 320z"
						/>
					</svg>
					{data.plant.name}
				</h1>
				<Warning />
			{:else}
				<h1 class="unstyled bg-gradient-to-tr from-green-700 to-emerald-900 p-2.5 text-4xl font-extralight text-white rounded-md">
					{data.plant.name}
				</h1>
			{/if}
		</div>

		<div class="mx-auto grid max-w-none grid-cols-1 lg:grid-cols-2 gap-4">
			<ColumnLeft plant={data.plant} />
			<ColumnRight plant={data.plant} />
		</div>
	</div>
</div>

<Seo title={data.plant.name} canonical="https://plantsm.art/plant/{data.plant.pid}" description="Affected animals, common names and symptoms for {data.plant.name}." image={getImageUrl(data.plant.images[0].relative_path, 'medium')} attribution={data.plant.images[0].attribution} />
