<script lang="ts">
	import '../app.postcss';
	import { Logo, Footer } from '../components';
</script>

<div class="bg-yellow-50 p-4 text-sm text-center text-yellow-700">
	<p>
		This project is currently under active development. Heaps of planned features have not yet been
		implemented and data may not be complete. Be sure to check back for regular updates!
	</p>
</div>

<div class="relative mx-auto max-w-7xl mb-8">
	<Logo classes={'p-2.5 text-center'} />
</div>
<div
	class="relative bg-gray-50 pb-5 mb-8 sm:px-6 sm:pt-4 lg:px-8 lg:pt-4 lg:pb-4 border-b border-t"
>
	<div class="relative mx-auto max-w-7xl">
		<slot />
	</div>
</div>
<Footer />
