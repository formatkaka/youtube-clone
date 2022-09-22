<script>
	import Header from './header.svelte';
	import Nav from './leftNav.svelte';
	import VideoTile from './videoTile.svelte';
	import VideoTileSkeleton from './videoTileSkeleton.svelte';
	import { onMount } from 'svelte';

	/**
	 * @type {string | any[]}
	 */
	var feed = [];
	var limit = 12;
	var page = 0;

	onMount(() => {
		let options = {
			root: null,
			rootMargin: '0px',
			threshold: 0.1
		};

		let observer = new IntersectionObserver(() => {
			page = page + 1;
			getData();
		}, options);

		let targetEl = document.querySelector('.continuation-row');
		// @ts-ignore
		observer.observe(targetEl);
	});

	onMount(() => {
		page = page + 1;
		getData();
	});

	function getData() {
		fetch(`http://localhost:8080/feed?page=${page}&limit=${limit}`)
			.then((res) => res.json())
			.then((res) => {
				feed = [...feed, ...res.response];
				console.log(feed);
			});
	}
</script>

<Header />
<Nav />

<div class="feed">
	{#each Array(feed.length) as _, i}
		<VideoTile data={feed[i]} />
	{/each}
</div>
<div class="continuation-row">
	{#each Array(8) as _, i}
		<VideoTileSkeleton />
	{/each}
</div>

<style>
	.feed,
	.continuation-row {
		position: relative;
		left: 12vw;

		display: grid;
		grid-template-columns: repeat(4, 1fr);
		column-gap: 20px;
		row-gap: 40px;

		padding: 20px 30px;

		width: 88vw;
		box-sizing: border-box;

		background: #f9f9f9;
	}

	@media (max-width: 600px) {
		.feed {
			left: 0px;
			width: 100vw;
			padding: 30px 0px;
			grid-template-columns: repeat(1, 1fr);
			row-gap: unset;
		}
	}
</style>
