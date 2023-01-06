<script lang="ts">
	import Time from 'svelte-time';

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
		<div class="absolute rounded-full p-1 text-indigo-800 bg-indigo-100" aria-hidden="true">
			<svg
				class="w-6 h-6"
				fill="none"
				stroke-width="1"
				viewBox="0 0 24 24"
				stroke="currentColor"
				xmlns="http://www.w3.org/2000/svg"
				><path
					d="M18 21a2 2 0 100-4 2 2 0 000 4zM6 7a2 2 0 100-4 2 2 0 000 4zM6 21a2 2 0 100-4 2 2 0 000 4zM6 7v10M18 17V7s0-2-2-2h-3"
					stroke-linecap="round"
					stroke-linejoin="round"
				/><path d="M15 7.5L12.5 5 15 2.5" stroke-linecap="round" stroke-linejoin="round" /></svg
			>
		</div>
		<h3 class="text-sm font-bold text-slate-800 pl-12">
			{event.payload.action.charAt(0).toUpperCase() + event.payload.action.slice(1)} the pull request
			<a
				href={event.payload.pull_request.html_url}
				title="View {event.payload.pull_request.title} on Github."
				class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500 font-normal"
				>{event.payload.pull_request.title}</a
			>
			about
			<span class="font-normal text-slate-400"><Time relative timestamp={event.created_at} /></span>
		</h3>
	</div>
</li>
