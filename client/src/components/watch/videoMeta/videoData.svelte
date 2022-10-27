<script>
	import Img from '@components/common/img.svelte';
	import MediaQuery from '@components/common/MediaQuery.svelte';
</script>

<MediaQuery query="(max-width: 480px)" let:matches>
	{#if matches}
		<div>
			<h1 class="px-4 pt-2 text-lg font-normal">
				Taarak Mehta Ka Ooltah Chashmah | तारक मेहता का उल्टा चश्मा | Episode 862 | 15th March, 2021
			</h1>
		</div>
		<div class="px-4 flex text-xs">
			<span>1.4M views</span>
			<span class="px-1">&#183;</span>
			<span>3 days ago</span>
		</div>
		<div class="w-full px-4 mt-4 flex gap-2">
			<img src="https://picsum.photos/35" alt="Logo" class="w-10 h-10 rounded-full" />
			<div class="">
				<p class="whitespace-nowrap text-sm">Rajshekhar Mamidana</p>
				<p class="text-xs">282K subscribers</p>
			</div>
			<button
				class="text-white bg-black px-4  ml-auto rounded-full flex items-center justify-between text-sm"
				>Subscribe</button
			>
		</div>
		<div class="flex gap-4 pl-4 mt-4 w-full overflow-scroll no-scrollbar">
			<!-- <div class="flex gap-2 ml-auto 	"> -->
			<div class="flex justify-center text-sm gap-2 bg-[#E8E8E8] px-8 py-2 rounded-full divide-x-2">
				<Img className="w-6" src="icons/watch/like.svg" alt="Like" />
				<p class="mr-4">84k</p>

				<Img className="w-6" src="icons/watch/dislike.svg" alt="Like" />
			</div>
			<div
				class="flex justify-center items-center text-sm gap-2 bg-[#E8E8E8] px-6 py-2 rounded-full"
			>
				<Img className="w-6" src="icons/watch/share.svg" alt="Like" />
				<p>Share</p>
			</div>
			<div
				class="flex justify-center items-center text-sm gap-2 bg-[#E8E8E8] px-6 py-2 rounded-full"
			>
				<Img className="w-6" src="icons/watch/add-to-playlist-desktop.svg" alt="Like" />
				<p>Save</p>
			</div>
			<div
				class="flex justify-center items-center text-sm gap-2 bg-[#E8E8E8] px-6 py-2 rounded-full"
			>
				<Img className="w-4" src="icons/watch/report.svg" alt="Like" />
				<p>Report</p>
			</div>
			<!-- </div> -->
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
					<Img className="w-4" src="icons/watch/report.svg" alt="Like" />
					<p>Report</p>
				</div>
			</div>
		</div>
	{/if}
</MediaQuery>
