<script lang="ts">
	import Time from 'svelte-time';
	import { eventIcons } from './events';

	export let index: number = 0;
	export let numEntries: number = 0;
	export let event: any;
</script>

<li class="relative py-2">
	<div class="flex items-center mb-1 pb-5">
		<div
			class:hidden={index == numEntries - 1}
			class="absolute left-0 h-full w-0.5 bg-slate-200 self-start ml-4 -translate-x-1/2 translate-y-3"
			aria-hidden="true"
		/>
		<div
			class="absolute rounded-full p-1 text-{eventIcons.IssuesEvent[0]}-800 bg-{eventIcons
				.IssuesEvent[0]}-100"
			aria-hidden="true"
		>
			<svg
				class="w-6 h-6"
				fill="none"
				stroke-width="1"
				viewBox="0 0 24 24"
				stroke="currentColor"
				xmlns="http://www.w3.org/2000/svg">{@html eventIcons.IssuesEvent[1]}</svg
			>
		</div>
		<h3 class="text-sm font-bold text-slate-800 pl-12">
			{event.payload.action.charAt(0).toUpperCase() + event.payload.action.slice(1)} issue
			<a
				href={event.payload.issue.html_url}
				title="View {event.payload.issue.title} on Github."
				class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500 font-normal"
				>{event.payload.issue.title}</a
			>
			{#if event.payload.issue.labels.length}
				labeled
				{#each event.payload.issue.labels as label}
					<span
						class="inline-flex font-normal text-xs items-center rounded-md px-2 text-blue-800 border"
						style="border-color: #{label.color}80; color: #{label.color}; background-color: #{label.color}25"
						>{label.name}
					</span>
				{/each}
			{/if} about
			<span class="font-normal text-slate-400"><Time relative timestamp={event.created_at} /></span>
		</h3>
	</div>
</li>
