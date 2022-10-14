<script>
	import CategoryTag from '@components/common/categoryTag.svelte';
</script>

<div class="w-full ">
	<div class="w-full flex gap-2 mb-2 overflow-x-scroll no-scrollbar">
		<CategoryTag tagName="All" selected />
		<CategoryTag tagName="Lorem" />
		<CategoryTag tagName="Ipsum" />
		<CategoryTag tagName="Dolor" />
		<CategoryTag tagName="Sit" />
		<CategoryTag tagName="Amet" />
	</div>

	{#each new Array(12) as _}
		<div class="bg-lime-300 h-32 mb-1 flex justify-center items-center	">Tile</div>
	{/each}
</div>

<style>
</style>
