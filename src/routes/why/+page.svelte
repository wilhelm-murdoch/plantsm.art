<script lang="ts">
	import { Seo } from '$components';
	import { lazy } from '$lib/utils/lazy';
</script>

<div class="relative bg-gray-50 mb-4 p-4 border-y">
	<div class="relative mx-auto max-w-7xl">
		<div class="overflow-hidden bg-white shadow-lg mb-4 pb-8 rounded-lg">
			<h2 class="unstyled bg-gradient-to-tr from-green-700 to-emerald-900 p-2.5 text-2xl font-extralight text-white text-center mb-8">why</h2>
			<div class="prose prose-lg prose-slate px-4 mx-auto text-gray-500 mb-8">
				<h2 class="block text-center text-3xl font-bold leading-8 tracking-tight text-gray-900 sm:text-4xl">Why go through all the trouble?</h2>
				<p class="text-xl leading-8 text-gray-500 text-justify">Simply put, I love plants and I love animals. I can't count how many times I've been to a plant nursery and had to stop and do a Google search to see if something was safe enough to bring home and keep around our little Pandora, or 🐼 for short.</p>
			</div>

			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 mb-8">
				<div class="background-fallback overflow-hidden bg-slate-50">
					<img class="is-lazy hover:scale-125 ease-in-out duration-1000 h-64 w-full object-cover object-center" use:lazy={'/images/why/panda-6.png'} alt="A very cute cat." />
				</div>
				<div class="background-fallback overflow-hidden bg-slate-50 hidden md:block">
					<img class="is-lazy hover:scale-125 ease-in-out duration-1000 h-64 w-full object-cover object-center" use:lazy={'/images/why/panda-2.png'} alt="A very cute cat." />
				</div>
				<div class="background-fallback overflow-hidden bg-slate-50 hidden lg:block">
					<img class="is-lazy hover:scale-125 ease-in-out duration-1000 h-64 w-full object-cover object-center" use:lazy={'/images/why/panda-7.png'} alt="A very cute cat." />
				</div>
			</div>

			<div class="prose prose-lg prose-slate px-4 mx-auto text-gray-500 mb-8">
				<h3>It's as simple as that?</h3>
				<p>That was the idea, at least. I just wanted a single place that had all the information I needed so I could make informed purchasing decisions. I also figured I couldn't possibly be the only person who has this problem. So, I set off to make this project during the holiday break of 2022.</p>
				<p>However, I didn't want to make it <i>too</i> easy for myself...</p>
				<h3>A challenge!</h3>
				<p>
					I work as a weird combination of <a href="https://en.wikipedia.org/wiki/Site_reliability_engineering" title="What does SRE mean?">SRE</a>
					and
					<a href="https://en.wikipedia.org/wiki/DevOps" title="What does DevOps mean?">DevOps engineer</a>. I have plenty of smarts when it comes to the technical side of hosting and running things, but it's been years since I've done any front-end work. Career-wise, I need to be able to speak the same language as teams I have to support. So, understanding how far behind that portion of my skillset had become, I thought this would be the perfect opportunity to brush up.
				</p>
				<h4>The Requirements:</h4>
				<ul>
					<li>Must keep operational costs as close to zero as possible. The only money I've put down on this so far is the $20 AUD to purchase the domain name.</li>
					<li>
						Must be hosted on <a href="https://pages.cloudflare.com/" title="Cloudflare Pages">Cloudflare Pages</a>. I could've chosen
						<a href="https://pages.github.com/" title="GitHub Pages">GitHub Pages</a> to keep everything in one place, but automated Cloudflare builds & deployments work out-of-the-box with minimal configuration. Besides, I've already worked with the latter and I wanted the challenge of trying something new.
					</li>
					<li>
						Must be completely static. There should be no server-side rendering, processing or other explicit backend dependencies to manage. For this, I use the pre-rendering functionality that comes packaged with <a href="https://kit.svelte.dev/" title="SvelteKit">SvelteKit</a>. I've been working with this for only a few weeks now and I'm a full convert.
					</li>
					<li>
						All data must be served statically as well. The entirety of the data set is contained within a single JSON file, which you can view at <a href="https://plantsm.art/plants.json" title="The primary Plant Smart data set.">/plants.json</a>. This is the source of truth for all derivative data sets and lookup tables used on this site. It's effectively what I call a "dumb API". Check out the
						<a href="/api" title="Plant Smart API documentation.">API documentation</a> if you'd like to know more about it.
					</li>
					<li>
						Must use <a href="https://kit.svelte.dev/" title="SvelteKit">SvelteKit</a>,
						<a href="https://www.typescriptlang.org/" title="TypeScript">TypeScript</a>,
						<a href="https://tailwindcss.com/" title="TailwindCSS">TailwindCSS</a>
						and <a href="https://vitejs.dev/" title="Vite">Vite</a>. I had zero working knowledge of any of these and my front-end peers can't seem to shut up about them. So... why not?
					</li>
					<li>
						Must be <i>fast</i>. Everything is static, compressed, cached and sitting behind a world-class CDN. I've worked in a network performance and load testing SaaS for close to 5 years now. I wouldn't be able to look myself in the mirror if I couldn't easily do this one. 😅
					</li>
					<li>
						Must be open source and community-driven. At one point, I would like to take the hands off the wheel and see if other interested parties would like to get involved and help out with managing datasets and fixing bugs. GitHub allows for pretty much all of this. Check how to <a href="/contribute" title="How to contribute.">contribute</a> or see the
						<a href="/updates" title="Latest GitHub contributions.">latest contributions</a>.
					</li>
				</ul>
				<p>So far, it's going quite well. I haven't had any issues with meeting any of these self-imposed development constraints.</p>
			</div>

			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 mb-8">
				<div class="background-fallback overflow-hidden bg-slate-50">
					<img class="is-lazy hover:scale-125 ease-in-out duration-1000 h-64 w-full object-cover object-center" use:lazy={'/images/why/panda-1.png'} alt="A very cute cat." />
				</div>
				<div class="background-fallback overflow-hidden bg-slate-50 hidden md:block">
					<img class="is-lazy hover:scale-125 ease-in-out duration-1000 h-64 w-full object-cover object-center" use:lazy={'/images/why/panda-4.png'} alt="A very cute cat." />
				</div>
				<div class="background-fallback overflow-hidden bg-slate-50 hidden lg:block">
					<img class="is-lazy hover:scale-125 ease-in-out duration-1000 h-64 w-full object-cover object-center" use:lazy={'/images/why/panda-5.png'} alt="A very cute cat." />
				</div>
			</div>

			<div class="prose prose-lg prose-slate px-4 mx-auto text-gray-500 mb-8">
				<h3>Where did you source all this data?</h3>
				<p>There are quite a few sources I've collated from, but these are the main ones:</p>
				<ol>
					<li>
						<a href="https://www.inaturalist.org/" title="iNaturalist">iNaturalist</a> is the best source of high-quality, community-driven <a href="https://creativecommons.org/" title="Creative Commons">Creative Commons</a> license photography. All images have been sourced from this site along with licensing and attribution data.
					</li>
					<li>
						<a href="https://www.aspca.org/" title="ASPCA">ASPCA</a> was used to initially prime the first dataset. This is also where I sourced most of the common name and symptom data.
					</li>
					<li>
						<a href="https://en.wikipedia.org/wiki/Plant" title="Wikipedia">Wikipedia</a> is the best source of scientific classification data out there.
					</li>
				</ol>
				<p>
					All of this disparate data was collated and munged together by several processing scripts written in <a href="https://go.dev/" title="Golang's main site.">Go</a> as
					<a href="https://magefile.org/" title="Magefiles' main site.">Magefiles</a>. It got me about 95% there, but it still needs quite a bit of handraulic finessing.
				</p>
				<h3>Where to go from here?</h3>
				<p>I think I've met most, if not all, of my requirements. The project still needs a lot of love and I plan on supporting listings for a variety of other pet species. I already have the data, just need to go through it with a fine-toothed comb.</p>
				<p>I'm having loads of fun at the moment learning new things. I hope I can keep doing this for a while longer. However, I do have other things planned for the future and will be using what I've learned here as a kind of launching pad.</p>
				<p>
					I sincerely hope you find <span class=" text-amber-400 mr-1 font-serif font-bold">Plant</span>🌱<span class="text-green-600 font-serif font-bold">Smart</span> as useful as I had fun making it.
				</p>
			</div>
		</div>
	</div>
</div>

<Seo title="Why?" canonical="https://plantsm.art/why" description="Why I made Plant Smart." />
