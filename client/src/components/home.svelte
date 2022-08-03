<script>
	import Header from './header.svelte';
	import Nav from './leftNav.svelte';
	import VideoTile from './videoTile.svelte';
</script>

<Header />
<Nav />

<div class="feed">
	{#each Array(15) as _, i}
		<VideoTile />
	{/each}
</div>

<style>
	.feed {
		position: relative;
		left: 10vw;

		display: grid;
		grid-template-columns: repeat(4, 1fr);
		column-gap: 20px;
		row-gap: 40px;

		padding: 20px 30px;

		width: 90vw;
		box-sizing: border-box;
	}

	@media (max-width: 700px) {
		.feed {
			grid-template-columns: repeat(3, 1fr);
		}
	}
</style>
