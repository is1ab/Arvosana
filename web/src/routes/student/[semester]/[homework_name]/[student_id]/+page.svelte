<script lang="ts">
	import type { PageData } from './$types';

	import * as Table from '$lib/components/ui/table';
	import { onMount } from 'svelte';
	import { dev } from '$app/environment';
	import { invalidateAll } from '$app/navigation';

	let { data }: { data: PageData } = $props();

	onMount(() => {
		// HACK: vite proxy doesn't work with this
		const host = dev ? 'http://127.0.0.1:8080' : '';
		const evtSrc = new window.EventSource(host + '/api/sse?stream=submit');

		evtSrc.onmessage = () => {
			// TODO: maybe add debounce?
			invalidateAll();
		};

		return () => evtSrc.close();
	});
</script>

<h1 class="scroll-m-20 text-4xl font-bold tracking-tight lg:text-5xl">
	{data.semester.toUpperCase()}
	{data.homework_name}
	{data.student_id}
</h1>

<div class="my-6 w-full">
	<Table.Root>
		<Table.Header>
			<Table.Row>
				<Table.Head class="w-2/5">Grade</Table.Head>
				<Table.Head>Submitted Time</Table.Head>
			</Table.Row>
		</Table.Header>

		<Table.Body>
			{#each data.info as i}
				<Table.Row>
					<Table.Cell>
						{i.grade}
					</Table.Cell>
					<Table.Cell>{new Date(i.submitted_at * 1000).toLocaleString()}</Table.Cell>
				</Table.Row>
			{/each}
		</Table.Body>
	</Table.Root>
</div>
