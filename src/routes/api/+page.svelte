<script lang="ts">
	import SvelteSeo from 'svelte-seo';
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

type Classification struct {
	// The assigned kingdom of the plant
	Kingdom string   \`json:"kingdom"\`

	// A list of any clades assigned to the plant
	Clades  []string \`json:"clades"\`

	// The assigned order of the plant
	Order   string   \`json:"order"\`

	// The assigned family of the plant
	Family  string   \`json:"family"\`

	// The assigned genus of the plant
	Genus   string   \`json:"genus"\`

	// The assigned species for the plant ( if applicable )
	Species string   \`json:"species"\`
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

	// A timestamp representing the latest changes
	DateLastUpdated string \`json:"date_last_updated"\`

	// An object representing classification data ( see above )
	Classification Classification \`json:"classification"\`
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

	const twitterMeta = {
		site: '@wilhelm',
		title: 'Plant Smart &middot; API Documentation',
		description: 'A free service that aims to provide a detailed listing of dangerous plants for your pets.',
		image: 'https://plantsm.art/images/og-cover.png',
		imageAlt: 'Plant Smart'
	};

	const openGraphMeta = {
		title: 'Plant Smart &middot; API Documentation',
		description: 'A free service that aims to provide a detailed listing of dangerous plants for your pets.',
		type: 'article',
		url: 'https://plantsm.art/api',
		images: [
			{
				url: 'https://plantsm.art/images/og-cover.png',
				alt: 'Plant Smart'
			}
		]
	};
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
									<span class="inline-flex items-center rounded-md bg-{getByAnimal(animal).background} px-2.5 py-1 text-{getByAnimal(animal).foreground}">{getByAnimal(animal).emoji}</span> <a href="/{animal}.json"><code>{animal}.json</code></a>
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

<SvelteSeo title="Plant Smart &middot; API Documentation" description="A free service that aims to provide a detailed listing of dangerous plants for your pets." canonical="https://plantsm.art/api" twitter={twitterMeta} openGraph={openGraphMeta} />
