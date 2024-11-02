<script lang="ts">
	import { ModeWatcher } from 'mode-watcher';
	import '../app.css';
	import ModeToggle from '$lib/components/mode-toggle.svelte';
	let { children } = $props();
</script>

<ModeWatcher />
<nav class="w-full px-4">
	<div class="mx-auto flex max-w-5xl items-center justify-between py-4">
		<a href="/" class="text-2xl font-bold">Arvosana</a>

		<ModeToggle />
	</div>
</nav>

<div class="w-full">
	<main class="mx-auto max-w-4xl px-4">
		{@render children()}
	</main>
</div>
