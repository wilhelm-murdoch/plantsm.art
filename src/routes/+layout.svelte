<script lang="ts">
	import '../app.postcss';
	import { Header, Footer } from '$components';
</script>

<div class="bg-yellow-50 p-4 text-sm text-center text-yellow-700 mb-8">
	<p>
		This project is currently under active development. Heaps of planned features have not yet been
		implemented and data may not be complete. Be sure to check back for regular updates!
	</p>
</div>

<Header />

<slot />

<Footer />
