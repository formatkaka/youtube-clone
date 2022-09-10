<script>
	import Icon from '@iconify/svelte';

	let loggedIn = false;
</script>

<header class="header">
	<div class="header__top">
		<div class="start">
			<img class="svg" alt="hamburger" src="icons/menu-outline.svg" />
			<div class="logo">
				<img class="svg logo__img" src="icons/logo-youtube.svg" alt="youtube logo" />
				<p class="logo__name">Youtube</p>
			</div>
		</div>
		<div class="search">
			<input
				class="search__query"
				type="text"
				name="search__text"
				id="search__text"
				placeholder="Search"
			/>
			<button class="search__icon">
				<img class="search__icon-img" src="icons/search-outline.svg" alt="search icon" />
			</button>
			<button class="search__mic">
				<img class="svg search__mic-img" src="icons/mic.svg" alt="mic" />
			</button>
		</div>
		<div class="end">
			<img class="svg" src="icons/videocam-outline.svg" alt="video create" />
			<img class="svg" src="icons/notifications-outline.svg" alt="notifications" />

			{#if !loggedIn}
				<button class="login__button flex justify-around items-center	">
					<img class="h-4/5" src="./icons/nav/login.svg" alt="login" />
					<p class="font-medium text-base">SIGN IN</p>
				</button>
			{:else}
				<img class="profile" src="https://picsum.photos/id/237/50/50" alt="profile pic" />
			{/if}
		</div>
	</div>

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
</header>

<style>
	.header {
		height: 120px;
		width: 100vw;
		/* max-height: 120px; */

		text-align: center;

		font-size: 2em;

		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;

		font-family: Verdana, Geneva, Tahoma, sans-serif;

		position: sticky;
		top: 0px;
		z-index: 2;
		background: white;
	}

	.header__top {
		height: 50%;
		width: 100%;

		/* background-color: bisque; */

		display: flex;
		justify-content: center;
		align-items: center;
	}

	.start {
		display: flex;
		justify-content: center;
		align-items: center;
		margin-left: 20px;
		gap: 20px;
		/* align-self: start; */
	}

	.header__top {
		justify-content: space-between;
		/* position: sticky; */
		/* top: 0px; */
	}

	.header__tags {
		/* background-color: coral; */
	}

	.svg {
		height: 25px;
	}

	.logo {
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 2px;
	}

	.logo__name {
		font-size: 25px;
		font-family: Oswald;
	}

	.search {
		display: flex;
		/* align-self: center; */
	}

	.search__query {
		width: 30vw;
		border: 1px solid #ccc;
		font-size: 14px;
		padding: 0px 10px;
	}

	.search__query::placeholder {
		font-size: 14px;
	}

	.search__icon {
		height: 40px;
		width: 60px;
		background: #f8f8f8;

		display: flex;
		justify-content: center;
		align-items: center;
	}
	.search__mic {
		background-color: #f9f9f9;
		border-radius: 50px;
		width: 40px;
		margin-left: 5px;
	}
	.search__icon-img {
		height: 20px;
	}

	.profile {
		border-radius: 50px;
		height: 40px;
	}

	.end {
		display: flex;
		gap: 25px;
		justify-content: center;
		align-items: center;
		margin-right: 30px;
	}
	.login__button {
		height: 40px;
		width: 120px;
		font-size: 12px;
		border: 1px solid black;
		padding: 5px 10px;
		box-sizing: border-box;

		font-family: Roboto;
	}

	.header__tags {
		display: flex;
		gap: 15px;
		margin-left: 12vw;
		overflow-x: scroll;
		width: 88vw;
		height: 50%;
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
</style>
