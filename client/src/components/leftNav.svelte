<script>
	const nav = {
		first: [
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Home'
			},
			{
				img: 'icons/nav/dice-outline.svg',
				text: 'Explore'
			},
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Shorts'
			},
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Subscriptions'
			},
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Originals'
			}
		],
		second: [
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Home'
			},
			{
				img: 'icons/nav/dice-outline.svg',
				text: 'Explore'
			},
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Shorts'
			},
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Subscriptions'
			},
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Originals'
			},
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Originals'
			}
		],
		third: [
			{
				img: 'https://picsum.photos/50',
				text: 'T-Series'
			},
			{
				img: 'https://picsum.photos/50',
				text: 'Visa2Explore'
			},
			{
				img: 'https://picsum.photos/50',
				text: 'After Skool'
			},
			{
				img: 'https://picsum.photos/50',
				text: 'Cinema Beyond Entertainment'
			},
			{
				img: 'https://picsum.photos/50',
				text: 'Rajshekhar Mamidana'
			},
			{
				img: 'https://picsum.photos/50',
				text: 'illneas'
			},
			{
				img: 'https://picsum.photos/50',
				text: 'Food Food'
			}
		],
		fourth: [
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Movies'
			},
			{
				img: 'icons/nav/dice-outline.svg',
				text: 'Gaming'
			},
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Live'
			},
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Fashion and Beauty'
			},
			{
				img: 'icons/nav/home-outline.svg',
				text: 'Sports'
			}
		]
	};
</script>

<nav class="nav">
	{#each Object.values(nav) as section}
		<div class="nav__section">
			{#each section as item}
				<span class="item">
					<img class="item__logo" src={item.img} alt="Home" />
					<p class="item__text">{item.text}</p>
				</span>
			{/each}
		</div>
	{/each}
</nav>

<style>
	.nav {
		position: fixed;
		left: 0px;
		top: 60px;

		height: calc(100vh - 60px);

		width: 12vw;

		/* background-color: burlywood; */

		display: flex;
		justify-content: flex-start;
		align-items: center;

		flex-direction: column;

		font-size: 2em;
		font-family: Verdana, Geneva, Tahoma, sans-serif;

		overflow-y: scroll;

		z-index: 3;
	}

	.nav__section {
		display: flex;
		flex-direction: column;
		border-bottom: 1px solid rgba(0, 0, 0, 0.1);

		width: 100%;

		box-sizing: border-box;

		padding: 10px 0px;
	}

	.item {
		display: flex;
		justify-content: flex-start;
		align-items: center;

		width: 100%;

		height: 40px;
	}

	.item__logo {
		height: 20px;
		margin-left: 20px;

		border-radius: 50px;
	}

	.item__text {
		font-size: 14px;
		font-family: Roboto;
		margin-left: 20px;
	}
</style>
