<script lang="ts">
	import { getByAnimal, normalizeAnimal } from '$utils/animals';

	export let plant: any = {};
</script>

<div class="flex flex-col overflow-hidden rounded-lg shadow-lg mb-4">
	<div class="flex-shrink-0">
		<h3 class="unstyled font-sans mb-3 p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500">Affects</h3>
		<p class="px-2.5 pb-3 text-md font-medium space-x-1 space-y-1">
			{#each plant.animals as animal}
				<span class="inline-flex items-center rounded-md bg-{getByAnimal(animal).background} px-2.5 py-1 text-{getByAnimal(animal).foreground}">{getByAnimal(animal).emoji} {normalizeAnimal(animal)}</span>
			{/each}
		</p>
<!-- 
		<span class="inline-flex">
			<svg id="emoji" class="h-8 w-8" viewBox="0 0 72 72" xmlns="http://www.w3.org/2000/svg">
				<g id="color">
					<path fill="#f4aa41" d="m24.473,15.1583l-5.0799,1.9352-7.2963,7.901-4.1377,10.5005,1.291,5.3405c1.2554,3.7911,3.3357,6.4338,7.0626,9.2506l2.6874-2.5839s3.8218,7.7098,10.7384,8.9598c0,0,10.2616,1.936,15.5949-.8765,1.4899-.7857,2.5141-1.8291,3.2921-2.5939,2.0702-2.0351,3.033-3.5201,4.5413-5.2395h0s1.6701,1.8077,1.6701,1.8077l1.838-.0557,5.0169-7.2292,2.0032-5.0703-.0215-4.255-2.1735-5.6141-4.8333-7.4167s-2.6368-4.2558-8.1667-3.9167c0,0-6.5-4.8333-11.8333-4.0833s-3.6104-.6772-12.1937,3.2395Z"/>
					<polygon fill="#ea5a47" points="36 46.7324 32.9167 49.1491 30.4167 49.1491 30.9413 53.0386 31.362 56.0539 32.1667 58.3991 35 59.8991 39.5833 59.3158 40.4429 57.1116 41.1441 52.9335 41.9167 49.3158 39.9167 49.5658 36 46.7324"/>
					<polygon fill="#3f3f3f" points="32.5 36.9188 30.9167 40.6688 33.0833 41.9188 34.3333 42.4188 38.6667 42.5855 41.5833 40.3355 39.8333 37.0855 32.5 36.9188"/>
				  </g>
				  <g id="line">
					<path d="m29.5059,30.1088s-1.8051,1.2424-2.7484.6679c-.9434-.5745-1.2424-1.8051-.6679-2.7484s1.805-1.2424,2.7484-.6679.6679,2.7484.6679,2.7484Z"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m33.1089,37.006h6.1457c.4011,0,.7634.2397.9203.6089l1.1579,2.7245-2.1792,1.1456c-.6156.3236-1.3654-.0645-1.4567-.754"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m34.7606,40.763c-.1132.6268-.7757.9895-1.3647.7471l-2.3132-.952,1.0899-2.9035c.1465-.3901.5195-.6486.9362-.6486"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m30.4364,50.0268s-.7187,8.7934,3.0072,9.9375c2.6459.8125,5.1497.5324,6.0625-.25.875-.75,2.6323-4.4741,1.8267-9.6875"/>
					<path d="m44.2636,30.1088s1.805,1.2424,2.7484.6679,1.2424-1.8051.6679-2.7484c-.5745-.9434-1.805-1.2424-2.7484-.6679s-.6679,2.7484-.6679,2.7484Z"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m25.6245,42.8393c-.475,3.6024,2.2343,5.7505,4.2847,6.8414,1.1968.6367,2.6508.5182,3.7176-.3181l2.581-2.0233,2.581,2.0233c1.0669.8363,2.5209.9548,3.7176.3181,2.0504-1.0909,4.7597-3.239,4.2847-6.8414"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19.9509,28.3572c-2.3166,5.1597-.5084,13.0249.119,15.3759.122.4571.0755.9355-.1271,1.3631l-1.9874,4.1937c-.623,1.3146-2.3934,1.5533-3.331.4409-3.1921-3.7871-8.5584-11.3899-6.5486-16.686,7.0625-18.6104,15.8677-18.1429,15.8677-18.1429,2.8453-1.9336,13.1042-6.9375,24.8125.875,0,0,8.6323-1.7175,14.9375,16.9375,1.8036,5.3362-3.4297,12.8668-6.5506,16.6442-.9312,1.127-2.7162.8939-3.3423-.4272l-1.9741-4.1656c-.2026-.4275-.2491-.906-.1271-1.3631.6275-2.3509,2.4356-10.2161.119-15.3759"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m52.6309,46.4628s-3.0781,6.7216-7.8049,8.2712"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19.437,46.969s3.0781,6.0823,7.8049,7.632"/>
					<line x1="36.2078" x2="36.2078" y1="47.3393" y2="44.3093" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>
				  </g>
			</svg>
		  </span>

		  <span class="inline-flex">
			<svg id="emoji" class="h-8 w-8" viewBox="0 0 72 72" xmlns="http://www.w3.org/2000/svg">
				<g id="color">
					<path fill="#f4aa41" d="m24.473,15.1583l-5.0799,1.9352-7.2963,7.901-4.1377,10.5005,1.291,5.3405c1.2554,3.7911,3.3357,6.4338,7.0626,9.2506l2.6874-2.5839s3.8218,7.7098,10.7384,8.9598c0,0,10.2616,1.936,15.5949-.8765,1.4899-.7857,2.5141-1.8291,3.2921-2.5939,2.0702-2.0351,3.033-3.5201,4.5413-5.2395h0s1.6701,1.8077,1.6701,1.8077l1.838-.0557,5.0169-7.2292,2.0032-5.0703-.0215-4.255-2.1735-5.6141-4.8333-7.4167s-2.6368-4.2558-8.1667-3.9167c0,0-6.5-4.8333-11.8333-4.0833s-3.6104-.6772-12.1937,3.2395Z"/>
					<polygon fill="#ea5a47" points="36 46.7324 32.9167 49.1491 30.4167 49.1491 30.9413 53.0386 31.362 56.0539 32.1667 58.3991 35 59.8991 39.5833 59.3158 40.4429 57.1116 41.1441 52.9335 41.9167 49.3158 39.9167 49.5658 36 46.7324"/>
					<polygon fill="#3f3f3f" points="32.5 36.9188 30.9167 40.6688 33.0833 41.9188 34.3333 42.4188 38.6667 42.5855 41.5833 40.3355 39.8333 37.0855 32.5 36.9188"/>
				  </g>
				  <g id="line">
					<path d="m29.5059,30.1088s-1.8051,1.2424-2.7484.6679c-.9434-.5745-1.2424-1.8051-.6679-2.7484s1.805-1.2424,2.7484-.6679.6679,2.7484.6679,2.7484Z"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m33.1089,37.006h6.1457c.4011,0,.7634.2397.9203.6089l1.1579,2.7245-2.1792,1.1456c-.6156.3236-1.3654-.0645-1.4567-.754"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m34.7606,40.763c-.1132.6268-.7757.9895-1.3647.7471l-2.3132-.952,1.0899-2.9035c.1465-.3901.5195-.6486.9362-.6486"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m30.4364,50.0268s-.7187,8.7934,3.0072,9.9375c2.6459.8125,5.1497.5324,6.0625-.25.875-.75,2.6323-4.4741,1.8267-9.6875"/>
					<path d="m44.2636,30.1088s1.805,1.2424,2.7484.6679,1.2424-1.8051.6679-2.7484c-.5745-.9434-1.805-1.2424-2.7484-.6679s-.6679,2.7484-.6679,2.7484Z"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m25.6245,42.8393c-.475,3.6024,2.2343,5.7505,4.2847,6.8414,1.1968.6367,2.6508.5182,3.7176-.3181l2.581-2.0233,2.581,2.0233c1.0669.8363,2.5209.9548,3.7176.3181,2.0504-1.0909,4.7597-3.239,4.2847-6.8414"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19.9509,28.3572c-2.3166,5.1597-.5084,13.0249.119,15.3759.122.4571.0755.9355-.1271,1.3631l-1.9874,4.1937c-.623,1.3146-2.3934,1.5533-3.331.4409-3.1921-3.7871-8.5584-11.3899-6.5486-16.686,7.0625-18.6104,15.8677-18.1429,15.8677-18.1429,2.8453-1.9336,13.1042-6.9375,24.8125.875,0,0,8.6323-1.7175,14.9375,16.9375,1.8036,5.3362-3.4297,12.8668-6.5506,16.6442-.9312,1.127-2.7162.8939-3.3423-.4272l-1.9741-4.1656c-.2026-.4275-.2491-.906-.1271-1.3631.6275-2.3509,2.4356-10.2161.119-15.3759"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m52.6309,46.4628s-3.0781,6.7216-7.8049,8.2712"/>
					<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19.437,46.969s3.0781,6.0823,7.8049,7.632"/>
					<line x1="36.2078" x2="36.2078" y1="47.3393" y2="44.3093" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>
				  </g>
			</svg>
		  </span>

		  <span class="inline-flex">
			<svg id="emoji" class="h-8 w-8" viewBox="0 0 72 72" xmlns="http://www.w3.org/2000/svg">
				<g id="color">
				<path fill="#f4aa41" d="m58.2673,11.3469s-10.4076,2.3754-15.5743,6.7088c0,0-9-2.5-13.8333.1667,0,0-9.6549-6.7318-15.6549-6.7318,0,0-5.0326,3.75.3216,21.0651,0,0-2.6667,10.6667,1.6667,16.3333.7823,1.023,1.6026,1.9862,2.4217,2.8779,3.4268,3.7306,7.5912,6.7046,12.1937,8.8205l1.696.7797c1.5277.7023,3.1777,1.1,4.8576,1.1707h0c.9304.0392,1.8573-.1359,2.7093-.5118l4.5429-2.0042c3.8082-1.6801,7.2734-4.0872,10.0486-7.1894,1.1585-1.295,2.2135-2.71,2.8635-4.11,4.4736-10.6191,1.5314-16.2624,1.5314-16.2624l1.2356-7.1292c.8094-3.1482.8268-6.4477.0506-9.6043l-1.0768-4.3796Z"/>
				<path fill="#fff" d="m30.8377,47.3355s-7.3487,2.8338-1.0987,9.3338c0,0-1.6971,4.2984,3.5271,4.6285.6823.0431,2.7339.0635,2.7339.0635l1.5797.0367c.4833.0112.9656-.0228,1.4424-.1026,1.8709-.3132,3.9279-.7821,3.181-4.5878,0,0,7.5513-6.3722-1.3654-9.3722l-4.875,2-5.125-1.9999Z"/>
				</g>
				<g id="line">
				<ellipse cx="45.0854" cy="38.1033" rx="1.6461" ry="2.8119"/>
				<ellipse cx="26.8427" cy="38.1033" rx="1.6461" ry="2.8119"/>
				<polyline fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" points="31.9328 47.2287 36.037 50.0204 39.8495 47.2287"/>
				<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m36.037,50.0204v4.2708s-1.1042,3.6875-5.5417,2.875"/>
				<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15.8717,48.4759c-4.8928-7.2535-2.0014-15.8722-2.0014-15.8722,0,0-5.25-14.875-.4375-21.25,0,0,9.1875,1.5,15.6875,7.375,4.5946-1.9379,9.1575-2.0128,13.6875-.1437,6.5-5.875,15.6875-7.375,15.6875-7.375,4.8125,6.375-.4375,21.25-.4375,21.25,0,0,2.8914,8.6187-2.0014,15.8722"/>
				<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.7453,15.1037s12.8125,6.1875,10.0625,11.8125"/>
				<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m24.8491,50.8753s-9.3615-.458-13.6525,7.5243"/>
				<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m35.8911,49.8767v4.2708s1.1042,3.6875,5.5417,2.875"/>
				<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m57.1828,14.96s-12.8125,6.1875-10.0625,11.8125"/>
				<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m47.2048,54.6836s8.2116,2.2454,8.6795,11.2958"/>
				<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m48.079,50.7316s9.3615-.458,13.6525,7.5243"/>
				<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m31.3859,60.7598c3.88,1.6845,5.6481,1.8093,9.3021,0"/>
				<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m25.4446,54.6836s-8.2116,2.2454-8.6795,11.2958"/>
				</g>
			</svg>
		  </span> -->
		  


		<h3 class="unstyled font-sans mb-3 p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500">Common Names</h3>
		<div class="prose text-md px-2.5 pb-2.5 text-gray-400">
			{#if plant.common}
				<ul>
					{#each plant.common as name, i}
						<li>{name.name}</li>
					{/each}
				</ul>
			{:else}
				<div class="rounded-md bg-slate-100 p-4">
					<p class="text-md">
						We currently don't have any other common names for this entry on record. Feel free to learn about how you can <a href="/contribute" title="Missing details? Submit your own changes for approval." class="border-b hover:text-slate-600">contribute</a> to make this page more complete.
					</p>
				</div>
			{/if}
		</div>
	</div>

	<h3 class="unstyled font-sans p-2.5 border-b text-md font-semibold bg-slate-100 text-slate-500">Symptoms</h3>
	<div class="prose text-md px-2.5 pb-2.5 text-gray-400">
		<ul>
			{#each plant.symptoms as symptom, i}
				<li>
					<span title="View other plants with this symptom.">{symptom.name}</span>
				</li>
			{/each}
		</ul>
	</div>
</div>
