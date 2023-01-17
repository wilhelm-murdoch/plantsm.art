<script lang="ts">
	import Logo from '$components/Logo/Logo.svelte';
</script>

<div class="px-6">
	<div class="relative mx-auto max-w-7xl">
		<nav class="py-4 flex text-center md:text-left">
			<div class="flex-1 space-x-4">
				<a href="/" class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500 font-medium">Home</a>

				<a href="/why" class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500 font-medium">Why?</a>

				<a href="/contribute" class="border-b border-dotted border-b-green-600 hover:border-b-green-800 text-green-600 hover:text-green-500 font-medium">Contribute</a>

				<a href="/api" class="border-b border-dotted border-b-green-600 text-green-600 hover:text-green-500 font-medium">API</a>
			</div>
			<div class="flex-2 space-x-2 hidden md:inline-block">
				<a href="https://twitter.com/wilhelm" title="Twitter" class="text-slate-300 hover:text-slate-400 inline-block align-middle border-r pr-4 border-r-slate-100">
					<span class="sr-only">Twitter</span>
					<svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
						<path d="M8.29 20.251c7.547 0 11.675-6.253 11.675-11.675 0-.178 0-.355-.012-.53A8.348 8.348 0 0022 5.92a8.19 8.19 0 01-2.357.646 4.118 4.118 0 001.804-2.27 8.224 8.224 0 01-2.605.996 4.107 4.107 0 00-6.993 3.743 11.65 11.65 0 01-8.457-4.287 4.106 4.106 0 001.27 5.477A4.072 4.072 0 012.8 9.713v.052a4.105 4.105 0 003.292 4.022 4.095 4.095 0 01-1.853.07 4.108 4.108 0 003.834 2.85A8.233 8.233 0 012 18.407a11.616 11.616 0 006.29 1.84" />
					</svg>
				</a>

				<a href="https://github.com/wilhelm-murdoch/plantsm.art" title="GitHub" class="text-slate-300 hover:text-slate-400 inline-block align-middle border-r pr-4 border-r-slate-100">
					<span class="sr-only">GitHub</span>
					<svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
						<path
							fill-rule="evenodd"
							d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z"
							clip-rule="evenodd"
						/>
					</svg>
				</a>

				<a href="https://ko-fi.com/devilmaycode" title="Ko-fi" class="text-slate-300 hover:text-slate-400 inline-block align-middle border-r pr-4 border-r-slate-100">
					<span class="sr-only">Ko-fi</span>
					<svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
						<path
							fill="currentColor"
							d="M23.881 8.948c-.773-4.085-4.859-4.593-4.859-4.593H.723c-.604 0-.679.798-.679.798s-.082 7.324-.022 11.822c.164 2.424 2.586 2.672 2.586 2.672s8.267-.023 11.966-.049c2.438-.426 2.683-2.566 2.658-3.734c4.352.24 7.422-2.831 6.649-6.916zm-11.062 3.511c-1.246 1.453-4.011 3.976-4.011 3.976s-.121.119-.31.023c-.076-.057-.108-.09-.108-.09c-.443-.441-3.368-3.049-4.034-3.954c-.709-.965-1.041-2.7-.091-3.71c.951-1.01 3.005-1.086 4.363.407c0 0 1.565-1.782 3.468-.963c1.904.82 1.832 3.011.723 4.311zm6.173.478c-.928.116-1.682.028-1.682.028V7.284h1.77s1.971.551 1.971 2.638c0 1.913-.985 2.667-2.059 3.015z"
						/>
					</svg>
				</a>

				<a href="/updates" title="Updates" class="text-red-600 hover:text-red-500 inline-block align-middle">
					<svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M10 2a6 6 0 00-6 6v3.586l-.707.707A1 1 0 004 14h12a1 1 0 00.707-1.707L16 11.586V8a6 6 0 00-6-6zM10 18a3 3 0 01-3-3h6a3 3 0 01-3 3z" /></svg>
				</a>
			</div>
		</nav>
	</div>
</div>

<div class="relative mx-auto max-w-7xl mb-8">
	<Logo classes={'p-2.5 text-center'} />
</div>
