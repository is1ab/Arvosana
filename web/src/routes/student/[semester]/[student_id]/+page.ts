import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

type GradeInfo = {
	name: string;
	grade: number | null;
};

export const load: PageLoad = async ({ params, fetch }) => {
	const res = await fetch(`/api/student/${params.semester}/${params.student_id}`);
	if (!res.ok) {
		error(res.status, await res.json());
	}

	const info: GradeInfo[] = await res.json();

	return {
		data: {
			info,
			semester: params.semester,
			student_id: params.student_id
		}
	};
};
