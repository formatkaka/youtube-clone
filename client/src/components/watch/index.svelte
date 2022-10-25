<script>
	import VideoPlayer from './videoPlayer.svelte';
	import VideoSuggestions from './suggestions/videoSuggestions.svelte';
	import VideoMeta from './videoMeta/index.svelte';
</script>

<div class="flex flex-col w-screen px-0 sm:px-12 gap-6 mt-4 sm:flex-row">
	<div class="w-full sm:w-9/12">
		<VideoPlayer />
		<VideoMeta />
	</div>

	<div class="w-full sm:w-1/4">
		<VideoSuggestions />
	</div>
</div>

<style>
</style>
