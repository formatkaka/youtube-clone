<script></script>

<div class="tile-skeleton">
	<div class="tile-skeleton__poster" />
	<div class="tile-skeleton__desc">
		<div class="desc__logo-shadow" />
		<div class="desc flex-column">
			<div class="desc__first flex-sb" />
			<div class="desc__second flex-column" />
		</div>
	</div>
</div>

<style>
	.tile-skeleton__poster {
		display: block;
		width: 100%;
		padding-top: 57.1%;
		background-color: gray;

		z-index: -1;
	}

	.tile-skeleton__desc {
		display: flex;
		padding: 10px 0px;
		gap: 10px;
	}

	.desc__logo-shadow {
		min-width: 35px;
		height: 35px;
		background: gray;
		border-radius: 30px;

		/* position: absolute; */
		/* z-index: -1; */
	}

	.flex-column {
		display: flex;
		flex-direction: column;
	}

	.desc {
		gap: 10px;
		width: 100%;
	}

	.desc__first {
		background: gray;
		height: 20px;
		width: 90%;
	}

	.desc__second {
		background: gray;
		height: 20px;
		width: 50%;
	}
</style>
