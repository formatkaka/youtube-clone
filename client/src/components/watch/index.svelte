<script>
	import VideoPlayer from './videoPlayer.svelte';
	import VideoSuggestions from './suggestions/videoSuggestions.svelte';
	import VideoMeta from './videoMeta/index.svelte';
</script>

<div class="flex w-screen px-12 gap-6 mt-4">
	<div class="w-9/12">
		<VideoPlayer />
		<VideoMeta />
	</div>

	<div class="w-1/4">
		<VideoSuggestions />
	</div>
</div>

<style>
</style>
