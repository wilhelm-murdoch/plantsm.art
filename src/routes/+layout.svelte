<script>
		import '@skeletonlabs/skeleton/themes/theme-skeleton.css';
		import '@skeletonlabs/skeleton/styles/all.css';
		import '../app.postcss';
		import { AppShell, AppBar, Divider } from '@skeletonlabs/skeleton';
		import SvgIcon from '@skeletonlabs/skeleton/components/SvgIcon/SvgIcon.svelte';

</script>

<AppShell slotSidebarLeft="bg-surface-500/5 w-56 p-4">
    <!-- Header -->
		<svelte:fragment slot="header">
				<!-- Insert the App Bar: -->
				<AppBar>
					<svelte:fragment slot="lead">
						<h1>Plant Smart</h1>
					</svelte:fragment>
					<svelte:fragment slot="trail">
						<Divider vertical borderWidth="border-l-2 opacity-20" />

						<!-- Social -->
						<section class="flex">
							<a class="btn btn-sm" href="https://discord.gg/EXqV7W8MtY" target="_blank" rel="noreferrer" aria-label="Discord">
								<SvgIcon name="discord" viewBox="0 0 640 512" />
							</a>
							<a class="btn btn-sm" href="https://twitter.com/SkeletonUI" target="_blank" rel="noreferrer" aria-label="Twitter">
								<SvgIcon name="twitter" />
							</a>
							<a class="btn btn-sm" href="https://github.com/skeletonlabs/skeleton" target="_blank" rel="noreferrer" aria-label="GitHub">
								<SvgIcon name="github" />
							</a>
						</section>
						<!-- --- -->
					</svelte:fragment>
				</AppBar>
				<!-- --- -->
			</svelte:fragment>
    <slot />
</AppShell>