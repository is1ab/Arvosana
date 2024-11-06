<script lang="ts">
	import type { PageData } from './$types';

	import * as Table from '$lib/components/ui/table';

	let { data }: { data: PageData } = $props();
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
