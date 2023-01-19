<script lang="ts">
	import SvelteSeo from 'svelte-seo';
	import { ForkEvent, IssueCommentEvent, IssuesEvent, PublicEvent, PullRequestEvent, PullRequestReviewEvent, PushEvent, WatchEvent } from '$components';
	import { eventIcons } from '$components/GithubFeed/events';

	export let data: any;

	const getEventsByUser = (id: string) => {
		return data.events[id];
	};

	const getEventUsers = () => {
		return Object.keys(data.events);
	};

	const getEventCounts = () => {
		let counts: any = {};

		Object.keys(eventIcons).forEach((e) => {
			counts[e] = 0;
		});

		getEventUsers().forEach((u: string) => {
			getEventsByUser(u).events.forEach((e: any) => {
				counts[e.type]++;
			});
		});

		return counts;
	};

	let eventCounts = getEventCounts();

	const twitterMeta = {
		site: '@wilhelm',
		title: 'Plant Smart &middot; Updates',
		description: 'A free service that aims to provide a detailed listing of dangerous plants for your pets.',
		image: 'https://plantsm.art/images/og-cover.png',
		imageAlt: 'Plant Smart'
	};

	const openGraphMeta = {
		title: 'Plant Smart &middot; Updates',
		description: 'A free service that aims to provide a detailed listing of dangerous plants for your pets.',
		type: 'article',
		url: 'https://plantsm.art/updates',
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
			<h2 class="unstyled bg-gradient-to-tr from-green-700 to-emerald-900 p-2.5 text-2xl font-extralight text-white text-center mb-8">updates</h2>
			<div class="prose prose-lg prose-slate px-4 mx-auto text-gray-500 mb-8">
				<h2 class="block text-center text-3xl font-bold leading-8 tracking-tight text-gray-900 sm:text-4xl">See how industrious we've been!</h2>

				<p class="text-xl leading-8 text-gray-500 text-justify">
					Here's a list of the last 100 development events grouped by contributor. Feel free to check out our <a href="https://github.com/wilhelm-murdoch/plantsm.art" title="The Plant Smart repository on GitHub.">GitHub repository</a> if you want to see a more exhaustive view.
				</p>

				<div class="flex mb-8 justify-center space-x-2 sm:space-x-8">
					{#each Object.keys(eventCounts) as eventType}
						{#if eventCounts[eventType] > 0}
							<div class="relative inline-block rounded-full p-1 text-{eventIcons[eventType][0]}-800 bg-{eventIcons[eventType][0]}-100" aria-hidden="true">
								<svg class="w-8 h-8" fill="none" stroke-width="1" viewBox="0 0 24 24" stroke="currentColor" xmlns="http://www.w3.org/2000/svg">{@html eventIcons[eventType][1]}</svg><span class="not-prose absolute -top-2 -right-2 block rounded-md text-center py-0.5 px-1 bg-green-500 ring-2 ring-white text-white text-xs">{eventCounts[eventType]}</span>
							</div>
						{/if}
					{/each}
				</div>

				{#each getEventUsers() as user}
					{@const grouped = getEventsByUser(user)}

					<article class="not-prose pb-6">
						<div class="xl:flex">
							<div class="grow">
								<header class="border-b mb-6">
									<div class="flex flex-nowrap items-center space-x-2 mb-4">
										<div class="flex shrink-0 -space-x-3 -ml-px">
											<a class="block" href={grouped.url} title="{grouped.name}'s Github profile.">
												<img class="rounded-full border-2 border-slate-50 box-content h-12 w-12 border-solid" src={grouped.avatar} alt={user} />
											</a>
										</div>
										<h3 class="text-lg">
											{user}
											<span class="text-slate-300 text-sm block">{grouped.events.length} events</span>
										</h3>
									</div>
								</header>
								<ul class="ml-2">
									{#each grouped.events as event, i}
										{#if event.type == 'PushEvent'}
											<PushEvent index={i} numEntries={grouped.events.length} {event} />
										{:else if event.type == 'ForkEvent'}
											<ForkEvent index={i} numEntries={grouped.events.length} {event} />
										{:else if event.type == 'IssueCommentEvent'}
											<IssueCommentEvent index={i} numEntries={grouped.events.length} {event} />
										{:else if event.type == 'IssuesEvent'}
											<IssuesEvent index={i} numEntries={grouped.events.length} {event} />
										{:else if event.type == 'PublicEvent'}
											<PublicEvent index={i} numEntries={grouped.events.length} {event} />
										{:else if event.type == 'PullRequestEvent'}
											<PullRequestEvent index={i} numEntries={grouped.events.length} {event} />
										{:else if event.type == 'WatchEvent'}
											<WatchEvent index={i} numEntries={grouped.events.length} {event} />
										{:else if event.type == 'PullRequestReviewEvent'}
											<PullRequestReviewEvent index={i} numEntries={grouped.events.length} {event} />
										{/if}
									{/each}
								</ul>
							</div>
						</div>
					</article>
				{/each}
			</div>
		</div>
	</div>
</div>

<SvelteSeo title="Plant Smart &middot; Updates" description="A free service that aims to provide a detailed listing of dangerous plants for your pets." canonical="https://plantsm.art/updates" twitter={twitterMeta} openGraph={openGraphMeta} />
