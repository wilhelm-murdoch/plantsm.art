<div class="mx-auto max-w-7xl overflow-hidden py-2.5 sm:px-6 lg:px-8">
	<p class="text-center text-sm text-slate-300">
		<span class="text-amber-400 font-serif">Plant</span>🪴<span
			class="text-green-600 italic font-serif">Smart</span
		>
		is made with
		<svg
			class="w-6 h-6 text-red-600 inline-block heartbeat-animation"
			fill="currentColor"
			viewBox="0 0 20 20"
			xmlns="http://www.w3.org/2000/svg"
			><path
				fill-rule="evenodd"
				d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z"
				clip-rule="evenodd"
			/></svg
		>
		by
		<a
			href="https://wilhelm.codes"
			title="It me."
			class="font-semibold border-b hover:text-slate-400">wilhelm</a
		>
		<img
			class="inline-block h-6 w-6 rounded-full border-2 border-slate-300 hover:border-slate-400 jello-animation"
			src="https://s.gravatar.com/avatar/7cc6bee5975704b299bd7355e8ba9356?s=40"
			alt="It me."
		/>
	</p>
	<!-- svelte-ignore missing-declaration -->
	<p class="text-center text-sm text-slate-200 mt-2.5 font-mono">{__BUILD_NUMBER__}</p>
</div>

<style>
	.heartbeat-animation:hover {
		-webkit-animation: heartbeat-animation 1.5s ease-in-out 2 both;
		animation: heartbeat-animation 1.5s ease-in-out 2 both;
	}

	@-webkit-keyframes heartbeat-animation {
		from {
			-webkit-transform: scale(1);
			transform: scale(1);
			-webkit-transform-origin: center center;
			transform-origin: center center;
			-webkit-animation-timing-function: ease-out;
			animation-timing-function: ease-out;
		}
		10% {
			-webkit-transform: scale(0.91);
			transform: scale(0.91);
			-webkit-animation-timing-function: ease-in;
			animation-timing-function: ease-in;
		}
		17% {
			-webkit-transform: scale(0.98);
			transform: scale(0.98);
			-webkit-animation-timing-function: ease-out;
			animation-timing-function: ease-out;
		}
		33% {
			-webkit-transform: scale(0.87);
			transform: scale(0.87);
			-webkit-animation-timing-function: ease-in;
			animation-timing-function: ease-in;
		}
		45% {
			-webkit-transform: scale(1);
			transform: scale(1);
			-webkit-animation-timing-function: ease-out;
			animation-timing-function: ease-out;
		}
	}

	@keyframes heartbeat-animation {
		from {
			-webkit-transform: scale(1);
			transform: scale(1);
			-webkit-transform-origin: center center;
			transform-origin: center center;
			-webkit-animation-timing-function: ease-out;
			animation-timing-function: ease-out;
		}
		10% {
			-webkit-transform: scale(0.91);
			transform: scale(0.91);
			-webkit-animation-timing-function: ease-in;
			animation-timing-function: ease-in;
		}
		17% {
			-webkit-transform: scale(0.98);
			transform: scale(0.98);
			-webkit-animation-timing-function: ease-out;
			animation-timing-function: ease-out;
		}
		33% {
			-webkit-transform: scale(0.87);
			transform: scale(0.87);
			-webkit-animation-timing-function: ease-in;
			animation-timing-function: ease-in;
		}
		45% {
			-webkit-transform: scale(1);
			transform: scale(1);
			-webkit-animation-timing-function: ease-out;
			animation-timing-function: ease-out;
		}
	}

	.jello-animation:hover {
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
