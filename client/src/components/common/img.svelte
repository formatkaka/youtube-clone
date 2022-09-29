<script>
	// @ts-nocheck

	const BASE_URL = 'https://ytspaces.sgp1.digitaloceanspaces.com/static/';
	export let src, alt, className;
</script>

<img class={className} src={BASE_URL + src} {alt} />
