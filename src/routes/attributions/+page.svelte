<script lang="ts">
	import { Seo } from '$components';
	import { getByAnimal, getAllAnimals } from '$lib/utils/animals';
</script>

<div class="relative bg-gray-50 mb-4 p-4 border-y">
	<div class="relative mx-auto max-w-7xl">
		<div class="overflow-hidden bg-white shadow-lg mb-4 pb-8 rounded-lg">
			<h2 class="unstyled bg-gradient-to-tr from-green-700 to-emerald-900 p-2.5 text-2xl font-extralight text-white text-center mb-8">attributions</h2>
			<div class="prose prose-lg prose-slate px-4 mx-auto text-gray-500">
				<h2 class="block text-center text-3xl font-bold leading-8 tracking-tight text-gray-900 sm:text-4xl">Cite your sources.</h2>

				<p class="text-xl leading-8 text-gray-500 text-justify">
					This site stands on the shoulder of giants. This is where we give credit where credit is due. If you see something that is not properly attributed, please <a href="https://twitter.com/wilhelm" class="border-b hover:text-slate-500">let us know</a>.
				</p>

				<h3>Icons</h3>

				<ul class="text-gray-400">
					<li>tba</li>
				</ul>

				<h3>Images</h3>

				<ul class="text-gray-400">
					<li>tba</li>
				</ul>

				<h3>Data Sources</h3>

				<ul class="text-gray-400">
					<li>tba</li>
				</ul>
			</div>
		</div>
	</div>
</div>

<Seo title="Attributions" canonical="https://plantsm.art/attributions" description="Credit and cited sources of all used external works." />
