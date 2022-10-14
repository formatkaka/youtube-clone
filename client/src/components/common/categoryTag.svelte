<script>
	// @ts-nocheck

	export let tagName,
		selected = false;
</script>

<button class:header__tags-btn--selected={selected} class="header__tags-btn">{tagName}</button>

<style>
	.header__tags-btn {
		background: rgba(0, 0, 0, 0.05);
		padding: 5px 20px;
		border-radius: 20px;
		white-space: nowrap;
		border: 1px solid rgba(0, 0, 0, 0.1);

		font-size: 15px;
	}

	.header__tags-btn--selected {
		background-color: black;
		color: white;
	}
</style>
