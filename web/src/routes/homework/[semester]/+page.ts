import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

type HomeworkInfo = {
	name: string;
	begin_at: number;
	end_at: number;
};

export const load: PageLoad = async ({ params, fetch }) => {
	const res = await fetch(`/api/homework/${params.semester}`);
	if (!res.ok) {
		error(res.status, await res.json());
	}

	const info: HomeworkInfo[] = await res.json();

	return {
		info,
		...params
	};
};
