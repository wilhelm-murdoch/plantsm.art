<script lang="ts">
	import Time from 'svelte-time';

	export let index: number = 0;
	export let numEntries: number = 0;
	export let event: any;
</script>

<li class="relative py-2">
	<div class="flex items-center mb-1">
		<div
			class:hidden={index == numEntries - 1}
			class="absolute left-0 h-full w-0.5 bg-slate-200 self-start ml-4 -translate-x-1/2 translate-y-3"
			aria-hidden="true"
		/>
		<div class="absolute rounded-full p-1 text-orange-800 bg-orange-100" aria-hidden="true">
			<svg
				class="w-6 h-6"
				fill="none"
				stroke-width="1"
				viewBox="0 0 24 24"
				stroke="currentColor"
				xmlns="http://www.w3.org/2000/svg"
				><path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z"
				/></svg
			>
		</div>
		<h3 class="text-sm font-bold text-slate-800 pl-12">
			{event.payload.action.charAt(0).toUpperCase() + event.payload.action.slice(1)} comment in
			<a
				href={event.payload.issue.html_url}
				title="View {event.payload.issue.title} on Github."
				class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500 font-normal"
				>{event.payload.issue.title}</a
			>
			about
			<span class="font-normal text-slate-400"><Time relative timestamp={event.created_at} /></span>
		</h3>
	</div>
	<div class="pl-12 prose text-sm pb-5">{event.payload.issue.body}</div>
</li>
