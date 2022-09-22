<script>
	import Home from '@components/home.svelte';
	import { page } from '$app/stores';
	import { user } from '@stores/user';

	if ($page.data.user) {
		user.loginUser($page.data.user);
	}
</script>

<Home />

<style>
	:global(body) {
		padding: 0px;
		margin: 0px;

		font-family: 'Open Sans', sans-serif;
		overflow-x: hidden;
	}

	:global(button) {
		background: none;
		color: inherit;
		border: none;
		padding: 0;
		font: inherit;
		cursor: pointer;
		outline: inherit;
	}

	:global(.no-scrollbar::-webkit-scrollbar) {
		display: none;
	}
	/* Hide scrollbar for Chrome, Safari and Opera */

	/* Hide scrollbar for IE, Edge and Firefox */
	.no-scrollbar {
		-ms-overflow-style: none; /* IE and Edge */
		scrollbar-width: none; /* Firefox */
	}
</style>
