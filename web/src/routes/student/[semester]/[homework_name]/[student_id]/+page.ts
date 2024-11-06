import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

type SubmitInfo = {
	grade: number;
	submitted_at: number; // unix timestamp
};

export const load: PageLoad = async ({ params, fetch }) => {
	const res = await fetch(
		`/api/grade/${params.semester}/${params.homework_name}/${params.student_id}`
	);
	if (!res.ok) {
		error(res.status, await res.json());
	}

	const info: SubmitInfo[] = await res.json();

	return {
		info,
		...params
	};
};
