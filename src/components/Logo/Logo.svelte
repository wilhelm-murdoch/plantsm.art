<script lang="ts">
	import { onMount } from 'svelte';

	export let classes = '';

	const classesBase = 'font-serif font-bold';

	$: compositeClasses = `${classesBase} ${classes}`;

	const icons = [
		'aglaonema',
		'anthurium',
		'begonia',
		'cactus-1',
		'calathea',
		'climbing-plant',
		'dragon-tree',
		'english-ivy',
		'ficus',
		'guzmania',
		'hoya',
		'indoor-plants-1',
		'kalanchoe',
		'lavender',
		'lithops',
		'monstera',
		'orchid',
		'paddle-plant',
		'philodendron',
		'sansevieria-bacularis',
		'spider-plant',
		'string-of-beads',
		'succulent-2',
		'violet',
		'yucca',
		'aloe-vera',
		'bamboo',
		'bonsai',
		'cactus-2',
		'cast-iron',
		'coral-cactus',
		'elephant-ear',
		'fern',
		'gerbera',
		'hanging-pot',
		'hyacinth',
		'indoor-plants-2',
		'lady-palm',
		'lilly',
		'mimosa',
		'moon-cactus',
		'oxalis-triangularis',
		'peperomia-watermelon',
		'pilea',
		'snake-plant',
		'staghorn-fern',
		'succulent-1',
		'venus-flytrap',
		'weeping-fig',
		'zz-plant'
	];

	let randomIcon: string;

	onMount(() => {
		randomIcon = icons[Math.floor(Math.random() * icons.length)];
	});
</script>

<h1 class={compositeClasses}>
	<a href="/" title="Back to the main page.">
		<span class=" text-amber-400 text-6xl mr-1">Plant</span>
		{#if randomIcon}<img
				src="/icons/plants/{randomIcon}.svg"
				alt="An image of {randomIcon}."
				class="inline-block h-14 w-14 align-baseline jello-animation"
			/>
		{:else}
			<span class="inline-block animate-pulse align-baseline h-14 w-14 bg-slate-50 rounded-lg" />
		{/if}
		<span class="text-green-600 italic text-6xl ml-1">Smart</span></a
	>
</h1>

<style>
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
