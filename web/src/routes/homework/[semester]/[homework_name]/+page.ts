import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

type GradeInfo = {
	student_id: string;
	grade: number | null;
};

type HomeworkInfo = {
	begin_at: number;
	end_at: number;
};

export const load: PageLoad = async ({ params, fetch }) => {
	const [gradeRes, hwRes] = await Promise.all([
		fetch(`/api/grade/${params.semester}/${params.homework_name}`),
		fetch(`/api/homework/${params.semester}/${params.homework_name}`)
	]);

	if (!gradeRes.ok) {
		error(gradeRes.status, await gradeRes.json());
	}
	if (!hwRes.ok) {
		error(hwRes.status, await hwRes.json());
	}

	const gradeInfo: GradeInfo[] = await gradeRes.json();
	const hwInfo: HomeworkInfo = await hwRes.json();

	let submitted = 0;

	for (const info of gradeInfo) {
		if (info.grade) {
			submitted += 1;
		}
	}
	return {
		info: gradeInfo,
		submitted,
		...hwInfo,
		...params
	};
};
