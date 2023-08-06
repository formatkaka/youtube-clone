<script>
	import Header from './header.svelte';
	import Nav from './leftNav.svelte';
	import VideoTile from './videoTile.svelte';
	import VideoTileSkeleton from './videoTileSkeleton.svelte';
	import { onMount } from 'svelte';
	import CategoryTag from './common/categoryTag.svelte';

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



<div class="flex w-full mt-4">
  <Nav />

  <div class="flex flex-col gap-8 ml-left_nav w-feed px-5">
    <div class="flex gap-4 w-full no-scrollbar pr-5">
      {#each Array(4) as _}
        <CategoryTag tagName="All" selected={true} />
        <CategoryTag tagName="Mixes" />
        <CategoryTag tagName="Music" />
        <CategoryTag tagName="Bollywood" />
        <CategoryTag tagName="Street Food" />
      {/each}
    </div>

    <div class="feed w-full grid-cols-1 sm:grid-cols-3 2xl:grid-cols-4">
      {#each Array(feed.length) as _, i}
        <VideoTile data={feed[i]} />
      {/each}
      {#each Array(8) as _, i}
        <VideoTileSkeleton />
      {/each}
    </div>

    <div class="continuation-row" />
  </div>
</div>

<style>
	.feed,
	.continuation-row {
		position: relative;

		display: grid;
		column-gap: 20px;
		row-gap: 40px;
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
