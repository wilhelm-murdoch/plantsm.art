<script lang="ts">
	import Time from 'svelte-time';
	import { eventIcons } from './events';

	export let index: number = 0;
	export let numEntries: number = 0;
	export let event: any;
</script>

<li class="relative py-2">
	<div class="flex items-center mb-1">
		<div
			class:hidden={index == numEntries - 1}
			class="absolute left-0 h-full w-0.5 bg-slate-100 self-start ml-4 -translate-x-1/2 translate-y-3"
			aria-hidden="true"
		/>
		<div
			class="absolute rounded-full p-1 text-{eventIcons.PushEvent[0]}-800 bg-{eventIcons
				.PushEvent[0]}-100"
			aria-hidden="true"
		>
			<svg
				class="w-6 h-6"
				fill="none"
				stroke-width="1"
				viewBox="0 0 24 24"
				stroke="currentColor"
				xmlns="http://www.w3.org/2000/svg">{@html eventIcons.PushEvent[1]}</svg
			>
		</div>
		<h3 class="text-sm font-bold text-slate-800 pl-12">
			Made <span class="text-slate-400 font-normal">{event.payload.size}</span>
			{#if event.payload.size == 1}change{:else}changes{/if} about
			<span class="text-slate-400 font-normal"><Time relative timestamp={event.created_at} /></span>
		</h3>
	</div>
	<div class="pb-5">
		{#each event.payload.commits as commit}
			<p class="ml-12 border-b border-b-slate-100">
				<span class="text-xs font-mono"
					><a
						href={commit.url.replace('api.', '').replace('repos/', '')}
						title="sha: {commit.sha}"
						class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500"
						>{commit.sha.slice(0, 8)}</a
					><svg
						class="w-4 h-4 inline-block ml-1"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
						xmlns="http://www.w3.org/2000/svg"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="1.5"
							d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
						/></svg
					></span
				>
				&middot; <span class="text-xs">{commit.message}</span>
			</p>
		{/each}
	</div>
</li>
