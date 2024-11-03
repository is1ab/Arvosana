import type { PageLoad } from './$types';

type GradeInfo = {
	name: string;
	grade: number;
};

export const load: PageLoad = async ({ params, fetch }) => {
	const res = await fetch(`/api/student/${params.semester}/${params.student_id}`);
	const info: GradeInfo[] = await res.json();

	return {
		data: {
			info,
			semester: params.semester,
			student_id: params.student_id
		}
	};
};
