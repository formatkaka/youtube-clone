<script>
	import CategoryTag from '@components/common/categoryTag.svelte';
	import MediaQuery from '@components/common/MediaQuery.svelte';
</script>

<MediaQuery query="(min-width: 480px)" let:matches>
	{#if matches}
		<!-- content here -->
		<div class="w-full ">
			<div class="w-full flex gap-2 mb-4 overflow-x-scroll no-scrollbar">
				<CategoryTag tagName="All" selected />
				<CategoryTag tagName="Lorem" />
				<CategoryTag tagName="Ipsum" />
				<CategoryTag tagName="Dolor" />
				<CategoryTag tagName="Sit" />
				<CategoryTag tagName="Amet" />
			</div>

			{#each new Array(12) as _}
				<div class="mb-1 flex justify-center items-start mt-2">
					<img class="rounded" src="https://picsum.photos/180/100" alt="hello" />
					<div class="text-sm ml-2">
						<p class="line-clamp-2 font-bold mb-1">
							दादा गोवा छोड़ दो हमे | Sanjay, Arshad, Javed, Riteish, Aashish, Javed | Dhamaal Part
							3
						</p>
						<p class="text-xs text-gray-600	mb-1">Shemaroo Comedy</p>
						<p class="text-xs text-gray-600	">
							<span>11M Views</span>
							<span>&#183;</span>
							<span>2 years ago</span>
						</p>
					</div>
				</div>
			{/each}
		</div>
	{/if}
</MediaQuery>

<MediaQuery query="(max-width: 481px)" let:matches>
	{#if matches}
		<p class="px-2 border-t-2 py-2 text-sm">Up Next</p>
		<div class="w-full ">
			{#each new Array(12) as _}
				<div class="mb-1 flex flex-col  items-center">
					<img class="w-full" src="https://picsum.photos/180/100" alt="hello" />
					<div class="text-sm ml-2 mt-4 px-2">
						<div class="flex gap-2">
							<img class="w-12 h-12 rounded-full" src="https://picsum.photos/50/50" alt="hello" />
							<div>
								<p class="line-clamp-2 font-bold mb-1">
									दादा गोवा छोड़ दो हमे | Sanjay, Arshad, Javed, Riteish, Aashish, Javed | Dhamaal
									Part 3
								</p>
								<p class="text-xs text-gray-600	mb-1">
									<span> Shemaroo Comedy</span>
									<span>&#183;</span>
									<span>11M Views</span>
								</p>
							</div>
						</div>
					</div>
				</div>
			{/each}
		</div>
	{/if}
</MediaQuery>

<style>
</style>
