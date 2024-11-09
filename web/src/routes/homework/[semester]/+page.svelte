<script lang="ts">
	import type { PageData } from './$types';

	import * as Table from '$lib/components/ui/table';

	let { data }: { data: PageData } = $props();
</script>

<h1 class="scroll-m-20 text-4xl font-bold tracking-tight lg:text-5xl">
	{data.semester.toUpperCase()}
	{' '}Homeworks
</h1>

<div class="my-6 w-full">
	<Table.Root>
		<Table.Header>
			<Table.Row>
				<Table.Head class="w-1/2">Student ID</Table.Head>
				<Table.Head>Deadline</Table.Head>
			</Table.Row>
		</Table.Header>

		<Table.Body>
			{#each data.info as i}
				<Table.Row>
					<Table.Cell>
						<a
							href={`/homework/${data.semester}/${i.name}`}
							class="text-primary underline hover:text-primary/70"
						>
							{i.name}
						</a>
					</Table.Cell>
					<Table.Cell>
						{#if new Date() > new Date(i.end_at * 1000)}
							Overdue
						{:else if new Date() < new Date(i.begin_at * 1000)}
							Unreleased
						{:else}
							Ending at {new Date(i.end_at * 1000).toLocaleString()}
						{/if}
					</Table.Cell>
				</Table.Row>
			{/each}
		</Table.Body>
	</Table.Root>
</div>
