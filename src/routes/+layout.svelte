<script lang="ts">
	import '../app.postcss';
	import { Header, Footer } from '$components';
</script>

<div class="hidden text-yellow-700 bg-yellow-100 text-center py-2.5 text-sm border border-b-yellow-300">
	<p>Hi! 👋 😬 This project is still under heavy development and heaps of features are still being worked on and finessed.</p>
</div>

<Header />

<slot />

<Footer />
