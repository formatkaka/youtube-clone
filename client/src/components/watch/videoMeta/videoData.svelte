<script>
	import Img from '@components/common/img.svelte';
	import MediaQuery from '@components/common/MediaQuery.svelte';
</script>

<MediaQuery query="(max-width: 480px)" let:matches>
	{#if matches}
		<div>
			<h1 class="px-2 pt-2 text-lg font-normal">
				Taarak Mehta Ka Ooltah Chashmah | तारक मेहता का उल्टा चश्मा | Episode 862 | 15th March, 2021
			</h1>
		</div>
		<div class="px-2 flex text-xs">
			<span>1.4M views</span>
			<span class="px-1">&#183;</span>
			<span>3 days ago</span>
		</div>
		<div class="w-full px-2 pt-2 flex justify-around">
			<div class="flex justify-between items-center text-xs flex-col">
				<Img className="w-6" src="icons/watch/like.svg" alt="Like" />
				<p>84k</p>
			</div>
			<div class="flex justify-between items-center text-xs flex-col">
				<Img className="w-6" src="icons/watch/dislike.svg" alt="Like" />
				<p>Dislike</p>
			</div>
			<div class="flex justify-between items-center text-xs flex-col">
				<Img className="w-6" src="icons/watch/share.svg" alt="Like" />
				<p>Share</p>
			</div>
			<div class="flex justify-between items-center text-xs flex-col">
				<Img className="w-6" src="icons/watch/add-to-playlist-mobile.png" alt="Like" />
				<p>Save</p>
			</div>
			<div class="flex justify-between items-center text-xs flex-col">
				<Img className="w-6" src="icons/watch/report.svg" alt="Like" />
				<p>Report</p>
			</div>
		</div>
	{/if}
</MediaQuery>

<MediaQuery query="(min-width: 480px)" let:matches>
	{#if matches}
		<div>
			<h1 class="pt-2 text-xl font-normal">
				Taarak Mehta Ka Ooltah Chashmah | तारक मेहता का उल्टा चश्मा | Episode 862 | 15th March, 2021
			</h1>
		</div>
		<div class="flex mt-4">
			<div class="flex gap-2">
				<img src="https://picsum.photos/35" alt="Logo" class="w-10 h-10 rounded-full" />
				<div class="">
					<p class="whitespace-nowrap text-sm">Rajshekhar Mamidana</p>
					<p class="text-xs">282K subscribers</p>
				</div>
				<button
					class="text-white bg-black px-4  ml-2 rounded-full flex items-center justify-between text-sm"
					>Subscribe</button
				>
			</div>

			<div class="px-2 flex items-center gap-2 ml-auto">
				<div
					class="flex justify-between items-center text-sm gap-2 bg-[#E8E8E8] px-4 py-2 rounded-full"
				>
					<Img className="w-6" src="icons/watch/like.svg" alt="Like" />
					<p>84k</p>

					<Img className="w-6 mx-2" src="icons/watch/dislike.svg" alt="Like" />
				</div>
				<div
					class="flex justify-between items-center text-sm gap-2 bg-[#E8E8E8] px-4 py-2 rounded-full"
				>
					<Img className="w-6" src="icons/watch/share.svg" alt="Like" />
					<p>Share</p>
				</div>
				<div
					class="flex justify-between items-center text-sm gap-2 bg-[#E8E8E8] px-4 py-2 rounded-full"
				>
					<Img className="w-6" src="icons/watch/add-to-playlist-desktop.svg" alt="Like" />
					<p>Save</p>
				</div>
				<div
					class="flex justify-between items-center text-sm gap-2 bg-[#E8E8E8] px-4 py-2 rounded-full"
				>
					<Img className="w-6" src="icons/watch/report.svg" alt="Like" />
					<p>Report</p>
				</div>
			</div>
		</div>
	{/if}
</MediaQuery>
