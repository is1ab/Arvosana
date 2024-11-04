import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

type GradeInfo = {
	student_id: string;
	grade: number | null;
};

export const load: PageLoad = async ({ params, fetch }) => {
	const res = await fetch(`/api/grade/${params.semester}/${params.homework_name}`);
	if (!res.ok) {
		error(res.status, await res.json());
	}

	const info: GradeInfo[] = await res.json();

	return {
		data: {
			semester: params.semester,
			name: params.homework_name,
			info
		}
	};
};
