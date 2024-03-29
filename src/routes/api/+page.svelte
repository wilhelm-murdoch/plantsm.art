<script lang="ts">
	import { Seo } from '$components';
	import { getByAnimal, getAllAnimals } from '$lib/utils/animals';

	let snippetPlantsStruct: string = `type Symptom struct {
	// The human-readable name of the symptom
	Name string \`json:"name"\`

	// A machine-friendly version of the name
	Slug string \`json:"slug"\`
}

type Common struct {
	// The human-readable name of the common name
	Name string \`json:"name"\`

	// A machine-friendly version of the common name
	Slug string \`json:"slug"\`
}

type Image struct {
	// The original remote source of the image
	SourceUrl string \`json:"source_url"\`

	// The original author of the image
	Attribution string \`json:"attribution"\`

	// The associated CC license
	License string \`json:"license"\`

	// The path to the image hosted on this site
	RelativePath string \`json:"relative_path"\`
}

type Plant struct {
	// The unique identifier of the plant
	Pid string \`json:"pid"\`

	// The human-friendly name of the plant
	Name string \`json:"name"\`

	// A list of any animals affected by the plant
	Animals []string \`json:"animals"\`

	// A list of common names ( see above )
	Common []Common \`json:"common"\`

	// A list of symptoms ( see above )
	Symptoms []Symptom \`json:"symptoms"\`

	// A list of images ( see above )
	Images []Image \`json:"images"\`

	// The wikipedia entry associated with the plant
	WikipediaUrl string \`json:"wikipedia_url"\`

	// The family classification of the plant
	Family string \`json:"family"\`
}`;

	let snippetSymptomsStruct: string = `type Symptom struct {
	// The human-readable name of the symptom
	Name string \`json:"name"\`

	// A machine-friendly version of the name
	Slug string \`json:"slug"\`

	// The number of plants sharing this symptom
	Count  int      \`json:"count"\`

	// A list of ids for each associated plant
	Plants []string \`json:"plants"\`
}`;

	let snippetAnimalsStruct: string = `type Animal struct {
	// The name of the animal
	Name string \`json:"name"\`

	// The number of plants affecting this animal
	Count  int      \`json:"count"\`

	// A list of ids for each affecting plant
	Plants []string \`json:"plants"\`
}`;
</script>

<div class="relative bg-gray-50 mb-4 p-4 border-y">
	<div class="relative mx-auto max-w-7xl">
		<div class="overflow-hidden bg-white shadow-lg mb-4 pb-8 rounded-lg">
			<h2 class="unstyled bg-gradient-to-tr from-green-700 to-emerald-900 p-2.5 text-2xl font-extralight text-white text-center mb-8">api</h2>
			<div class="prose prose-lg prose-slate px-4 mx-auto text-gray-500">
				<h2 class="block text-center text-3xl font-bold leading-8 tracking-tight text-gray-900 sm:text-4xl">Play around with the data.</h2>

				<p class="text-xl leading-8 text-gray-500 text-justify">
					We've collated <a href="https://www.inaturalist.org/" title="iNaturalist">data</a>
					<a href="https://www.aspca.org/" title="ASPCA">from</a>
					<a href="https://en.wikipedia.org/wiki/Plant" title="Wikipedia">several</a> sources and munged it into a minimal JSON-based static database. The entirety of the site is powered by this single file and you're more than welcome to use it for your own purposes.
				</p>

				<p>
					We also expose several other JSON files which can be used as lookup tables in combination with <code>plants.json</code> to help sort through the data.
				</p>

				<h3>Files & Models</h3>

				<ul class="text-gray-400">
					<li>
						<a href="/api/plants.json"><code>plants.json</code></a>
						<p>Our primary database. This includes a breakdown of all currently-supported plants along with their common names, symptoms, images and affected animal types.</p>
						<p>This is an array of objects with the following structure:</p>
						<pre>{snippetPlantsStruct}</pre>
						<h4>Download</h4>
						<pre>$ curl -sL plantsm.art/api/plants.json > plants.json</pre>
					</li>

					<li>
						<a href="/api/symptoms.json"><code>symptoms.json</code></a>
						<p>A distinct listing of all known symptoms with an associated listing of any related plant records.</p>
						<p>This is an array of objects with the following structure:</p>
						<pre>{snippetSymptomsStruct}</pre>
						<h4>Download</h4>
						<pre>$ curl -sL plantsm.art/api/symptoms.json > symptoms.json</pre>
					</li>

					<li>
						<a href="/api/animals.json"><code>animals.json</code></a>
						<p>A distinct listing of all supported animal types with an associated listing of any related plant records.</p>
						<p>This is an array of objects with the following structure:</p>
						<pre>{snippetAnimalsStruct}</pre>
						<h4>Download</h4>
						<pre>$ curl -sL plantsm.art/api/animals.json > animals.json</pre>
						<h4>Supplemental</h4>
						<p>Use these files if you wish to have animal-specific listings.</p>
						<ul>
							{#each getAllAnimals() as animal}
								<li>
									<span class="inline-flex bg-{getByAnimal(animal).background} p-1 rounded-md align-middle">
										<svg id="emoji" class="h-6 w-6" viewBox="0 0 72 72" xmlns="http://www.w3.org/2000/svg">
											{@html getByAnimal(animal).svg}
										</svg>
									</span>
									<a href="/api/{animal}.json"><code>{animal}.json</code></a>
								</li>
							{/each}
						</ul>
					</li>
				</ul>

				<h3>Usage</h3>
				<p>These are known as "dumb" API endpoints. They are static flat files that are generated at "build time" before we deploy changes to production. They, like this site, live behind a CDN which is managed by Cloudflare. We do not impose rate-limits or any other restrictions with regards to access.</p>
				<p>You may download these files locally without attribution and use them as you wish.</p>
				<p>
					If you wish to use the associated images, be sure to abide by the associated <a href="https://creativecommons.org/" title="Creative Commons">Creative Commons</a>
					license. You can find the specific license for each image using the
					<code>.[].images[].license</code> path in the <code>plants.json</code> dataset.
				</p>
			</div>
		</div>
	</div>
</div>

<Seo title="API Documentation" canonical="https://plantsm.art/api" description="Learn how you can access the Plant Smart API." />
