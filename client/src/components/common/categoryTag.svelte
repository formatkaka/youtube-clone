<script>
	// @ts-nocheck

	export let tagName,
		selected = false;
</script>

<button class:header__tags-btn--selected={selected} class="bg-gray-300 px-5 py-1 rounded-lg whitespace-nowrap">{tagName}</button>

<style>
	.header__tags-btn--selected {
		background-color: black;
		color: white;
	}
</style>
