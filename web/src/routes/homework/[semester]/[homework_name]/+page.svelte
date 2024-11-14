<script lang="ts">
	import type { PageData } from './$types';

	import * as Table from '$lib/components/ui/table';
	import { Separator } from '$lib/components/ui/separator';

	let { data }: { data: PageData } = $props();
</script>

<div class="flex items-end justify-between">
	<h1 class="scroll-m-20 text-4xl font-bold tracking-tight lg:text-5xl">
		{data.semester.toUpperCase()}
		{data.homework_name}
	</h1>

	<a href={`/homework/${data.semester}`} class="text-primary underline hover:text-primary/70">
		{data.semester}
		{' '} Homeworks
	</a>
</div>

<div class="my-6 flex justify-between text-muted-foreground">
	<p>
		{new Date(data.begin_at * 1000).toLocaleString()}
		-
		{new Date(data.end_at * 1000).toLocaleString()}
	</p>

	<p>
		{data.submitted}/{data.info.length} Submitted
	</p>
</div>

<Separator />

<div class="my-6 w-full">
	<Table.Root>
		<Table.Header>
			<Table.Row>
				<Table.Head>Student ID</Table.Head>
				<Table.Head>Grade</Table.Head>
			</Table.Row>
		</Table.Header>

		<Table.Body>
			{#each data.info as i}
				<Table.Row>
					<Table.Cell>
						<a
							href={`/student/${data.semester}/${i.student_id}`}
							class="text-primary underline hover:text-primary/70"
						>
							{i.student_id}
						</a>
					</Table.Cell>
					<Table.Cell>
						{#if i.grade}
							<a
								href={`/student/${data.semester}/${data.homework_name}/${i.student_id}`}
								class="text-primary underline hover:text-primary/70"
							>
								{i.grade}
							</a>
						{:else}
							Not submitted
						{/if}
					</Table.Cell>
				</Table.Row>
			{/each}
		</Table.Body>
	</Table.Root>
</div>
