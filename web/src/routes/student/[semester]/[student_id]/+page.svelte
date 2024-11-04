<script lang="ts">
	import type { PageData } from './$types';

	import * as Table from '$lib/components/ui/table';

	let { data }: { data: PageData } = $props();
</script>

<h1 class="scroll-m-20 text-4xl font-bold tracking-tight lg:text-5xl">
	{data.data.semester.toUpperCase()}
	{data.data.student_id}
</h1>

<div class="my-6 w-full">
	<Table.Root>
		<Table.Header>
			<Table.Row>
				<Table.Head>Homework</Table.Head>
				<Table.Head>Grade</Table.Head>
			</Table.Row>
		</Table.Header>

		<Table.Body>
			{#each data.data.info as i}
				<Table.Row>
					<Table.Cell>
						<a
							href={`/homework/${data.data.semester}/${i.name}`}
							class="text-primary underline hover:text-primary/70"
						>
							{i.name}
						</a>
					</Table.Cell>
					<Table.Cell>{i.grade ?? 'Not submitted'}</Table.Cell>
				</Table.Row>
			{/each}
		</Table.Body>
	</Table.Root>
</div>
