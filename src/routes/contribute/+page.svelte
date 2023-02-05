<script lang="ts">
	import { Seo } from '$components';

	export let data: any;
</script>

<div class="relative bg-gray-50 mb-4 p-4 border-y">
	<div class="relative mx-auto max-w-7xl">
		<div class="overflow-hidden bg-white shadow-lg mb-4 pb-8 rounded-lg">
			<h2 class="unstyled bg-gradient-to-tr from-green-700 to-emerald-900 p-2.5 text-2xl font-extralight text-white text-center mb-8">contribute</h2>
			<div class="prose prose-lg prose-slate px-4 mx-auto text-gray-500">
				<h2 class="block text-center text-3xl font-bold leading-8 tracking-tight text-gray-900 sm:text-4xl">Learn how you can help out.</h2>

				<p class="text-xl leading-8 text-gray-500 text-justify">
					<span class=" text-amber-400 mr-1 font-serif font-bold">Plant</span>🌱<span class="text-green-600 italic font-serif font-bold">Smart</span> is a free service built upon open source software. We welcome constructive feedback and participation from the public.
				</p>

				<ul class="text-gray-400">
					<li>Found some new plants that aren’t listed yet?</li>
					<li>Found a <strike>speeling missed steak</strike> spelling mistake?</li>
					<li>Found something that’s not working properly?</li>
					<li>Found some better images for an existing record?</li>
					<li>Got a suggestion for a new feature?</li>
					<li>Felt the deep and undeniable urge to buy me a ☕️?</li>
				</ul>

				<p>
					All participation and contributions are done using <a href="https://github.com" title="Github">Github</a>, so if you’d like to help out, you’re going to need to
					<a href="https://github.com/signup" title="Register a Github account.">register an account</a> if you don’t already have one.
				</p>

				<h3>Submit an issue</h3>

				<p>
					This is probably the most pragmatic option for most people. <a href="https://github.com/login" title="Log into Github.">Log into Github</a>
					and
					<a href="https://github.com/wilhelm-murdoch/plantsm.art/issues/new" title="Submit a new issue.">submit a new issue</a>. Be sure to add a detailed title, description and — most importantly — ensure you’ve selected an appropriate tag to associate with the issue. This allows me to quickly sort though and prioritise other entries.
				</p>

				<h3>Fork the project</h3>

				<p>
					If you’d like to get your hands dirty with the source code, or update the plant database yourself with new records and images, feel free to <a href="https://github.com/wilhelm-murdoch/plantsm.art/fork" title="Fork the repository.">fork the repository</a>
					yourself and
					<a href="https://github.com/wilhelm-murdoch/plantsm.art/compare" title="Submit a pull request.">submit a pull request</a> when you’re finished.
				</p>

				<p>
					All changes are peer-reviewed before merging into the <code>main</code> branch and deploying to production. If your change is approved, you’ll automatically show up in the
					<a href="#contributors" title="Contributors">contributors' list</a> below. 🙂
				</p>

				<h3>Shoot me a message</h3>

				<p>
					If you’d rather not go through the hassle of dealing with Github, feel free to fire off a message to me on <a href="https://twitter.com/wilhelm" title="Wilhelm's Twitter account.">Twitter</a> and I’ll try to get back to you as soon as I’m able.
				</p>

				<h3>Buy me a ☕️!</h3>

				<p>
					By design, operational costs are pretty low for this service. At the moment, I only pay to renew the domain name. But, I do spend a lot of my free time maintaining the site. If you value this service and want to help out in some small way, feel free to send me the gift of caffeine via my <a href="https://ko-fi.com/devilmaycode" title="My Ko-fi account.">Ko-fi account</a>. Any amount is greatly appreciated!
				</p>

				<h3 id="contributors">Contributors</h3>
				<p>The following is a list of people who have contributed on Github.</p>
				<div class="flex space-x-5">
					{#each data.contributors as contributor}
						<a class="not-prose relative inline-block" title="{contributor.name}'s Github profile." href={contributor.url}>
							<img class="h-16 w-16 rounded-full border-solid border-2 border-black/5" src={contributor.avatar} alt="{contributor.name}'s profile image." />
							<span class="not-prose absolute top-0 right-0 block rounded-md text-center py-0.5 px-1 bg-green-500 ring-2 ring-white text-white text-xs">{contributor.contributions}</span>
						</a>
					{/each}
				</div>
			</div>
		</div>
	</div>
</div>

<Seo title="How to contribute" canonical="https://plantsm.art/contribute" />
