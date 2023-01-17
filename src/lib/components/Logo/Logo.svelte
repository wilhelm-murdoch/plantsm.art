<script lang="ts">
	import { randomIcon } from './icons';
	import { browser } from '$app/environment';

	export let classes = '';

	const classesBase = 'font-serif font-bold logo';

	$: compositeClasses = `${classesBase} ${classes}`;
</script>

<div class="h-16">
	<h1 class={compositeClasses}>
		{#if browser}
			<a href="/" title="Back to the main page.">
				<span class="text-amber-400 text-4xl md:text-6xl mr-1">Plant</span>
				<svg
					class="inline-block h-12 w-12 md:h-14 md:w-14 align-baseline jello-animation"
					enable-background="new 0 0 512 512"
					height="512"
					viewBox="0 0 512 512"
					width="512"
					xmlns="http://www.w3.org/2000/svg">{@html randomIcon.svg}</svg
				>

				<span class="text-green-600 italic text-4xl md:text-6xl ml-1">Smart</span></a
			>
		{/if}
	</h1>
</div>

<style>
	@keyframes fade-in {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}

	.logo {
		animation: fade-in 0.2s;
	}

	.jello-animation {
		transform-origin: 50% 100%;
		-webkit-animation: jello-animation 0.9s both;
		animation: jello-animation 0.9s both;
	}

	@-webkit-keyframes jello-animation {
		0% {
			-webkit-transform: scale3d(1, 1, 1);
			transform: scale3d(1, 1, 1);
		}

		30% {
			-webkit-transform: scale3d(1.25, 0.75, 1);
			transform: scale3d(1.25, 0.75, 1);
		}

		40% {
			-webkit-transform: scale3d(0.75, 1.25, 1);
			transform: scale3d(0.75, 1.25, 1);
		}

		50% {
			-webkit-transform: scale3d(1.15, 0.85, 1);
			transform: scale3d(1.15, 0.85, 1);
		}

		65% {
			-webkit-transform: scale3d(0.95, 1.05, 1);
			transform: scale3d(0.95, 1.05, 1);
		}

		75% {
			-webkit-transform: scale3d(1.05, 0.95, 1);
			transform: scale3d(1.05, 0.95, 1);
		}

		100% {
			-webkit-transform: scale3d(1, 1, 1);
			transform: scale3d(1, 1, 1);
		}
	}

	@keyframes jello-animation {
		0% {
			-webkit-transform: scale3d(1, 1, 1);
			transform: scale3d(1, 1, 1);
		}

		30% {
			-webkit-transform: scale3d(1.25, 0.75, 1);
			transform: scale3d(1.25, 0.75, 1);
		}

		40% {
			-webkit-transform: scale3d(0.75, 1.25, 1);
			transform: scale3d(0.75, 1.25, 1);
		}

		50% {
			-webkit-transform: scale3d(1.15, 0.85, 1);
			transform: scale3d(1.15, 0.85, 1);
		}

		65% {
			-webkit-transform: scale3d(0.95, 1.05, 1);
			transform: scale3d(0.95, 1.05, 1);
		}

		75% {
			-webkit-transform: scale3d(1.05, 0.95, 1);
			transform: scale3d(1.05, 0.95, 1);
		}

		100% {
			-webkit-transform: scale3d(1, 1, 1);
			transform: scale3d(1, 1, 1);
		}
	}
</style>
