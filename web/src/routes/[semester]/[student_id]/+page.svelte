<script lang="ts">
	import type { PageData } from './$types';

	import * as Table from '$lib/components/ui/table';

	let { data }: { data: PageData } = $props();
</script>

<Table.Root>
	<Table.Header>
		<Table.Row>
			<Table.Head>Homework</Table.Head>
			<Table.Head>Grade</Table.Head>
		</Table.Row>
	</Table.Header>

	<Table.Body>
		{#each data.info as i}
			<Table.Row>
				<Table.Cell>{i.name}</Table.Cell>
				<Table.Cell>{i.grade}</Table.Cell>
			</Table.Row>
		{/each}
	</Table.Body>
</Table.Root>
