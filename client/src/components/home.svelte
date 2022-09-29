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

<div class="header__tags no-scrollbar">
	<!-- <p>Hello</p> -->
	{#each Array(4) as _}
		<button class="header__tags-btn">All</button>
		<button class="header__tags-btn">Mixes</button>
		<button class="header__tags-btn">Music</button>
		<button class="header__tags-btn">Bollywood</button>
		<button class="header__tags-btn">Street Food</button>
	{/each}
</div>
<!-- <Header /> -->
<Nav />

<div class="feed">
	{#each Array(feed.length) as _, i}
		<VideoTile data={feed[i]} />
	{/each}
	{#each Array(8) as _, i}
		<VideoTileSkeleton />
	{/each}
</div>

<div class="continuation-row" />

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

	.header__tags {
		display: flex;
		gap: 15px;
		margin-left: 12vw;
		overflow-x: scroll;
		width: 88vw;
		height: 60px;
		align-items: center;

		border-bottom: 1px solid rgba(0, 0, 0, 0.1);
		border-top: 1px solid rgba(0, 0, 0, 0.1);
	}

	.header__tags-btn {
		background: rgba(0, 0, 0, 0.05);
		padding: 5px 20px;
		border-radius: 20px;
		white-space: nowrap;
		border: 1px solid rgba(0, 0, 0, 0.1);

		font-size: 15px;
	}

	.header__tags-btn:first-child {
		margin-left: 30px;
	}

	.header__tags-btn:last-child {
		margin-right: 30px;
	}
	@media (max-width: 600px) {
		.feed,
		.continuation-row {
			left: 0px;
			width: 100vw;
			padding: 30px 0px;
			grid-template-columns: repeat(1, 1fr);
			row-gap: unset;
		}

		.header__tags {
			width: 100vw;
			margin-left: 0px;
		}

		.header__tags-btn:first-child {
			margin-left: 10px;
		}

		.header__tags-btn:last-child {
			margin-right: 10px;
		}
	}
</style>
