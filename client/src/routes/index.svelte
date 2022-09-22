<script context="module">
	export async function load({ session }) {
		// console.log('🚀 ~ file: index.svelte ~ line 3 ~ load ~ ctx', session);
		return {
			props: {
				token: session.user.username
			}
		};
	}
</script>

<script>
	// @ts-nocheck

	import Home from '../components/home.svelte';

	// export const ssr = false;
	export let token;
	console.log('🚀 ~ file: index.svelte ~ line 19 ~ user', token);
	// console.log('🚀 ~ file: index.svelte ~ line 18 ~ data', data);
</script>

<!-- <p class="text-3xl ">Hello world!</p> -->
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
