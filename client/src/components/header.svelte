<script>
	import { user } from '@stores/user';
	import Login from '@components/login/login.svelte';
	import Img from './common/img.svelte';

	let showLoginPopup = false;

	function togglePopup() {
		showLoginPopup = !showLoginPopup;
	}
</script>

<header class="header z-10">
	<div class="header__top">
		<div class="start">
			<Img className="hidden sm:block h-6" alt="hamburger" src="icons/menu-outline.svg" />
			<div class="logo">
				<Img className="h-6 logo__img" src="icons/logo-youtube.svg" alt="youtube logo" />
				<p class="logo__name">Youtube</p>
			</div>
		</div>
		<div class="flex items-center gap-4 rounded-lg">
      <div class="flex">
        <input
				class="search__query hidden rounded-l-full sm:block focus:outline-none focus:border-indigo-500"
				type="text"
				name="search__text"
				id="search__text"
				placeholder="Search"
        />
        <button class="search__icon rounded-r-full">
          <Img className="h-5" src="icons/search-outline.svg" alt="search icon" />
        </button>
      </div>
      <button class="sm:flex rounded-full items-center justify-center w-10 h-10 bg-gray-100 hidden">
        <Img className="h-6 search__mic-img" src="icons/mic.svg" alt="mic" />
      </button>
		</div>
		<div class="end">
			<Img className="h-6 hidden sm:block" src="icons/videocam-outline.svg" alt="video create" />
			<Img
				className="h-6 hidden sm:block"
				src="icons/notifications-outline.svg"
				alt="notifications"
			/>

			{#if !$user.isLoggedIn}
				<button class="login__button flex justify-around items-center	" on:click={togglePopup}>
					<Img className="h-4/5" src="./icons/nav/login.svg" alt="login" />
					<p class="hidden sm:block font-medium text-base">LOG IN</p>
				</button>
			{:else}
				<Img className="profile" src="https://picsum.photos/id/237/50/50" alt="profile pic" />
			{/if}
		</div>
	</div>

	{#if showLoginPopup}
		<Login closePopup={togglePopup} />
	{/if}
</header>

<style>
	.header {
		height: 60px;
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
	.search__iconh-5 {
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

	@media (max-width: 600px) {
		.hide-in-mobile {
			display: none;
		}

		.login__button {
			border: unset;
			width: unset;
		}

		.search {
			margin-left: auto;
		}

		.search__icon {
			background: unset;
		}

		.start {
			margin-left: 10px;
		}
		.end {
			margin-right: 10px;
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
