<header class="header">
	<div class="header__top">Top header</div>
	<div class="header__tags">Bottom header</div>
</header>

<style>
	.header {
		height: 15vh;
		width: 100vw;
		max-height: 120px;

		text-align: center;

		font-size: 2em;

		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;

		font-family: Verdana, Geneva, Tahoma, sans-serif;
	}

	.header__top,
	.header__tags {
		height: 50%;
		width: 100%;

		background-color: bisque;

		display: flex;
		justify-content: center;
		align-items: center;
	}

	.header__tags {
		background-color: coral;
	}
</style>
