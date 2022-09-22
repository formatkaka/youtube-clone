<script>
	// @ts-nocheck

	import { secondsToHms } from '../utils/time';

	export /**
	 * @type {{ [x: string]: any; }}
	 */
	let data;
	let arr = [1001, 1002, 1003, 1004, 1005, 1006, 1008, 1009, 100, 101];
	function randomIntFromInterval(min, max) {
		// min and max included
		return Math.floor(Math.random() * (max - min + 1) + min);
	}
	let rand = randomIntFromInterval(0, 9);
	let src = `https://picsum.photos/id/${arr[rand]}/350/200`;
</script>

<div class="video-tile flex-column">
	<div class="video-tile__poster">
		<img class="poster__img" {src} alt="Picsum" />
		<div class="poster__duration">{secondsToHms(data['Duration'])}</div>
	</div>
	<div class="video-tile__desc">
		<div class="desc__logo-shadow">
			<img src="https://picsum.photos/35" alt="Logo" class="desc__logo" />
		</div>
		<div class="desc flex-column">
			<div class="desc__first flex-sb">
				<p class="desc__title">
					{data['Title']}
				</p>
			</div>
			<div class="desc__stats flex-column">
				<p class="desc__channel-name">{data['Name']}</p>
				<span class="show-in-mobile">&#183;</span>
				<div class="desc__third flex-sb">
					<p class="desc__views">538K views</p>
					<span class="show-in-mobile">&#183;</span>
					<p class="desc__upload-info">8 days ago</p>
				</div>
			</div>
		</div>
	</div>
</div>

<style>
	.flex-column {
		display: flex;
		flex-direction: column;
	}

	.video-title {
	}

	.video-tile__desc {
		display: flex;
		padding: 10px 0px;
		gap: 10px;
	}

	.flex-sb {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.desc__logo {
		height: 35px;
		border-radius: 100px;
	}

	p {
		margin: 0px;
	}

	.desc__first {
		height: 35px;
	}

	.desc__third {
		justify-content: flex-start;
		gap: 5px;
	}

	.desc {
		gap: 10px;
	}

	.video-tile__poster {
		position: relative;
		width: 100%;
	}

	.video-tile__poster::before {
		display: block;
		content: '';
		width: 100%;
		padding-top: 57.1%;
		background-color: gray;

		/* position: absolute; */
		z-index: -1;
	}

	.desc__logo-shadow {
		min-width: 35px;
		height: 35px;
		background: gray;
		border-radius: 30px;

		/* position: absolute; */
		/* z-index: -1; */
	}

	.poster__img {
		width: 100%;
		position: absolute;
		top: 0px;
		left: 0px;
	}

	.poster__duration {
		position: absolute;
		right: 5px;
		bottom: 5px;

		background-color: black;
		font-size: 12px;
		color: white;

		padding: 2px 5px;
	}

	.show-in-mobile {
		display: none;
	}

	@media (max-width: 600px) {
		.video-tile__desc {
			align-items: center;
			justify-content: center;

			font-size: 12px;
			padding: 10px;
		}

		.desc__stats {
			flex-direction: unset;
			gap: 5px;
			position: relative;
		}

		.desc__title {
			font-weight: bold;
		}

		.desc {
			gap: 2px;
		}

		.show-in-mobile {
			display: unset;
			font-size: 14px;
			font-weight: bold;
		}
	}
</style>
