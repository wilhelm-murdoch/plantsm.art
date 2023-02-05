<script lang="ts">
	import SvelteSeo from 'svelte-seo';

	export let title = 'Helping plants and pets peacefully coexist.';
	export let description = 'A free service that aims to provide a detailed listing of dangerous plants for your pets.';
	export let canonical = 'https://plantsm.art';
	export let image = 'https://plantsm.art/images/og-cover.png';
	export let attribution = 'Plant Smart';

	const twitter = {
		site: '@wilhelm',
		title: 'Plant Smart &middot; {title}',
		description: description,
		image: image,
		imageAlt: attribution
	};

	const openGraph = {
		title: 'Plant Smart &middot; {title}',
		description: description,
		type: 'article',
		url: 'https://plantsm.art',
		images: [
			{
				url: image,
				alt: attribution
			}
		]
	};

	const ldJson = {
		'@context': 'https://schema.org/',
		'@type': 'Article',
		headline: { title },
		author: {
			'@type': 'Person',
			name: 'Wilhelm Murdoch',
			url: 'https://wilhelm.codes/'
		},
		image: [{ image }],
		description: { description }
	};
</script>

<svelte:head>
	{@html '<script type="application/ld+json">' + JSON.stringify(ldJson) + '</script>'}
</svelte:head>

<SvelteSeo title="Plant Smart &middot; {title}" {description} {canonical} {twitter} {openGraph} />
